package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/jasperges/go-turbine-api/graph/generated"
	"github.com/jasperges/go-turbine-api/graph/model"
)

func (r *mutationResolver) CreateFrameRate(ctx context.Context, input model.NewFrameRate) (*model.FrameRate, error) {
	frameRate := &model.FrameRate{
		ID:           fmt.Sprintf("T%d", rand.Int()),
		FrameRateNum: input.FrameRateNum,
		FrameRateDen: input.FrameRateDen,
	}
	r.frameRates = append(r.frameRates, frameRate)
	return frameRate, nil
}

func (r *mutationResolver) CreateResolution(ctx context.Context, input model.NewResolution) (*model.Resolution, error) {
	resolution := &model.Resolution{
		ID: fmt.Sprintf("T%d", rand.Int()),
		X:  input.X,
		Y:  input.Y,
	}
	r.resolutions = append(r.resolutions, resolution)
	return resolution, nil
}

func (r *mutationResolver) CreateProject(ctx context.Context, input model.NewProject) (*model.Project, error) {
	frameRate, err := r.frameRateById(ctx, input.FrameRateID)
	if err != nil {
		return nil, err
	}
	resolution, err := r.resolutionById(ctx, input.ResolutionID)
	if err != nil {
		return nil, err
	}
	project := &model.Project{
		ID:          fmt.Sprintf("T%d", rand.Int()),
		Name:        input.Name,
		Label:       input.Label,
		Description: input.Description,
		FrameRate:   frameRate,
		Resolution:  resolution,
	}
	r.projects = append(r.projects, project)
	return project, nil
}

func (r *queryResolver) FrameRates(ctx context.Context) ([]*model.FrameRate, error) {
	return r.frameRates, nil
}

func (r *queryResolver) Resolutions(ctx context.Context) ([]*model.Resolution, error) {
	return r.resolutions, nil
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	return r.projects, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func (r *Resolver) frameRateById(ctx context.Context, id string) (*model.FrameRate, error) {
	frameRates, err := r.Query().FrameRates(ctx)
	if err != nil {
		return nil, err
	}
	for _, frameRate := range frameRates {
		if frameRate.ID == id {
			return frameRate, nil
		}
	}
	return nil, fmt.Errorf("Could not find frame rate with ID '%s'", id)
}

func (r *Resolver) resolutionById(ctx context.Context, id string) (*model.Resolution, error) {
	resolutions, err := r.Query().Resolutions(ctx)
	if err != nil {
		return nil, err
	}
	for _, resolution := range resolutions {
		if resolution.ID == id {
			return resolution, nil
		}
	}
	return nil, fmt.Errorf("Could not find resolution with ID '%s'", id)
}
