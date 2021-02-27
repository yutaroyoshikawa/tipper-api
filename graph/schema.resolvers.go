package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/yutaroyoshikawa/tipper-api/graph/generated"
	"github.com/yutaroyoshikawa/tipper-api/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		ID:            "hoge",
		UserType:      model.UserTypeUser,
		Name:          input.Name,
		ImageIcon:     *input.ImageIcon,
		FollowArtists: []string{},
		Performances:  []string{},
	}
	return user, nil
}

func (r *mutationResolver) CreatePerformance(ctx context.Context, input *model.NewPerformance) (*model.Performance, error) {
	performance := &model.Performance{
		ID:          "hoge",
		Name:        input.Name,
		Description: input.Description,
		Start:       input.Start,
		Finish:      input.Finish,
		Tags:        input.Tags,
		Thumbnail:   input.Thumbnail,
		Location:    (*model.Locate)(input.Location),
		Address:     input.Address,
		Artist: &model.User{
			ID:            "hoge",
			UserType:      model.UserTypeUser,
			Name:          "hoge",
			ImageIcon:     "huga",
			FollowArtists: []string{},
			Performances:  []string{},
		},
		Comments: []*model.Comment{},
	}
	return performance, nil
}

func (r *queryResolver) User(ctx context.Context, id *string) (*model.User, error) {
	user := &model.User{
		ID:            *id,
		UserType:      model.UserTypeUser,
		Name:          "hoge",
		ImageIcon:     "huga",
		FollowArtists: []string{},
		Performances:  []string{},
	}
	return user, nil
}

func (r *queryResolver) Performance(ctx context.Context, id *string) (*model.Performance, error) {
	thumbnail := "hoge"

	performance := &model.Performance{
		ID:          *id,
		Name:        "hoge",
		Description: "huga",
		Start:       "hoge",
		Finish:      "huga",
		Tags:        []string{},
		Thumbnail:   &thumbnail,
		Location: &model.Locate{
			Lat: 0.0,
			Lng: 0.0,
		},
		Address: "hoge",
		Artist: &model.User{
			ID:            *id,
			UserType:      model.UserTypeUser,
			Name:          "hoge",
			ImageIcon:     "huga",
			FollowArtists: []string{},
			Performances:  []string{},
		},
		Comments: []*model.Comment{},
	}
	return performance, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
