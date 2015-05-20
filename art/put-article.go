package art

import (
	"github.com/eduncan911/es/api"
	"github.com/eduncan911/facet/ctx"
	"github.com/eduncan911/facet/domain"
)

type putArticleRequest struct {
	URL       string `json:"urlKey,omitempty"`
	Headline  string `json:"headline,omitempty"`
	Body      string `json:"body,omitempty"`
	CommitMsg string `json:"commitmsg,omitempty"`
}

type putArticleResponse struct {
	URL string `json:"urlKey,omitempty"`
}

func (m *Module) putArticle(req *api.Request) api.Response {

	var r putArticleRequest
	if err := req.ParseBody(&r); err != nil {
		return api.BadRequestResponse(err.Error())
	}
	if r.Headline == "" || r.Body == "" || r.CommitMsg == "" {
		return api.BadRequestResponse("Expected fields: headline, body and commitmsg")
	}

	c := ctx.Parse(req)
	articleID := c.URLArticleID

	a, ok := m.s.articles[*articleID]
	if false == ok {
		return api.BadRequestResponse("Article with that ID does not exist: " + articleID.String())
	}

	// TODO future logic:
	// - watchdog
	// - anti-spam
	// - suggest better headlines
	// - spell check headlines
	//

	url := &r.URL
	if *url != "" && *url != a.URL {
		if url = m.s.uniqueURL(*url); url == nil {
			return api.BadRequestResponse("That new URL has too many collisions")
		}

		m.pub.MustPublish(
			domain.NewArticleURLChanged(
				domain.NewEventID(),
				*articleID,
				a.URL,
				*url,
				r.CommitMsg,
			),
		)
	}

	authorID := c.User.UserID
	m.pub.MustPublish(
		domain.NewArticleUpdated(
			domain.NewEventID(),
			*articleID,
			*authorID,
			*url,
			r.Headline,
			r.Body,
			r.CommitMsg,
		),
	)

	return api.NewJSON(&putArticleResponse{
		URL: *url,
	})
}
