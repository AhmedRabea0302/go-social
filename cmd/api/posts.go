package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/AhmedRabea0302/go-social/internal/store"
	"github.com/go-chi/chi/v5"
)

type CreatePostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=1000"`
	Tags    []string `json:"tags"`
}

func (app *application) CreatePostsHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
		// TODO: Change after auth
		UserID: 1,
	}

	ctx := r.Context()
	if err := app.store.Posts.Create(ctx, post); err != nil {
		app.intetrnalServerError(w, r, err)
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		app.intetrnalServerError(w, r, err)
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	idpParam := chi.URLParam(r, "postID")
	postID, err := strconv.ParseInt(idpParam, 10, 64)
	if err != nil {
		app.intetrnalServerError(w, r, err)
		return
	}

	ctx := r.Context()
	post, err := app.store.Posts.GetPostByID(ctx, postID)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrorNotFound):
			app.notFoundResponse(w, r, err)
		default:
			app.intetrnalServerError(w, r, err)
		}
		return
	}

	comments, err := app.store.Comments.GetCommentsByPostID(ctx, postID)
	if err != nil {
		app.intetrnalServerError(w, r, err)
		return
	}

	post.Comments = comments

	if err := writeJSON(w, http.StatusOK, post); err != nil {
		app.intetrnalServerError(w, r, err)
		return
	}
}
