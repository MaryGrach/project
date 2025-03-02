package delivery

import (
	"context"
	"strconv"

	"homework_ipl/internal/entities"
	"homework_ipl/internal/http-server/server/db"
	"homework_ipl/utils/logger"
	"homework_ipl/utils/wrapper"

	sightRep "homework_ipl/internal/repository/postgres"
)

type SightsHandler struct{}

type SightComments struct {
	Sight entities.Sight     `json:"sight"`
	Comms []entities.Comment `json:"comments"`
}

func (h *SightsHandler) GetSights(ctx context.Context, _ entities.Sight) (entities.Sights, error) {
	db, err := db.GetPostgres()

	if err != nil {
		logger.Logger().Error(err.Error())
	}
	sightsRepo := sightRep.NewSightRepo(db)
	sights, err := sightsRepo.GetSightsList()
	if err != nil {
		return entities.Sights{}, err
	}

	return entities.Sights{Sight: sights}, nil
}

func (h *SightsHandler) GetSightByID(ctx context.Context, requestData entities.Sight) (SightComments, error) {
	db, err := db.GetPostgres()
	if err != nil {
		logger.Logger().Error(err.Error())
	}

	pathParams := wrapper.GetPathParamsFromCtx(ctx)
	id, err := strconv.Atoi(pathParams["id"])
	if err != nil {
		logger.Logger().Error("Cannot convert string to integer to get sight")
		return SightComments{}, err
	}
	sightsRepo := sightRep.NewSightRepo(db)
	sight, _ := sightsRepo.GetSightByID(id)

	comments, err := sightsRepo.GetCommentsBySightID(id)

	return SightComments{Sight: sight, Comms: comments}, err
}

func (h *SightsHandler) GetFilteredSights(ctx context.Context, _ entities.Sight) (entities.Sights, error) {
	// Получаем соединение с базой данных
	db, err := db.GetPostgres()
	if err != nil {
		logger.Logger().Error("Failed to connect to database: " + err.Error())
		return entities.Sights{}, err
	}

	sightsRepo := sightRep.NewSightRepo(db)
	sights, err := sightsRepo.GetSightsList()
	if err != nil {
		return entities.Sights{}, err
	}

	queryRow := wrapper.GetQueryParamsFromCtx(ctx)
	name_of_city_of_country := queryRow["name"]

	if name_of_city_of_country == "" {
		return entities.Sights{Sight: sights}, nil
	}

	// Вызываем метод репозитория с фильтрами
	sights, err = sightsRepo.GetFilteredSights(name_of_city_of_country)
	if err != nil {
		logger.Logger().Error("Failed to fetch filtered sights: " + err.Error())
		return entities.Sights{}, err
	}

	return entities.Sights{Sight: sights}, nil
}
