package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

import (
	"context"
	"server/graphql/generated"
	"server/graphql/generated/model"
)

// Gender is the resolver for the gender field.
func (r *guideResolver) Gender(ctx context.Context, obj *model.Guide) (*model.Gender, error) {
	res, err := r.gender.GetGenderByGuideID(ctx, obj.GuideID)
	if err != nil {
		return nil, err
	}

	gender := &model.Gender{
		GenderID: res.GenderID,
		Gender:   res.Gender,
	}

	return gender, nil
}

// Country is the resolver for the country field.
func (r *guideResolver) Country(ctx context.Context, obj *model.Guide) (*model.Country, error) {
	res, err := r.country.GetCountryByGuideID(ctx, obj.GuideID)
	if err != nil {
		return nil, err
	}

	country := &model.Country{
		CountryID: res.CountryID,
		Name:      res.Name,
		ImgURL:    &res.ImgURL,
	}
	return country, nil
}

// Guides is the resolver for the guides field.
func (r *queryResolver) Guides(ctx context.Context) ([]*model.Guide, error) {
	res, err := r.guide.GetGuides(ctx)
	if err != nil {
		return nil, err
	}

	guides := make([]*model.Guide, len(res))
	for i, v := range res {
		guides[i] = &model.Guide{
			GuideID:           v.GuideID,
			Name:              v.Name,
			Age:               v.Age,
			Comment:           &v.Comment,
			Stance:            &v.Stance,
			FavoriteCharacter: &v.FavoriteCharacter,
		}
	}
	return guides, nil
}

// Guide is the resolver for the guide field.
func (r *queryResolver) Guide(ctx context.Context, guideID int) (*model.Guide, error) {
	res, err := r.guide.GetGuideByID(ctx, guideID)
	if err != nil {
		return nil, err
	}

	guide := &model.Guide{
		GuideID:           res.GuideID,
		Name:              res.Name,
		Age:               res.Age,
		Comment:           &res.Comment,
		Stance:            &res.Stance,
		FavoriteCharacter: &res.FavoriteCharacter,
	}
	return guide, nil
}

// GuideByAnimeIdsAndPrefectureIds is the resolver for the guideByAnimeIdsAndPrefectureIds field.
func (r *queryResolver) GuideByAnimeIdsAndPrefectureIds(ctx context.Context, animeIds []int, prefectureIds []int) ([]*model.Guide, error) {
	res, err := r.guide.GetGuidesByAnimeIDsAndPrefectureIDs(ctx, animeIds, prefectureIds)
	if err != nil {
		return nil, err
	}
	animes := make([]*model.Guide, len(res))
	for i, v := range res {
		animes[i] = &model.Guide{
			GuideID:           v.GuideID,
			Name:              v.Name,
			Age:               v.Age,
			Comment:           &v.Comment,
			Stance:            &v.Stance,
			FavoriteCharacter: &v.FavoriteCharacter,
		}
	}
	return animes, nil
}

// Guide returns generated.GuideResolver implementation.
func (r *Resolver) Guide() generated.GuideResolver { return &guideResolver{r} }

type guideResolver struct{ *Resolver }
