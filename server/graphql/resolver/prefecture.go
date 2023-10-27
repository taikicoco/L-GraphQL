package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

import (
	"context"
	"server/graphql/generated"
	"server/graphql/generated/model"
)

// Spot is the resolver for the spot field.
func (r *prefectureResolver) Spot(ctx context.Context, obj *model.Prefecture) ([]*model.Spot, error) {
	res, err := r.spot.GetSpotByPrefectureID(ctx, *obj.AnimeID, obj.PrefectureID)
	if err != nil {
		return nil, err
	}

	spot := make([]*model.Spot, len(res))
	for i, v := range res {
		spot[i] = &model.Spot{
			SpotID:      v.SpotID,
			Name:        v.Name,
			AnimeImgURL: &v.AnimeImgURL,
			RealImgURL:  &v.RealImgURL,
			Address:     &v.Address,
		}
	}
	return spot, nil
}

// Prefecture is the resolver for the prefecture field.
func (r *queryResolver) Prefecture(ctx context.Context, prefectureID int) (*model.Prefecture, error) {
	res, err := r.prefecture.GetPrefectureByID(ctx, prefectureID)
	if err != nil {
		return nil, err
	}

	prefecture := &model.Prefecture{
		PrefectureID: res.PrefectureID,
		Name:         res.Name,
	}
	return prefecture, nil
}

// Prefectures is the resolver for the prefectures field.
func (r *queryResolver) Prefectures(ctx context.Context) ([]*model.Prefecture, error) {
	res, err := r.prefecture.PrefectureList(ctx)
	if err != nil {
		return nil, err
	}

	prefectures := make([]*model.Prefecture, len(res))
	for i, v := range res {
		prefectures[i] = &model.Prefecture{
			PrefectureID: v.PrefectureID,
			Name:         v.Name,
		}
	}
	return prefectures, nil
}

// Prefecture returns generated.PrefectureResolver implementation.
func (r *Resolver) Prefecture() generated.PrefectureResolver { return &prefectureResolver{r} }

type prefectureResolver struct{ *Resolver }
