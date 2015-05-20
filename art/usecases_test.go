package art

import (
	"testing"

	"github.com/eduncan911/es/spec"
)

func TestUseCases(t *testing.T) {
	ctx := spec.NewContext(Spec)
	mod := NewModule(ctx.Pub())
	ctx.Verify(mod).ToTesting(t)
}
