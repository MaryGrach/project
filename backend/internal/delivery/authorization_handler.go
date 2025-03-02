package delivery

import (
	"context"
	"net/http"

	"homework_ipl/internal/entities"
	"homework_ipl/internal/http-server/server/db"
	userRep "homework_ipl/internal/repository/postgres"
	"homework_ipl/internal/usecase"
	"homework_ipl/utils/errors"
	"homework_ipl/utils/httputils"
	"homework_ipl/utils/logger"
)

type AuthorizationHandler struct{}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

var (
	errLoginUser = errors.HttpError{
		Code:    http.StatusBadRequest,
		Message: "failed authorize",
	}
	errSetSession = errors.HttpError{
		Code:    http.StatusInternalServerError,
		Message: "failed setting session",
	}
	errClearSession = errors.HttpError{
		Code:    http.StatusInternalServerError,
		Message: "failed clearing session",
	}
	errSessionNotSet = errors.HttpError{
		Code:    http.StatusUnauthorized,
		Message: "session is not set",
	}
)

// Хэндлер авторизации
func (h *AuthorizationHandler) Authorize(ctx context.Context, requestData entities.User) (UserResponse, error) {
	username := requestData.Email
	password := requestData.Passwrd

	db, err := db.GetPostgres()
	if err != nil {
		logger.Logger().Error(err.Error())
	}

	dataStr := make(map[string]string)

	dataStr["email"] = username
	dataStr["passwrd"] = password

	UserRepo := userRep.NewUserRepo(db)
	user, err := UserRepo.AuthorizeUser(dataStr)
	if err != nil {
		return UserResponse{}, errLoginUser
	}

	responseWriter, ok := httputils.ContextWriter(ctx)
	if !ok {
		return UserResponse{}, errInternal
	}

	err = usecase.SetSession(responseWriter, user.ID)
	if err != nil {
		return UserResponse{}, errSetSession
	}

	userResponse := UserResponse{
		ID:       user.ID,
		Username: user.Email,
	}

	return userResponse, nil
}

// Выход
func (h *AuthorizationHandler) LogOut(ctx context.Context, requestData entities.User) (UserResponse, error) {
	request, ok := httputils.HttpRequest(ctx)
	if !ok {
		return UserResponse{}, errInternal
	}

	responseWriter, ok := httputils.ContextWriter(ctx)
	if !ok {
		return UserResponse{}, errInternal
	}

	userID := usecase.GetSession(request)

	if userID == 0 {
		return UserResponse{}, errSessionNotSet
	}

	err := usecase.ClearSession(responseWriter, request)
	if err != nil {
		return UserResponse{}, errClearSession
	}

	return UserResponse{}, nil
}
