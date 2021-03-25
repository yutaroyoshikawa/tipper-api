package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/yutaroyoshikawa/tipper-api/domain"
	"github.com/yutaroyoshikawa/tipper-api/graph/generated"
	"github.com/yutaroyoshikawa/tipper-api/graph/model"
)

func (r *mutationResolver) UpdateUser(ctx context.Context, userUID string, input model.UpdateUserInput) (*string, error) {
	err := r.Database.UpdateUser(input.ID, domain.UpdateUserInput{
		Email:   nil,
		IconUrl: input.ImageIcon,
		Name:    input.Name,
	})

	if err != nil {
		return nil, err
	}

	return &input.ID, nil
}

func (r *mutationResolver) UpdateUserID(ctx context.Context, input model.UpdateUserIDInput) (*string, error) {
	user, err := r.Database.GetUserByUserID(input.ID)

	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, errors.New("userID is exists")
	}

	err = r.Database.UpdateUserID(input.ID, input.NewID)

	if err != nil {
		return nil, err
	}

	return &input.NewID, nil
}

func (r *mutationResolver) CreatePerformance(ctx context.Context, input model.PerformanceInput) (*model.Performance, error) {
	performance := &model.Performance{
		ID:          "hoge",
		Name:        input.Name,
		Discription: input.Discription,
		Start:       input.Start,
		Finish:      input.Finish,
		Tags:        input.Tags,
		Thumbnail:   input.Thumbnail,
		Location:    (*model.Locate)(input.Location),
		Address:     input.Address,
		Artist: &model.User{
			ID:            "hoge",
			Name:          "hoge",
			ImageIcon:     "huga",
			FollowArtists: []string{},
			Performances:  []string{},
		},
	}
	return performance, nil
}

func (r *mutationResolver) UpdatePerformance(ctx context.Context, performanceID string, input model.PerformanceInput) (*model.Performance, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePerformance(ctx context.Context, input string) (*string, error) {
	if r.LoginUser == nil {
		return nil, errors.New("no permission")
	}

	performance := r.Database.GetPerformance(input)

	if performance.ArtistId != r.LoginUser.UID {
		return nil, errors.New("no permission")
	}

	res := r.Database.DeletePerformance(input)

	return &res, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user, err := r.Database.GetUserByUserID(id)

	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:            id,
		Name:          user.Name,
		ImageIcon:     user.IconUrl,
		FollowArtists: []string{},
		Performances:  []string{},
	}, nil
}

func (r *queryResolver) Performance(ctx context.Context, id string) (*model.Performance, error) {
	performance := r.Database.GetPerformance(id)
	artist := r.Database.GetUserByUID(performance.ArtistId)

	return &model.Performance{
		ID:          id,
		Name:        performance.Name,
		Discription: performance.Discription,
		Start:       performance.Start.String(),
		Finish:      performance.Finish.String(),
		Tags:        []string{},
		Thumbnail:   &performance.Thumbnail,
		Location: &model.Locate{
			Lat: performance.GeoLocate.GetLatitude(),
			Lng: performance.GeoLocate.GetLongitude(),
		},
		Address: performance.PlaceName,
		Artist: &model.User{
			ID:            artist.Id,
			Name:          artist.Name,
			ImageIcon:     artist.IconUrl,
			FollowArtists: []string{},
			Performances:  []string{},
		},
	}, nil
}

func (r *queryResolver) NearByPerformance(ctx context.Context, locate model.LocateInput) ([]*model.Performance, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
