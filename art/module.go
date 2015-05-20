package art

import (
	"github.com/eduncan911/es/env"
)

// Module represents the code contract for this package
type Module struct {
	pub env.Publisher
	d   *denormalizer
	s   *store
}

// NewModule instantiates a new Module with the specified `env.Publisher`
func NewModule(pub env.Publisher) *Module {
	store := newStore()
	denormalizer := newDenormalizer(store)
	return &Module{
		pub,
		denormalizer,
		store,
	}
}

// Register takes an `env.Registrar` to register API endpoints and the module's event handler
func (m *Module) Register(r env.Registrar) {
	r.ResetData("store", m.s.reset)
	r.HandleEvents("art-denormalizer", m.d)

	r.HandleHTTP("POST", "/article", m.postArticle)
	r.HandleHTTP("PUT", "/article/{id:[0-9a-f]{32}+}", m.putArticle)
	//r.HandleHTTP("DELETE", "/article/{id:[0-9a-f]{32}+}", m.removeArticle)
	//r.HandleHTTP("PUT", "/article/{id:[0-9a-f]{32}+}/publish", m.putPublishArticle)

}
