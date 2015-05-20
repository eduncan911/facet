package art

import (
	"crypto/sha1"
	"fmt"

	"github.com/eduncan911/es/api"
	"github.com/eduncan911/es/log"
	"github.com/eduncan911/facet/ctx"
	"github.com/eduncan911/facet/domain"
	"github.com/op/go-logging"
)

type postArticleRequest struct {
	Headline  string `json:"headline,omitempty"`
	Body      string `json:"body,omitempty"`
	CommitMsg string `json:"commitmsg,omitempty"`
}

type postArticleResponse struct {
	ArticleID domain.ArticleID `json:"articleId,omitempty"`
	URL       string           `json:"urlKey,omitempty"`
}

func (m *Module) postArticle(req *api.Request) api.Response {

	log.Init("art")
	logging.

	var r postArticleRequest
	if err := req.ParseBody(&r); err != nil {
		return api.BadRequestResponse(err.Error())
	}
	if r.Headline == "" || r.Body == "" || r.CommitMsg == "" {
		return api.BadRequestResponse("Expected fields: headline, body and commitmsg")
	}

	c := ctx.Parse(req)

	// TODO move this to context
	d := fmt.Sprintf("POST /article %v %s", r, c.Fingerprint.String())
	sig := fmt.Sprintf("%x", sha1.Sum([]byte(d)))
	if url, ok := m.s.requests[sig]; ok {
		return api.NewJSON(&postArticleResponse{
			URL: *url,
		})
	}

	var url *string
	if url = m.s.uniqueURL(r.Headline); url == nil {
		return api.BadRequestResponse("That Headline has too many URL collisions")
	}

	// TODO future logic:
	// - watchdog
	// - anti-spam
	// - suggest better headlines
	// - spell check headlines
	//

	authorID := c.User.UserID
	articleID := domain.NewArticleID()
	m.pub.MustPublish(
		domain.NewArticleCreated(
			domain.NewEventID(),
			articleID,
			*authorID,
			*url,
			r.Headline,
			r.Body,
			r.CommitMsg,
		),
	)

	m.s.requests[sig] = url
	return api.NewJSON(&postArticleResponse{
		ArticleID: articleID,
		URL:       *url,
	})
}
