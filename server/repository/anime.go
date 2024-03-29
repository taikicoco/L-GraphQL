package repository

import (
	"context"
	"server/model"

	"github.com/jmoiron/sqlx"
)

type AnimeRepository struct{}

func (ar *AnimeRepository) GetAll(ctx context.Context, db *sqlx.DB) ([]*model.Anime, error) {
	animes := []*model.Anime{}
	err := db.SelectContext(ctx, &animes, `SELECT anime_id, name, img_url FROM animes`)

	if err != nil {
		return nil, err
	}
	return animes, nil
}

func (ar *AnimeRepository) GetByID(ctx context.Context, db *sqlx.DB, animeID int) (*model.Anime, error) {
	anime := &model.Anime{}
	err := db.Get(anime, `select anime_id, name, img_url from animes where anime_id = ?`, animeID)
	if err != nil {
		return nil, err
	}
	return anime, nil
}

func (ar *AnimeRepository) GetByIDs(ctx context.Context, db *sqlx.DB, animeIDs []int) ([]*model.Anime, error) {
	animes := []*model.Anime{}

	query := `
	select anime_id, name, img_url from animes where anime_id in (?)
	`
	query, args, err := sqlx.In(query, animeIDs)
	if err != nil {
		return nil, err
	}

	query = db.Rebind(query)

	err = db.Select(&animes, query, args...)
	if err != nil {
		return nil, err
	}
	return animes, nil
}
