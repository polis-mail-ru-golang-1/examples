package main

import (
	"github.com/polis-mail-ru-golang-1/examples/lec6/blog/config"
	"github.com/polis-mail-ru-golang-1/examples/lec6/blog/model"

	"github.com/go-pg/pg"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.Load()
	die(err)

	logLevel, err := zerolog.ParseLevel(cfg.LogLevel)
	if err != nil {
		panic(err)
	}
	zerolog.MessageFieldName = "msg"
	log.Level(logLevel)

	log.Print(cfg)

	pgOpt, err := pg.ParseURL(cfg.PgSQL)
	die(err)
	pgDb := pg.Connect(pgOpt)
	defer pgDb.Close()

	m := model.New(pgDb)

	test := model.Post{
		Title:   "title",
		Content: "content",
	}
	test, err = m.AddPost(test)
	die(err)
	log.Print(test)

	p, err := m.Posts()
	die(err)
	for _, post := range p {
		log.Print(post)
		c, err := m.Comments(post)
		die(err)
		for _, comment := range c {
			log.Print(comment)
		}
	}
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
