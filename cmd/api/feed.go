package main

import "net/http"

func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	feed, err := app.store.Posts.GetUserFeed(ctx, int64(99))
	if err != nil {
		app.intetrnalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, feed); err != nil {
		app.intetrnalServerError(w, r, err)
		return
	}
}
