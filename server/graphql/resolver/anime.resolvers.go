package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

import (
	"context"
	"fmt"
	"server/graphql/generated/model"
)

// AnimeListByAnimeID is the resolver for the animeListByAnimeId field.
func (r *queryResolver) AnimeListByAnimeID(ctx context.Context, animeID []*int) ([]*model.Anime, error) {
	panic(fmt.Errorf("not implemented: AnimeListByAnimeID - animeListByAnimeId"))
}

// AnimeList is the resolver for the animeList field.
func (r *queryResolver) AnimeList(ctx context.Context) ([]*model.Anime, error) {
	animes := []*model.Anime{
		{
			AnimeID: 1,
			Name:    "anime1",
		},
		{
			AnimeID: 2,
			Name:    "anime2",
		},
	}
	return animes, nil
}
