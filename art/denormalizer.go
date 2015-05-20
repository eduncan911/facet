package art

import (
	"github.com/eduncan911/es"
	"github.com/eduncan911/facet/domain"
)

type denormalizer struct {
	s *store
}

func newDenormalizer(s *store) *denormalizer {
	return &denormalizer{
		s,
	}
}

func (d *denormalizer) HandleEvent(e es.Event) error {

	// route the event type
	switch t := e.(type) {
	case *domain.ArticleCreated:
		d.s.addArticle(t)
	//case *domain.ArticleUpdated:
	//	d.s.updateArticle(t)
	case *domain.ArticleRemoved:
		d.s.removeArticle(t)
	case *domain.ArticleURLChanged:
		d.s.changeArticleURL(t)
	}

	return nil
}

func (d *denormalizer) Contracts() []string {
	return []string{
		"ArticleCreated",
		"ArticleUpdated",
		"ArticleRemoved",
		"ArticleURLChanged",
		"UserNameChanged",
	}
}
