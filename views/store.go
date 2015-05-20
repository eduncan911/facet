package views

import (
	"github.com/eduncan911/facet/domain"
	"time"
)

func newStore() *store {
	return &store{}
}

type store struct {
	all map[domain.ArticleID]*articleItem
}

func (s *store) reset() {
	s.all = make(map[domain.ArticleID]*articleItem)
}

func (s *store) addArticle(id domain.ArticleID, headline, body string) error {

	// TODO protect against updating existing?
	//

	// create the article
	s.all[id] = &articleItem{
		ArticleID: id,
		Headline:  headline,
		Body:      body,
		BodyHTML:  "", // TODO store Markdown->HTML conversions here?
		Published: false,
		Created:   time.Now(),
		Updated:   time.Now(),
	}

	// TODO: add author's commit
	//
	// another module handling history/diff views w/author's info?

	return nil
}

func (s *store) publishArticle(id domain.ArticleID) {
	s.all[id].Published = true
}

func (s *store) removeArticle(id domain.ArticleID) {
	delete(s.all, id)
}

type articleItem struct {
	ArticleID domain.ArticleID `json:"articleId,omitempty"`
	Headline  string           `json:"headline,omitempty"`
	Body      string           `json:"body,omitempty"`
	BodyHTML  string           `json:"bodyHtml,omitempty"`
	Published bool             `json:"published,omitempty"`
	Created   time.Time        `json:"created,omitempty"` // TODO save state here?
	Updated   time.Time        `json:"updated,omitempty"` // TODO save state here?
}
