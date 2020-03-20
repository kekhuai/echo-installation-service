package model

import "github.com/Kamva/mgm/v2"

// Partner model
type Partner struct {
	mgm.DefaultModel `bson:"inline"`
	Name             string `json:"name" bson:"name"`
}
