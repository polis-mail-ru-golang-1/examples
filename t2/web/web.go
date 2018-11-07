package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/polis-mail-ru-golang-1/examples/t2/index"
)

// Web provides web user interface
type Web struct {
	index   index.Index
	address string
}

// New web interface
func New(index index.Index, address string) Web {
	return Web{
		index:   index,
		address: address,
	}
}

// Search handler to perform index usage
func (web Web) Search(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	if search == "" {
		fmt.Fprintf(w, "Usage: %s/?search=query\n", web.address)
		return
	}
	result, err := web.index.Search(search)
	if err != nil {
		fmt.Fprintln(w, "Error: ", err)
		return
	}
	fmt.Fprintln(w, "Query: ", search)
	for _, row := range result {
		fmt.Fprintln(w, row.File, " -- ", row.Count)
	}
}

// Count handler shows info about index
func (web Web) Count(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Count: ", web.index.Info().Count)
}

// Start http server
func (web Web) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", web.Search)
	mux.HandleFunc("/count", web.Count)

	server := http.Server{
		Addr:         web.address,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return server.ListenAndServe()
}
