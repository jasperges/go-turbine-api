package graph

import "github.com/jasperges/go-turbine-api/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	frameRates  []*model.FrameRate
	resolutions []*model.Resolution
	projects    []*model.Project
}
