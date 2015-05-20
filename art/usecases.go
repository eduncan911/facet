package art

import (
	"github.com/abdullin/seq"
	"github.com/eduncan911/es/env"
	"github.com/eduncan911/es/spec"
	"github.com/eduncan911/facet/domain"
)

var useCases = []env.UseCaseFactory{

	when_POST_Article_then_ArticleCreated_published,
}

func newEventID() domain.EventID {
	return domain.NewEventID()
}

func newArticleID() domain.ArticleID {
	return domain.NewArticleID()
}

var IgnoreEventID domain.EventID

func when_POST_Article_then_ArticleCreated_published() *env.UseCase {

	articleID := domain.NewArticleID()

	return &env.UseCase{
		Name: "when POST /article, then 1 ArticleCreated event is published",
		When: spec.PostJSON("/article", seq.Map{
			"headline":  "The Headline",
			"body":      "The Body of Text",
			"commitmsg": "Creating initial article",
		}),
		ThenResponse: spec.ReturnJSON(seq.Map{
			"urlKey":    "the-headline",
			"articleId": articleID,
		}),
		ThenEvents: spec.Events(
			domain.NewArticleCreated(
				IgnoreEventID,
				articleID,
				domain.NewUserID(),
				"the-headline",
				"The Headline",
				"The Body of Text",
				"Creating initial article",
			),
		),
		Where: spec.Where{
			articleID:     "sameArticleID",
			IgnoreEventID: "ignore",
		},
	}

}
