package repository

import (
	"context"
	"server/model"

	"github.com/jmoiron/sqlx"
)

type PrefectureRepository struct{}

func (pr *PrefectureRepository) GetAll(ctx context.Context, db *sqlx.DB) ([]*model.Prefecture, error) {
	prefectures := []*model.Prefecture{}
	err := db.SelectContext(ctx, &prefectures, "SELECT prefecture_id, name FROM prefectures")

	if err != nil {
		return nil, err
	}
	return prefectures, nil
}

func (pr *PrefectureRepository) GetByID(ctx context.Context, db *sqlx.DB, prefectureID int) (*model.Prefecture, error) {
	prefecture := &model.Prefecture{}
	err := db.Get(prefecture, `select prefecture_id, name from prefectures where prefecture_id = ?`, prefectureID)
	if err != nil {
		return nil, err
	}
	return prefecture, nil
}

func (r *PrefectureRepository) GetByAnimeID(ctx context.Context, db *sqlx.DB, animeID int) ([]*model.Prefecture, error) {
	prefecture := []*model.Prefecture{}
	err := db.Select(&prefecture,
		`SELECT  DISTINCT sa.prefecture_id, sa.name
		FROM prefectures sa
		INNER JOIN spots s ON sa.prefecture_id = s.prefecture_id
		WHERE s.anime_id = ?`, animeID)
	if err != nil {
		return nil, err
	}
	return prefecture, nil
}