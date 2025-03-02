package delivery

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"homework_ipl/internal/entities"
	"homework_ipl/internal/http-server/server/db"
	userRep "homework_ipl/internal/repository/postgres"
	"homework_ipl/internal/usecase"
	"homework_ipl/utils/errors"
	"homework_ipl/utils/httputils"
	"homework_ipl/utils/logger"
	"homework_ipl/utils/wrapper"

	"golang.org/x/crypto/bcrypt"
)

type ProfileHandler struct{}

type ProfileResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Avatar   string `json:"avatar"`
}

var (
	errEditProfile = errors.HttpError{
		Code:    http.StatusInternalServerError,
		Message: "failed edit profile",
	}
	errDeleteProfile = errors.HttpError{
		Code:    http.StatusInternalServerError,
		Message: "failed deleting profile",
	}
	errProfilePermissionDenied = errors.HttpError{
		Code:    http.StatusUnauthorized,
		Message: "permission denied",
	}

	errProfileResetPassword = errors.HttpError{
		Code:    http.StatusUnauthorized,
		Message: "weak password",
	}

	errIncorrectOldPassword = errors.HttpError{
		Code:    http.StatusUnauthorized,
		Message: "incorrect old password",
	}
)

func (h *ProfileHandler) GetUserProfile(ctx context.Context, requestData entities.User) (ProfileResponse, error) {
	db, err := db.GetPostgres()
	if err != nil {
		logger.Logger().Error(err.Error())
	}

	pathParams := wrapper.GetPathParamsFromCtx(ctx)
	id, err := strconv.Atoi(pathParams["id"])
	if err != nil {
		logger.Logger().Error("Cannot convert string to integer to get sight")
		return ProfileResponse{}, err
	}

	dataInt := make(map[string]int)

	dataInt["userID"] = id

	UserRepo := userRep.NewUserRepo(db)
	user, err := UserRepo.GetUserProfile(dataInt)
	if err != nil {
		return ProfileResponse{}, errLoginUser
	}

	profileResponse := ProfileResponse{
		ID:       user.UserID,
		Username: user.Username,
		Bio:      user.Bio,
		Avatar:   user.Avatar,
	}

	return profileResponse, nil
}

func (h *ProfileHandler) DeleteUser(ctx context.Context, requestData entities.User) (ProfileResponse, error) {
	logger := logger.Logger()
	db, err := db.GetPostgres()

	if err != nil {
		logger.Error(err.Error())
	}

	pathParams := wrapper.GetPathParamsFromCtx(ctx)
	userID, err := strconv.Atoi(pathParams["id"])
	if err != nil {
		logger.Error("Cannot convert string to integer to get sight")
		return ProfileResponse{}, errParsing
	}

	if r, _ := httputils.HttpRequest(ctx); userID != usecase.GetSession(r) {
		logger.Error("Cannot edit other's profile")
		return ProfileResponse{}, errProfilePermissionDenied
	}

	dataInt := make(map[string]int)

	dataInt["userID"] = userID

	userRepo := userRep.NewUserRepo(db)
	err = userRepo.DeleteUserProfile(dataInt)

	if err != nil {
		return ProfileResponse{}, errDeleteProfile
	}

	return ProfileResponse{}, nil
}

func (h *ProfileHandler) EditUserProfile(ctx context.Context, requestData entities.UserProfile) (entities.UserProfile, error) {
	logger := logger.Logger()
	db, err := db.GetPostgres()

	if err != nil {
		logger.Error(err.Error())
	}

	pathParams := wrapper.GetPathParamsFromCtx(ctx)
	userID, err := strconv.Atoi(pathParams["id"])
	if err != nil {
		logger.Error("Error while converting string to int", "error", err)
		return entities.UserProfile{}, errParsing
	}
	r, _ := httputils.HttpRequest(ctx)
	if userID != usecase.GetSession(r) {
		logger.Error("Cannot edit other's profile")
		return entities.UserProfile{}, errProfilePermissionDenied
	}

	dataInt := make(map[string]int)
	dataStr := make(map[string]string)

	dataInt["userID"] = userID
	dataStr["username"] = requestData.Username
	dataStr["bio"] = requestData.Bio

	userRepo := userRep.NewUserRepo(db)
	profile, err := userRepo.EditUserProfile(dataInt, dataStr)

	if err != nil {
		return entities.UserProfile{}, errEditProfile
	}

	return profile, nil
}

func (h *ProfileHandler) UpdateUserPassword(ctx context.Context, requestData entities.Password) (ProfileResponse, error) {
	logger := logger.Logger()

	db, err := db.GetPostgres()
	if err != nil {
		logger.Error("Error while connecting to db", "error", err)
	}

	pathParams := wrapper.GetPathParamsFromCtx(ctx)
	userID, err := strconv.Atoi(pathParams["id"])
	if err != nil {
		logger.Error("Error while converting string to int", "error", err)
		return ProfileResponse{}, errParsing
	}
	r, _ := httputils.HttpRequest(ctx)
	if userID != usecase.GetSession(r) {
		logger.Error("Cannot edit other's profile")
		return ProfileResponse{}, errProfilePermissionDenied
	}

	userRepo := userRep.NewUserRepo(db)
	// ПРОВЕРКА СТАРОГО ПАРОЛЯ ---------

	old_passwrd_hash, err := userRepo.GetHashPassword(userID)

	if err != nil {
		return ProfileResponse{}, errLoginUser
	}

	err = bcrypt.CompareHashAndPassword([]byte(old_passwrd_hash), []byte(requestData.Passwrd))

	if err != nil {
		fmt.Println("Passwords not match!")
		return ProfileResponse{}, errIncorrectOldPassword
	}
	// --------------------------------

	if !entities.ValidatePassword(requestData.NewPasswrd) {
		return ProfileResponse{}, errProfileResetPassword
	}

	err = userRepo.UpdateUserPassword(userID, requestData.NewPasswrd)
	if err != nil {
		return ProfileResponse{}, err
	}

	return ProfileResponse{}, nil
}

func (f *ProfileHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	logger := logger.Logger()
	// Подключение к базе данных
	db, err := db.GetPostgres()
	if err != nil {
		logger.Error("Ошибка подключения к базе данных:", "error", err)
		errors.WriteHttpError(err, w)
		return
	}
	// Получение userID из параметров
	id := wrapper.GetPathParams(r)["id"]
	userID, err := strconv.Atoi(id)
	if err != nil {
		logger.Error("Ошибка получения userID из параметров:", "error", err)
		errors.WriteHttpError(err, w)
		return
	}
	// Проверка прав пользователя
	if userID != usecase.GetSession(r) {
		logger.Error("Пользователь пытается изменить чужой профиль:", "userID", userID)
		errors.WriteHttpError(errProfilePermissionDenied, w)
		return
	}
	// Сохранение файла
	logger.Info("Начата загрузка файла для пользователя:", "userID", userID)
	path, err := SaveFile(r, id)
	if err != nil {
		logger.Error("Ошибка при сохранении файла:", "error", err)
		errors.WriteHttpError(err, w)
		return
	}
	logger.Info("Файл успешно сохранен:", "path", path)
	// Сохранение имени файла в базе данных
	fileName := filepath.Base(path)
	logger.Info("Сохранение имени файла в БД:", "fileName", fileName)
	dataInt := map[string]int{"userID": userID}
	dataStr := map[string]string{"avatar": "/public/avatars/" + fileName}
	userRepo := userRep.NewUserRepo(db)
	profile, err := userRepo.EditUserProfile(dataInt, dataStr)
	if err != nil {
		logger.Error("Ошибка при обновлении профиля в БД:", "error", err)
		errors.WriteHttpError(err, w)
		return
	}
	logger.Info("Профиль успешно обновлен в БД для пользователя:", "userID", userID)
	// Формирование JSON-ответа
	rawJSON, err := json.Marshal(profile)
	if err != nil {
		logger.Error("Ошибка при кодировании профиля в JSON:", "error", err)
		errors.WriteHttpError(err, w)
		return
	}
	logger.Info("JSON-ответ сформирован успешно")
	// Отправка ответа
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(rawJSON)
	if err != nil {
		logger.Error("Ошибка при отправке JSON-ответа:", "error", err)
		return
	}
	logger.Info("JSON-ответ успешно отправлен")
}
