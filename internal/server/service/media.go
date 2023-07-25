package service

import (
	"context"
	"tages/internal/server/models"
	"tages/internal/server/repository"
)

type MediaInf interface {
	Upload(ctx context.Context, file string) (bool, error)
	List(ctx context.Context) ([]models.Media, error)
	GetMedia(ctx context.Context, name string) (models.Media, error)
}

type MediaSvc struct {
	repo *repository.MediaRepo
}

func NewMediaSvc(repo *repository.MediaRepo) MediaSvc {
	return MediaSvc{
		repo: repo,
	}
}

func (m *MediaSvc) Upload(file string) (bool, error) {
	saved, err := m.repo.Save(context.TODO(), file)
	if err != nil {
		return saved, err
	}
	return saved, nil
}

func (m *MediaSvc) List(ctx context.Context) ([]models.Media, error) {
	medias, err := m.repo.GetAll(ctx)
	if err != nil {
		return medias, err
	}
	return medias, nil
}

func (m *MediaSvc) GetMedia(ctx context.Context, name string) (models.Media, error) {
	media, err := m.repo.GetOne(ctx, name)
	if err != nil {
		return media, err
	}
	return media, nil
}
