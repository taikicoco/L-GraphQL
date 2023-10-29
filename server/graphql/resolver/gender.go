package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

import (
	"context"
	"server/graphql/generated/model"
)

// Genders is the resolver for the genders field.
func (r *queryResolver) Genders(ctx context.Context) ([]*model.Gender, error) {
	res, err := r.gender.GetGenders(ctx)
	if err != nil {
		return nil, err
	}

	genders := make([]*model.Gender, len(res))
	for i, v := range res {
		genders[i] = &model.Gender{
			GenderID: v.GenderID,
			Gender:   v.Gender,
		}
	}

	return genders, nil
}

// Gender is the resolver for the gender field.
func (r *queryResolver) Gender(ctx context.Context, genderID int) (*model.Gender, error) {
	res, err := r.gender.GetGenderByID(ctx, genderID)
	if err != nil {
		return nil, err
	}

	gender := &model.Gender{
		GenderID: res.GenderID,
		Gender:   res.Gender,
	}

	return gender, nil
}
