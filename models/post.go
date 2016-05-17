package models

import (
	"github.com/sunbirdframework/sunbird"
)

// Post ...
type Post struct {
	sunbird.Model
}

// TableName ...
func (p Post) TableName() string {
	return "posts"
}
