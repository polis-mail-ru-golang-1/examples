package view

import (
	"html/template"
	"io"

	"github.com/polis-mail-ru-golang-1/examples/lec6/blog/model"
)

type View struct {
	postsT *template.Template
	postT  *template.Template
	addT   *template.Template
	errorT *template.Template
}

func New() View {
	return View{
		postsT: template.Must(template.ParseFiles("view/posts.html")),
		postT:  template.Must(template.ParseFiles("view/post.html")),
		addT:   template.Must(template.ParseFiles("view/add.html")),
		errorT: template.Must(template.ParseFiles("view/error.html")),
	}
}

func (v View) Posts(posts []model.Post, wr io.Writer) {
	v.postsT.Execute(wr,
		struct {
			Posts []model.Post
		}{
			Posts: posts,
		})
}

func (v View) Post(post model.Post, comments []model.Comment, wr io.Writer) {
	v.postT.Execute(wr,
		struct {
			Post     model.Post
			Comments []model.Comment
		}{
			Post:     post,
			Comments: comments,
		})
}

func (v View) Error(err string, status int, wr io.Writer) {
	v.postT.Execute(wr,
		struct {
			Status int
			Error  string
		}{
			Status: status,
			Error:  err,
		})
}
