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

type RegistrationHandler struct{}

var (
	errCreateUser = errors.HttpError{
		Code:    http.StatusInternalServerError,
		Message: "failed creating new profile",
	}
)

func (h *RegistrationHandler) SignUp(ctx context.Context, requestData entities.User) (UserResponse, error) {
	db, err := db.GetPostgres()
	if err != nil {
		logger.Logger().Error(err.Error())
	}

	username := requestData.Email
	password := requestData.Passwrd

	if err := entities.UserDataVerification(username, password); err != nil {
		return UserResponse{}, errors.HttpError{Code: http.StatusBadRequest, Message: err.Error()}
	}

	user, err := entities.CreateUser(username, password)
	if err != nil {
		return UserResponse{}, errCreateUser
	}

	responseWriter, ok := httputils.ContextWriter(ctx)
	if !ok {
		return UserResponse{}, errInternal
	}

	err = usecase.SetSession(responseWriter, user.ID)
	if err != nil {
		return UserResponse{}, errSetSession
	}

	dataStr := make(map[string]string)

	dataStr["email"] = user.Email
	dataStr["passwrd"] = user.Passwrd

	UserRepo := userRep.NewUserRepo(db)
	user, err = UserRepo.CreateUser(dataStr)
	if err != nil {
		return UserResponse{}, errCreateUser
	}

	err = usecase.SetSession(responseWriter, user.ID)
	if err != nil {
		return UserResponse{}, errSetSession
	}

	return UserResponse{ID: user.ID, Username: user.Email}, nil
}
