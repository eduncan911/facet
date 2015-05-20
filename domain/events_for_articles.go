package domain

import (
	"github.com/eduncan911/es"
)

// ArticleCreated represents an article that has been posted
type ArticleCreated struct {
	EventID   EventID   `json:"eventId,omitempty"`
	ArticleID ArticleID `json:"articleId,omitempty"`
	Author    UserID    `json:"authorId,omitempty"`
	URL       string    `json:"urlKey"`
	Headline  string    `json:"headline,omitempty"`
	Body      string    `json:"body,omitempty"`
	CommitMsg string    `json:"commitMsg,omitempty"`
}

func NewArticleCreated(eventID EventID, articleID ArticleID, author UserID, url, headline, body, commitMsg string) *ArticleCreated {
	return &ArticleCreated{
		EventID:   eventID,
		ArticleID: articleID,
		Author:    author,
		URL:       url,
		Headline:  headline,
		Body:      body,
		CommitMsg: commitMsg,
	}
}

func (e *ArticleCreated) Meta() *es.Info { return i("ArticleCreated", e.EventID) }

// ArticleUpdated represents an article that has been updated with a PUT
type ArticleUpdated struct {
	EventID   EventID   `json:"eventId,omitempty"`
	ArticleID ArticleID `json:"articleId,omitempty"`
	Author    UserID    `json:"authorId,omitempty"`
	URL       string    `json:"urlKey"`
	Headline  string    `json:"headline,omitempty"`
	Body      string    `json:"body,omitempty"`
	CommitMsg string    `json:"commitMsg,omitempty"`
}

func NewArticleUpdated(eventID EventID, articleID ArticleID, author UserID, url, headline, body, commitMsg string) *ArticleUpdated {
	return &ArticleUpdated{
		EventID:   eventID,
		ArticleID: articleID,
		Author:    author,
		URL:       url,
		Headline:  headline,
		Body:      body,
		CommitMsg: commitMsg,
	}
}

func (e *ArticleUpdated) Meta() *es.Info { return i("ArticleUpdated", e.EventID) }

// ArticlePublished represents an article that has been published
type ArticlePublished struct {
	EventID     EventID   `json:"eventId,omitempty"`
	ArticleID   ArticleID `json:"articleId,omitempty"`
	PublishedBy UserID    `json:"publishedBy,omitempty"`
	CommitMsg   string    `json:"commitMsg,omitempty"`
}

func NewArticlePublished(eventID EventID, articleID ArticleID, publishedBy UserID, commitMsg string) *ArticlePublished {
	return &ArticlePublished{
		EventID:     eventID,
		ArticleID:   articleID,
		PublishedBy: publishedBy,
		CommitMsg:   commitMsg,
	}
}

func (e *ArticlePublished) Meta() *es.Info { return i("ArticlePublished", e.EventID) }

// ArticleRemoved represents an article that has been marked removed/deleted
type ArticleRemoved struct {
	EventID   EventID   `json:"eventId,omitempty"`
	ArticleID ArticleID `json:"articleId,omitempty"`
	DeletedBy UserID    `json:"deletedBy,omitempty"`
	CommitMsg string    `json:"commitMsg,omitempty"`
}

func NewArticleRemoved(eventID EventID, articleID ArticleID, deletedBy UserID, commitMsg string) *ArticleRemoved {
	return &ArticleRemoved{
		EventID:   eventID,
		ArticleID: articleID,
		DeletedBy: deletedBy,
		CommitMsg: commitMsg,
	}
}

func (e *ArticleRemoved) Meta() *es.Info { return i("ArticleRemoved", e.EventID) }

// ArticleURLChanged denotes that an `Article`'s URL has changed
type ArticleURLChanged struct {
	EventID   EventID   `json:"eventId,omitempty"`
	ArticleID ArticleID `json:"articleId,omitempty"`
	OldURL    string    `json:"oldUrl,omitempty"`
	NewURL    string    `json:"newUrl,omitempty"`
	CommitMsg string    `json:"commitMsg,omitempty"`
}

func NewArticleURLChanged(eventID EventID, articleID ArticleID, oldURL, newURL, commitMsg string) *ArticleURLChanged {
	return &ArticleURLChanged{
		EventID:   eventID,
		ArticleID: articleID,
		OldURL:    oldURL,
		NewURL:    newURL,
		CommitMsg: commitMsg,
	}
}

func (e *ArticleURLChanged) Meta() *es.Info { return i("ArticleURLChanged", e.EventID) }
