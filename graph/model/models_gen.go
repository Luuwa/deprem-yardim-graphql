// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Feed struct {
	ID               int       `json:"ID"`
	FullText         string    `json:"FullText"`
	IsResolved       bool      `json:"IsResolved"`
	Channel          string    `json:"Channel"`
	Timestamp        time.Time `json:"Timestamp"`
	Epoch            int       `json:"Epoch"`
	ExtraParameters  *string   `json:"ExtraParameters"`
	FormattedAddress string    `json:"FormattedAddress"`
	Reason           *string   `json:"Reason"`
}

type FeedLocation struct {
	EntryID   int     `json:"EntryID"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	Address   string  `json:"Address"`
}

type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type LiteResult struct {
	ID  int       `json:"ID"`
	Loc []float64 `json:"Loc"`
}

type Location struct {
	ID               int       `json:"ID"`
	FormattedAddress string    `json:"FormattedAddress"`
	Latitude         float64   `json:"Latitude"`
	Longitude        float64   `json:"Longitude"`
	NortheastLat     float64   `json:"NortheastLat"`
	NortheastLng     float64   `json:"NortheastLng"`
	SouthwestLat     float64   `json:"SouthwestLat"`
	SouthwestLng     float64   `json:"SouthwestLng"`
	EntryID          int       `json:"EntryID"`
	Timestamp        time.Time `json:"Timestamp"`
	Epoch            int       `json:"Epoch"`
	Reason           string    `json:"Reason"`
	Channel          string    `json:"Channel"`
}

type NewFeed struct {
	FullText         string    `json:"FullText"`
	IsResolved       bool      `json:"IsResolved"`
	Channel          string    `json:"Channel"`
	Timestamp        time.Time `json:"Timestamp"`
	Epoch            int       `json:"Epoch"`
	ExtraParameters  *string   `json:"ExtraParameters"`
	FormattedAddress string    `json:"FormattedAddress"`
	Reason           *string   `json:"Reason"`
}

type NewLatLng struct {
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
}

type Response struct {
	Count   int       `json:"Count"`
	Results []*Result `json:"Results"`
}

type Result struct {
	ID                 int       `json:"ID"`
	Loc                []float64 `json:"Loc"`
	EntryID            int       `json:"Entry_ID"`
	Timestamp          *string   `json:"Timestamp"`
	Epoch              int       `json:"Epoch"`
	Reason             *string   `json:"Reason"`
	Channel            *string   `json:"Channel"`
	ExtraParameters    *string   `json:"ExtraParameters"`
	IsLocationVerified bool      `json:"IsLocationVerified"`
	IsNeedVerified     bool      `json:"IsNeedVerified"`
}

type UpdateFeedLocationsRequest struct {
	FeedLocations []*FeedLocation `json:"FeedLocations"`
}
