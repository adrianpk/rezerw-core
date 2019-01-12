package core

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Account - Account
type Account struct {
	ID               uuid.UUID `json:"id" schema:"id"`
	Name             string    `json:"name" schema:"name"`
	Description      string    `json:"description" schema:"description"`
	AccountType      string    `json:"accountType" schema:"account-type"`
	OwnerID          uuid.UUID `json:"ownerID" schema:"owner-id"`
	ParentID         uuid.UUID `json:"parentID" schema:"parent-id"`
	Email            string    `json:"email" schema:"email"`
	GeoData          GeoJSON   `json:"geoDate" schema:"geo-data"`
	GeoHash          string    `json:"geoHash" schema:"geo-hash"`
	StartsAt         time.Time `json:"startsAt,omitempty" schema:"-"`
	EndsAt           time.Time `json:"endsAt,omitempty" schema:"-"`
	IsActive         bool      `json:"isActive,omitempty" schema:"is-active"`
	IsLogicalDeleted bool      `json:"isLogicalDeleted,omitempty" schema:"is-logical-deleted"`
	CreatedBy        uuid.UUID `json:"CreatedBy,omitempty" schema:"-"`
	UpdatedBy        uuid.UUID `json:"UpdatedBy,omitempty" schema:"-"`
	CreatedAt        time.Time `json:"createdAt,omitempty" schema:"-"`
	UpdatedAt        time.Time `json:"updatedAt,omitempty" schema:"-"`
}

// GeoJSON - Geo JSON.
type GeoJSON struct {
	ID         uuid.UUID `json:"id" schema:"id"`
	GeoType    string    `json:"geoType" schema:"id"`
	Properties map[string]interface{}
}

// GeoGeometry - Geo geometry.
type GeoGeometry struct {
	GeometryType string `json:"geometryType" schema:"geometry-type"`
	Coordinates  []GeoCoordinate
}

// GeoCoordinate - Geo coordinate.
type GeoCoordinate struct {
	LatLong []float64
}
