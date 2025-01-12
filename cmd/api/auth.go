package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/AhmedRabea0302/go-social/internal/mailer"
	"github.com/AhmedRabea0302/go-social/internal/store"
	"github.com/google/uuid"
)

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

type userWithToken struct {
	*store.User
	Token string `json:"token"`
}

// RegisterUser gdoc
//
//	@Summary		Registers a user
//	@Description	Register a user by ID
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		RegisterUserPayload	true	"User Credentials"
//	@Success		201		{object}	userWithToken		"User registered"
//	@Failure		400		{object}	error				"User payload missing"
//	@Failure		500		{object}	error				"User not found"
//	@Security		ApiKeyAuth
//	@Router			/authentication/user [post]
func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &store.User{
		Username: payload.Username,
		Email:    payload.Email,
	}

	// Hash the user password
	if err := user.Password.Set(payload.Password); err != nil {
		app.intetrnalServerError(w, r, err)
		return
	}

	// Create and hash token
	plainToken := uuid.New().String()
	hash := sha256.Sum256([]byte(plainToken))
	hashedToken := hex.EncodeToString(hash[:])

	// Store the user
	err := app.store.Users.CreateAndInvite(r.Context(), user, hashedToken, app.config.mail.expiry)
	if err != nil {
		switch err {
		case store.ErrDuplicateEmail:
			app.badRequestResponse(w, r, err)
			return
		case store.ErrDuplicateUsername:
			app.badRequestResponse(w, r, err)
			return
		default:
			app.intetrnalServerError(w, r, err)
			return
		}
	}

	userWithToken := userWithToken{
		User:  user,
		Token: plainToken,
	}
	activationURL := fmt.Sprintf("%s/confirm/%s", app.config.frontendURL, plainToken)

	// send the mail
	isProdEnv := app.config.env == "production"
	vars := struct {
		Username      string
		ActivationURL string
	}{
		Username:      user.Username,
		ActivationURL: activationURL,
	}
	// send email
	status, err := app.mailer.Send(mailer.UserWelcomeTemplate, user.Username, user.Email, vars, !isProdEnv)
	if err != nil {
		app.logger.Errorw("error sending welcome message email", "error", err)

		// rollback user creation if email fails
		ctx := r.Context()
		if err = app.store.Users.Delete(ctx, user.ID); err != nil {
			app.logger.Errorw("error deleting user", "error", err)
			return
		}

		app.intetrnalServerError(w, r, err)
		return
	}

	app.logger.Infow("Email sent", "status code", status)

	if err := app.jsonResponse(w, http.StatusCreated, userWithToken); err != nil {
		app.intetrnalServerError(w, r, err)
		return
	}
}
