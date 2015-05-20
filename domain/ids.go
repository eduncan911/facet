package domain

import (
	"github.com/eduncan911/go/pkg/uuid"
)

// NoEventID represents an empty EventID
var NoEventID = EventID{}

// NoArticleID represents an empty ArticleID
var NoArticleID = ArticleID{}

// NoUserID represents an empty UserID
var NoUserID = UserID{}

// EventID is the uuid for the Event
type EventID struct{ uuid.ID }

// ArticleID is the uuid for Articles
type ArticleID struct{ uuid.ID }

// CatID is the uuid for Categories
type CatID struct{ uuid.ID }

// UserID is the uuid for Users
type UserID struct{ uuid.ID }

// NewEventID instantiates a new EventID
func NewEventID() EventID {
	return EventID{uuid.NewID()}
}

// NewArticleID instantiates a new ArticleID
func NewArticleID() ArticleID {
	return ArticleID{uuid.NewID()}
}

// NewCatID instantiates a new CatID
func NewCatID() CatID {
	return CatID{uuid.NewID()}
}

// NewUserID instantiates a new UserID
func NewUserID() UserID {
	return UserID{uuid.NewID()}
}
