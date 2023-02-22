package graph

import (
	"github.com/Luuwa/deprem-yardim-graphql/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
func New() Config {
	c := Config{
		Resolvers: &Resolver{
			latlngs: []*model.LatLng{
				{Lat: 1.3, Lng: 2.3}, {Lat: 2.3, Lng: 2.4},
			},
		},
	}

	return c
}

type Resolver struct {
	latlngs []*model.LatLng
}
