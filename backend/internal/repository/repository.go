package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"homework_ipl/internal/entities"
	su "homework_ipl/internal/repository/postgres"
)

type SightRepositoryI interface {
	NewSightRepo(db *pgxpool.Pool) *su.SightRepo
	GetSightsList() ([]entities.Sight, error)
	GetSightByID(id int) (entities.Sight, error)
	GetCommentsBySightID(id int) ([]entities.Comment, error)
}
