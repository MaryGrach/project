package delivery

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
	"homework_ipl/internal/entities"
	"homework_ipl/internal/http-server/server/db"
	sightRep "homework_ipl/internal/repository/postgres"
	"homework_ipl/utils/errors"
	"homework_ipl/utils/logger"
	"homework_ipl/utils/wrapper"
)

type JourneyHandler struct{}

type JourneyResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

var (
	errCreateJourney = errors.HttpError{
		Code:    http.StatusInternalServerError,
		Message: "failed creating new journey",
	}
	errDeleteJourney = errors.HttpError{
		Code:    http.StatusInternalServerError,
		Message: "failed deleting journey",
	}
	errAddJourneySight = errors.HttpError{
		Code:    http.StatusInternalServerError,
		Message: "failed adding journey sight",
	}
	errDeleteJourneySight = errors.HttpError{
		Code:    http.StatusInternalServerError,
		Message: "failed deleting journey sight",
	}
	errGetJourneySights = errors.HttpError{
		Code:    http.StatusInternalServerError,
		Message: "failed getting journey sight",
	}
	errInternal = errors.HttpError{
		Code:    http.StatusInternalServerError,
		Message: "internal Error",
	}
)

func (h JourneyHandler) CreateJourney(ctx context.Context, requestData entities.Journey) (entities.Journey, error) {
	db, err := db.GetPostgres()

	if err != nil {
		logger.Logger().Error(err.Error())
	}

	dataStr := make(map[string]string)
	dataInt := make(map[string]int)

	dataInt["userID"] = requestData.UserID
	dataStr["name"] = requestData.Name
	dataStr["description"] = requestData.Description

	fmt.Println("Create journey:", dataStr)

	sightsRepo := sightRep.NewSightRepo(db)
	journey, err := sightsRepo.CreateJourney(dataInt, dataStr)

	if err != nil {
		return entities.Journey{}, errCreateJourney
	}

	return journey, nil
}

func (h JourneyHandler) DeleteJourney(ctx context.Context, requestData entities.Journey) (entities.Journey, error) {
	db, err := db.GetPostgres()

	if err != nil {
		logger.Logger().Error(err.Error())
	}

	pathParams := wrapper.GetPathParamsFromCtx(ctx)
	journeyID, err := strconv.Atoi(pathParams["id"])
	if err != nil {
		logger.Logger().Error("Cannot convert string to integer to get sight")
		return entities.Journey{}, errParsing
	}

	dataInt := make(map[string]int)

	dataInt["journeyID"] = journeyID

	sightsRepo := sightRep.NewSightRepo(db)
	err = sightsRepo.DeleteJourneyByID(dataInt)

	if err != nil {
		return entities.Journey{}, errDeleteJourney
	}

	return entities.Journey{}, nil
}

func (h *JourneyHandler) GetJourneys(ctx context.Context, requestData entities.Journey) (entities.Journeys, error) {
	db, err := db.GetPostgres()
	if err != nil {
		logger.Logger().Error(err.Error())
	}

	pathParams := wrapper.GetPathParamsFromCtx(ctx)
	userID, err := strconv.Atoi(pathParams["userID"])
	if err != nil {
		logger.Logger().Error("Cannot convert string to integer to get sight")
		return entities.Journeys{}, errParsing
	}

	sightsRepo := sightRep.NewSightRepo(db)
	journeys, _ := sightsRepo.GetJourneys(userID)

	return entities.Journeys{Journey: journeys}, err
}

func (h *JourneyHandler) AddJourneySight(ctx context.Context, requestData entities.JourneySightID) (entities.JourneySight, error) {
	db, err := db.GetPostgres()
	if err != nil {
		logger.Logger().Error(err.Error())
	}
	pathParams := wrapper.GetPathParamsFromCtx(ctx)
	journeyID, err := strconv.Atoi(pathParams["id"])
	if err != nil {
		logger.Logger().Error("Cannot convert string to integer to get sight")
		return entities.JourneySight{}, errParsing
	}

	dataInt := make(map[string]int)
	dataInt["journeyID"] = journeyID

	dataStr := make(map[string]string)
	dataStr["name"] = requestData.Name
	dataStr["description"] = requestData.Description

	logrus.Info(dataInt)
	logrus.Info(requestData.ListID)
	logrus.Info(dataStr)
	sightsRepo := sightRep.NewSightRepo(db)
	err = sightsRepo.AddJourneySight(dataInt, requestData.ListID, dataStr)

	if err != nil {
		return entities.JourneySight{}, errAddJourneySight
	}

	return entities.JourneySight{JourneyID: journeyID}, nil
}

func (h *JourneyHandler) DeleteJourneySight(ctx context.Context, requestData entities.JourneySight) (entities.JourneySight, error) {
	db, err := db.GetPostgres()
	if err != nil {
		logger.Logger().Error(err.Error())
	}
	pathParams := wrapper.GetPathParamsFromCtx(ctx)
	journeyID, err := strconv.Atoi(pathParams["id"])
	if err != nil {
		logger.Logger().Error("Cannot convert string to integer to get sight")
		return entities.JourneySight{}, errParsing
	}

	dataInt := make(map[string]int)
	dataInt["journeyID"] = journeyID
	dataInt["sightID"] = requestData.SightID

	sightsRepo := sightRep.NewSightRepo(db)
	err = sightsRepo.DeleteJourneySight(dataInt)

	if err != nil {
		return entities.JourneySight{}, errDeleteJourneySight
	}

	return entities.JourneySight{}, nil
}

func (h *JourneyHandler) GetJourneySights(ctx context.Context, requestData entities.JourneySight) (entities.JourneySights, error) {
	db, err := db.GetPostgres()
	if err != nil {
		logger.Logger().Error(err.Error())
	}
	pathParams := wrapper.GetPathParamsFromCtx(ctx)
	journeyID, err := strconv.Atoi(pathParams["id"])
	if err != nil {
		logger.Logger().Error("Cannot convert string to integer to get sight")
		return entities.JourneySights{}, errParsing
	}

	sightsRepo := sightRep.NewSightRepo(db)
	sights, err := sightsRepo.GetJourneySights(journeyID)

	if err != nil {
		return entities.JourneySights{}, errGetJourneySights
	}

	journey, err := sightsRepo.GetJourney(journeyID)
	if err != nil {
		return entities.JourneySights{}, errGetJourneySights
	}

	return entities.JourneySights{Journey: journey, Sight: sights}, nil
}
