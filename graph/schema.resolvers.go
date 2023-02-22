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
	LatLng := &model.LatLng{
		Lat: input.Latitude,
		Lng: input.Longitude,
	}

	r.latlngs = append(r.latlngs, LatLng)
	return LatLng, nil
}

// CreateFeed is the resolver for the createFeed field.
func (r *mutationResolver) CreateFeed(ctx context.Context, input model.NewFeed) (*model.Feed, error) {
	feed := &model.Feed{
		FullText:         input.FullText,
		IsResolved:       input.IsResolved,
		Channel:          input.Channel,
		Timestamp:        input.Timestamp,
		Epoch:            input.Epoch,
		ExtraParameters:  input.ExtraParameters,
		FormattedAddress: input.FormattedAddress,
		Reason:           input.Reason,
	}

	r.feeds = append(r.feeds, feed)
	return feed, nil
}

// LatLngs is the resolver for the latLngs field.
func (r *queryResolver) LatLngs(ctx context.Context) ([]*model.LatLng, error) {
	var batch []*model.LatLng
	for _, ll := range r.latlngs {
		latlng := model.LatLng{Lat: ll.Lat, Lng: ll.Lng}
		batch = append(batch, &latlng)
	}
	return batch, nil
}

// Feeds is the resolver for the feeds field.
func (r *queryResolver) Feeds(ctx context.Context) ([]*model.Feed, error) {
	var batch []*model.Feed
	for _, f := range r.feeds {
		feed := model.Feed{ID: f.ID,
			FullText:         f.FullText,
			IsResolved:       f.IsResolved,
			Channel:          f.Channel,
			Timestamp:        f.Timestamp,
			Epoch:            f.Epoch,
			ExtraParameters:  f.ExtraParameters,
			FormattedAddress: f.FormattedAddress,
			Reason:           f.Reason}
		batch = append(batch, &feed)
	}
	return batch, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
