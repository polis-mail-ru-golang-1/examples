package controller

import (
	"net/http"

	"github.com/polis-mail-ru-golang-1/examples/lec6/blog/model"
	"github.com/polis-mail-ru-golang-1/examples/lec6/blog/view"

	"github.com/rs/zerolog/log"
)

type Controller struct {
	view  view.View
	model model.Model
}

func New(v view.View, m model.Model) Controller {
	return Controller{
		view:  v,
		model: m,
	}
}

func (c Controller) Posts(w http.ResponseWriter, r *http.Request) {
	posts, err := c.model.Posts()
	if err != nil {
		c.error(w, r, err)
		return
	}
	c.view.Posts(posts, w)
}

func (c Controller) error(w http.ResponseWriter, r *http.Request, err error) {
	log.Error().Err(err).Msgf("Error %s", err)
	c.view.Error("server error", 500, w)
}
