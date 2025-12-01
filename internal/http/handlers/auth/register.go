package auth

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/bulgil/im-proc-svc/internal/domain/user"
	"github.com/bulgil/im-proc-svc/internal/http/handlers"
	"github.com/bulgil/im-proc-svc/internal/middleware"
	"github.com/bulgil/im-proc-svc/internal/validator"
)

type RegisterInput struct {
	Username string `json:"username" validate:"alphanum,min=5,max=35"`
	Password string `json:"password" validate:"min=7,max=50"`
}

func Register(userRepo user.Repository, val *validator.Validator, log *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input RegisterInput

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			handlers.JSONInternalErrorResponse(w)
			log.Error("a problem with decode body", "error", err.Error(), "request_id", middleware.GetRequestID(r.Context()))
			return
		}

		if err := val.Validate(input); err != nil {
			handlers.JSONErrorResponse(w, handlers.JSONError{
				Err:            "validation error",
				ErrDescription: err.Error(),
			}, http.StatusBadRequest)
			return
		}

		exists, err := userRepo.CheckUsername(r.Context(), input.Username)
		if err != nil {
			handlers.JSONInternalErrorResponse(w)
			log.Error("a problem with check username", "error", err.Error(), "request_id", middleware.GetRequestID(r.Context()))
			return
		}
		if exists {
			handlers.JSONErrorResponse(w, handlers.JSONError{
				Err:            "bad username",
				ErrDescription: "user with such username already exists",
			}, http.StatusUnprocessableEntity)
			return
		}

		user := user.User{
			Username: input.Username,
			Password: []byte(input.Password),
		}
		if err := user.HashPassword(); err != nil {
			handlers.JSONInternalErrorResponse(w)
			log.Error("a problem with hash password", "error", err.Error(), "request_id", middleware.GetRequestID(r.Context()))
			return
		}

		if err := userRepo.Create(r.Context(), &user); err != nil {
			handlers.JSONInternalErrorResponse(w)
			log.Error("a problem with register user", "error", err.Error(), "request_id", middleware.GetRequestID(r.Context()))
			return
		}

		handlers.JSONResponse(w, map[string]any{
			"user": user,
		}, http.StatusOK)
	})
}
