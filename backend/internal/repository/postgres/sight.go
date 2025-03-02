// МЕТОДЫ ДЛЯ ОБРАЩЕНИЯ К БД с достопримечательностями (sight)
package repository

import (
	"context"
	"fmt"

	"homework_ipl/internal/entities"
	"homework_ipl/utils/logger"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"

	"github.com/georgysavva/scany/v2/pgxscan"
)

// Структура вызывальщика
type SightRepo struct {
	// технология пулов
	db *pgxpool.Pool
}

// Конструктор вызывальщика запросов в бд
func NewSightRepo(db *pgxpool.Pool) *SightRepo {
	return &SightRepo{
		db: db,
	}
}

// возвращает (четкие) поля ВСЕХ достопримечательностей (sight)
func (repo *SightRepo) GetSightsList() ([]entities.Sight, error) {
	// такая переменная создается везде - в нее будет записано через &
	var sights []*entities.Sight
	ctx := context.Background()

	// через & записываем результат
	err := pgxscan.Select(ctx, repo.db, &sights, `SELECT sight.id, rating, name, description, city_id, country_id, im.path FROM sight INNER JOIN image_data AS im ON sight.id = im.sight_id`)
	if err != nil {
		logger.Logger().Error(err.Error())
		return nil, err
	}

	var sightList []entities.Sight
	for _, s := range sights {
		sightList = append(sightList, *s)
	}
	return sightList, nil
}

// Возвращает данные ОДНОЙ достопримечательности по айди
func (repo *SightRepo) GetSightByID(id int) (entities.Sight, error) {
	// такая переменная создается везде - в нее будет записано через &
	var sight []*entities.Sight
	ctx := context.Background()

	err := pgxscan.Select(ctx, repo.db, &sight, `SELECT sight.id, rating, sight.name, description, city_id, sight.country_id, im.path, city.city, country.country, latitude, longitude FROM sight INNER JOIN image_data AS im ON sight.id = im.sight_id INNER JOIN city ON sight.city_id = city.id INNER JOIN country ON sight.country_id = country.id WHERE sight.id = $1`, id)
	if err != nil {
		logger.Logger().Error(err.Error())
		return entities.Sight{}, err
	}

	return *sight[0], nil
}

// Возвращает данные НЕСКОЛЬКИХ достопримечательностей, которые были отфильтрованы по ключу
// (в данном случае ключи: либо город, либо страна, но можно добавить еще кучу вариантов)
func (repo *SightRepo) GetFilteredSights(key string) ([]entities.Sight, error) {
	// такая переменная создается везде - в нее будет записано через &
	var sights []*entities.Sight
	ctx := context.Background()

	// если хочешь добавить еще ключи - OR поле IN (SELECT вт_поле FROM таблица WHERE условие)
	err := pgxscan.Select(ctx, repo.db, &sights, `SELECT sight.id, rating, name, description, city_id, country_id, im.path FROM sight INNER JOIN image_data AS im ON sight.id = im.sight_id
		WHERE city_id IN (SELECT id FROM city WHERE city = $1)
   		OR country_id IN (SELECT id FROM country WHERE country = $1)`, key)

	if err != nil {
		logger.Logger().Error(err.Error())
		return []entities.Sight{}, err
	}

	var sightList []entities.Sight
	for _, s := range sights {
		sightList = append(sightList, *s)
	}
	return sightList, nil
}

// Комментарии по айди
func (repo *SightRepo) GetCommentsBySightID(id int) ([]entities.Comment, error) {
	var comments []*entities.Comment
	ctx := context.Background()

	err := pgxscan.Select(ctx, repo.db, &comments, `SELECT f.id, f.user_id, p.username, p.avatar, f.sight_id, f.rating, f.feedback FROM feedback AS f INNER JOIN profile_data AS p ON f.user_id = p.user_id WHERE sight_id =  $1 `, id)
	if err != nil {
		logger.Logger().Error(err.Error())
		return nil, err
	}

	var commentsList []entities.Comment
	for _, s := range comments {
		commentsList = append(commentsList, *s)
	}

	return commentsList, nil
}

// Добавление комментария по айди поста (синоним достопримечательности)
func (repo *SightRepo) CreateCommentBySightID(dataStr map[string]string, dataInt map[string]int) error {
	ctx := context.Background()

	_, err := repo.db.Exec(ctx, `INSERT INTO feedback(user_id, sight_id, rating, feedback) VALUES($1, $2, $3, $4)`, dataInt["userID"], dataInt["sightID"], dataInt["rating"], dataStr["feedback"])
	if err != nil {
		logger.Logger().Error(err.Error())
		return err
	}

	return nil
}

// Редактирование комментария по айди поста
func (repo *SightRepo) EditCommentByCommentID(dataStr map[string]string, dataInt map[string]int) error {
	ctx := context.Background()

	_, err := repo.db.Exec(ctx, `UPDATE feedback SET rating = $1, feedback = $2 WHERE id = $3`, dataInt["rating"], dataStr["feedback"], dataInt["id"])
	if err != nil {
		logger.Logger().Error(err.Error())
		return err
	}

	return nil
}

// Удаление комментария по айди поста
func (repo *SightRepo) DeleteCommentByCommentID(dataInt map[string]int) error {
	ctx := context.Background()

	_, err := repo.db.Exec(ctx, `DELETE FROM feedback WHERE id = $1`, dataInt["id"])
	if err != nil {
		logger.Logger().Error(err.Error())
		return err
	}

	return nil
}

// Создание Поездки
func (repo *SightRepo) CreateJourney(dataInt map[string]int, dataStr map[string]string) (entities.Journey, error) {
	var journey entities.Journey
	ctx := context.Background()
	logrus.Info(dataStr["name"], dataInt["userID"], dataStr["description"])
	row := repo.db.QueryRow(ctx, `INSERT INTO journey(name, user_id, description) VALUES ($1, $2, $3) RETURNING id, name, user_id, description;`, dataStr["name"], dataInt["userID"], dataStr["description"])
	err := row.Scan(&journey.ID, &journey.Name, &journey.UserID, &journey.Description)
	res, _ := repo.GetJourneys(1)
	logrus.Info(res)
	if err != nil {
		logger.Logger().Error(err.Error())
		return entities.Journey{}, err
	}

	return journey, nil
}

// Удаление поездки по айди
func (repo *SightRepo) DeleteJourneyByID(dataInt map[string]int) error {
	ctx := context.Background()

	_, err := repo.db.Exec(ctx, `DELETE FROM journey_sight WHERE journey_id = $1`, dataInt["journeyID"])
	if err != nil {
		logger.Logger().Error(err.Error())
		return err
	}

	_, err = repo.db.Exec(ctx, `DELETE FROM journey WHERE id = $1`, dataInt["journeyID"])
	if err != nil {
		logger.Logger().Error(err.Error())
		return err
	}

	return nil
}

// Возвращает поездки по айди пользователя
func (repo *SightRepo) GetJourneys(userID int) ([]entities.Journey, error) {
	var journey []*entities.Journey
	ctx := context.Background()

	err := pgxscan.Select(ctx, repo.db, &journey, `SELECT j.id, j.name, j.description, p.username FROM journey AS j INNER JOIN profile_data AS p ON p.user_id = $1 WHERE j.user_id = $1`, userID)
	if err != nil {
		logger.Logger().Error(err.Error())
		return nil, err
	}

	var journeyList []entities.Journey
	for _, j := range journey {
		journeyList = append(journeyList, *j)
	}
	return journeyList, nil
}

// Добавить в поездку Достопримечательности (связующая таблица)
func (repo *SightRepo) AddJourneySight(dataInt map[string]int, ids []int, dataStr map[string]string) error {
	ctx := context.Background()
	precedence := 0
	// _, err := repo.db.Exec(ctx, `UPDATE journey SET name = $1, description = $2 WHERE id = $3;`, dataStr["name"], dataStr["description"], dataInt["journeyID"])
	// if err != nil {
	// 	logger.Logger().Error(err.Error())
	// 	return err
	// }
	// fmt.Println("Succes update")

	_, err := repo.db.Exec(ctx, `DELETE FROM journey_sight WHERE journey_sight.journey_id = $1;`, dataInt["journeyID"])
	if err != nil {
		logger.Logger().Error(err.Error())
		return err
	}

	for _, id := range ids {
		precedence += 1
		_, err := repo.db.Exec(ctx, `INSERT INTO journey_sight(journey_id, sight_id, priority) VALUES($1, $2, $3) `, dataInt["journeyID"], id, precedence)
		if err != nil {
			logger.Logger().Error(err.Error())
			return err
		}
	}

	return nil
}

// Удаление связующего
func (repo *SightRepo) DeleteJourneySight(dataInt map[string]int) error {
	ctx := context.Background()

	_, err := repo.db.Exec(ctx, `DELETE FROM journey_sight WHERE journey_id = $1 AND sight_id = $2 `, dataInt["journeyID"], dataInt["sightID"])
	if err != nil {
		logger.Logger().Error(err.Error())
		return err
	}

	return nil
}

// вернуть что то.. смотри sql запрос
func (repo *SightRepo) GetJourneySights(journeyID int) ([]entities.Sight, error) {
	var sights []entities.Sight
	var idList []*int
	ctx := context.Background()

	err := pgxscan.Select(ctx, repo.db, &idList, `SELECT js.sight_id FROM journey_sight AS js WHERE js.journey_id = $1`, journeyID)
	if err != nil {
		logger.Logger().Error(err.Error())
		return nil, err
	}

	for _, id := range idList {
		sight, err := repo.GetSightByID(*id)
		if err != nil {
			logger.Logger().Error(err.Error())
			continue
		}
		sights = append(sights, sight)
	}

	return sights, nil
}

// Возвращает поездку по айди поездки
func (repo *SightRepo) GetJourney(journeyID int) (entities.Journey, error) {
	var journey []*entities.Journey
	ctx := context.Background()

	err := pgxscan.Select(ctx, repo.db, &journey, `SELECT j.id, j.name, j.description, p.username, p.user_id FROM journey AS j INNER JOIN profile_data AS p ON p.user_id = j.user_id WHERE j.id = $1;`, journeyID)
	if err != nil {
		logger.Logger().Error(err.Error())
		return entities.Journey{}, err
	}

	fmt.Println(*journey[0])

	return *journey[0], nil
}
