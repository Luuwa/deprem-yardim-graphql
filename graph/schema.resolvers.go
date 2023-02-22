package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"

	"github.com/Luuwa/deprem-yardim-graphql/graph/model"
)

// CreateLatLng is the resolver for the createLatLng field.
func (r *mutationResolver) CreateLatLng(ctx context.Context, input model.NewLatLng) (*model.LatLng, error) {
	Cor := &model.LatLng{
		Lat: input.Latitude,
		Lng: input.Longitude,
	}

	r.latlngs = append(r.latlngs, Cor)
	return Cor, nil
}

// LatLngs is the resolver for the latLngs field.
func (r *queryResolver) LatLngs(ctx context.Context) ([]*model.LatLng, error) {
	var latlngArr []*model.LatLng
	for _, latlng := range r.latlngs {
		var latlngx = model.LatLng{Lat: latlng.Lat, Lng: latlng.Lng}
		latlngArr = append(latlngArr, &latlngx)
	}
	return latlngArr, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
