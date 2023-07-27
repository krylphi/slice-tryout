package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Repo interface {
	Save(context.Context, string) (string, error)
	Load(context.Context, string) (string, error)
}

type Handler struct {
	srv  *http.Server
	repo Repo
}

func NewHandler(repo Repo, addr, port string) Handler {
	router := gin.Default()

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", addr, port),
		Handler: router,
	}

	h := Handler{
		srv:  srv,
		repo: repo,
	}

	router.POST("/encode", h.encodeHandler)
	router.POST("/decode", h.decodeHandler)

	return h
}

func (h Handler) Start() error {
	return h.srv.ListenAndServe()
}

func (h Handler) Shutdown(ctx context.Context) error {
	return h.srv.Shutdown(ctx)
}
