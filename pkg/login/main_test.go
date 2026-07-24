package login

import (
	"os"
	"testing"

	"github.com/stripe/stripe-cli/pkg/config"
	"github.com/stripe/stripe-cli/pkg/keyring"
)

func TestMain(m *testing.M) {
	config.KeyRing = keyring.NewMemoryStore(nil)
	os.Exit(m.Run())
}
