package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"tages/internal/server/models"
)

type MediaRepo struct {
	dbConn *pgxpool.Pool
}

func NewMediaRepo(conn *pgxpool.Pool) *MediaRepo {
	return &MediaRepo{
		dbConn: conn,
	}
}

const insertNewFile = `INSERT INTO media (name) VALUES ($1)`

func (mr *MediaRepo) Save(ctx context.Context, filePath string) (bool, error) {
	if _, err := mr.dbConn.Exec(ctx, insertNewFile, filePath); err != nil {
		return false, err
	}

	return true, nil
}

const getAllMedia = `SELECT * FROM media`

func (mr *MediaRepo) GetAll(ctx context.Context) ([]models.Media, error) {
	var medias []models.Media
	rows, err := mr.dbConn.Query(ctx, getAllMedia)
	if err != nil {
		return medias, err
	}

	for rows.Next() {
		var media models.Media
		if err = rows.Scan(&media.ID, &media.Name, &media.CreatedAt, &media.UpdatedAt); err != nil {
			return nil, err
		}

		medias = append(medias, media)
	}

	return medias, nil
}

const getMedia = `SELECT * FROM media WHERE name LIKE $1`

func (mr *MediaRepo) GetOne(ctx context.Context, name string) (models.Media, error) {
	var media models.Media

	if err := mr.dbConn.QueryRow(ctx, getMedia, "%"+name+"%").Scan(&media.ID, &media.Name, &media.CreatedAt, &media.UpdatedAt); err != nil {
		return media, err
	}

	return media, nil
}
