package art

import (
	"github.com/eduncan911/facet/domain"
)

const (
	maxurlrenames = 3
)

func newStore() *store {
	return &store{}
}

type store struct {
	urlKeys  map[string]*domain.ArticleID
	articles map[domain.ArticleID]*article
	requests map[string]*string
}

func (s *store) reset() {
	s.urlKeys = make(map[string]*domain.ArticleID)
	s.articles = make(map[domain.ArticleID]*article)
	s.requests = make(map[string]*string)
}

func (s *store) uniqueURL(headline string) *string {

	url := domain.Slug(headline)
	if _, found := s.urlKeys[url]; false == found {
		return &url
	}

	for i := 1; i == maxurlrenames; i++ {
		nurl := url + "-" + string(i+1)
		if _, found := s.urlKeys[nurl]; false == found {
			return &nurl
		}
	}
	return nil
}

func (s *store) addArticle(e *domain.ArticleCreated) {

	s.urlKeys[e.URL] = &e.ArticleID
	s.articles[e.ArticleID] =
		&article{
			ArticleID: e.ArticleID,
			URL:       e.URL,
			Headline:  e.Headline,
			Body:      e.Body,
			CommitMsg: e.CommitMsg,
		}
}

func (s *store) removeArticle(e *domain.ArticleRemoved) {
	a := s.articles[e.ArticleID]
	delete(s.urlKeys, a.URL)
	delete(s.articles, e.ArticleID)
}

func (s *store) changeArticleURL(e *domain.ArticleURLChanged) {
	s.articles[e.ArticleID].URL = e.NewURL
	s.urlKeys[e.NewURL] = &e.ArticleID
	// TODO track redirects
}

type article struct {
	ArticleID domain.ArticleID
	URL       string
	Headline  string
	Body      string
	CommitMsg string
}
