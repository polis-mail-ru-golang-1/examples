package controller

import (
	"net/http"
	"strconv"

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
		log.Error().Err(err).Msgf("Error loading posts %s", err)
		c.error(w, r, "server error", 500)
		return
	}
	c.view.Posts(posts, w)
}

func (c Controller) Post(w http.ResponseWriter, r *http.Request) {
	strid := r.URL.Query().Get("id")
	id, err := strconv.Atoi(strid)
	if err != nil {
		c.error(w, r, "Incorrect id "+strid, 400)
		return
	}
	post, err := c.model.Post(id)
	if err != nil {
		log.Error().Err(err).Msgf("Error loading post %d %s", id, err)
		c.error(w, r, "server error", 500)
		return
	}
	comments, err := c.model.Comments(post)
	if err != nil {
		log.Error().Err(err).Msgf("Error loading comments for post %d %s", id, err)
		c.error(w, r, "server error", 500)
		return
	}
	c.view.Post(post, comments, w)
}

func (c Controller) error(w http.ResponseWriter, r *http.Request, err string, status int) {
	log.Info().Msgf("Report error %s status %d", err, status)
	c.view.Error(err, status, w)
}
