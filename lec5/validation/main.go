package main

import (
	"fmt"
	"net/http"
	"strconv"

	// валидатор
	"github.com/asaskevich/govalidator"
)

// http://127.0.0.1:8080/?priority=low&subject=Hello!&inner=ignored&id=12&recipient=test@host.ru

type SendMessage struct {
	Id        int    `valid:"range(1|100),optional"`
	Priority  string `valid:"in(low|normal|high),required"`
	Recipient string `valid:"email,required"`
	Subject   string `valid:"msgSubject"`
	Inner     string `valid:"-"`
	flag      int
}

func handler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("request " + r.URL.String() + "\n\n"))

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	msg := &SendMessage{
		Id:        id,
		Priority:  r.URL.Query().Get("priority"),
		Recipient: r.URL.Query().Get("recipient"),
		Subject:   r.URL.Query().Get("subject"),
		Inner:     r.URL.Query().Get("inner"),
	}

	w.Write([]byte(fmt.Sprintf("Msg: %#v\n\n", msg)))

	_, err := govalidator.ValidateStruct(msg)

	if err != nil {

		if allErrs, ok := err.(govalidator.Errors); ok {
			for _, fld := range allErrs.Errors() {
				data := []byte(fmt.Sprintf("field: %#v\n\n", fld))
				w.Write(data)
			}
		}

		w.Write([]byte(fmt.Sprintf("error: %s\n\n", err)))
	} else {
		w.Write([]byte(fmt.Sprintf("msg is correct\n\n")))
	}
}

func init() {
	govalidator.CustomTypeTagMap.Set("msgSubject", govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		subject, ok := i.(string)
		if !ok {
			return false
		}
		if len(subject) == 0 || len(subject) > 10 {
			return false
		}
		return true
	}))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
