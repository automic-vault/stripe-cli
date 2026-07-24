package keys

import (
	"os"
	"testing"

	zkr "github.com/zalando/go-keyring"

	"github.com/stripe/stripe-cli/pkg/config"
	"github.com/stripe/stripe-cli/pkg/keyring"
)

func TestMain(m *testing.M) {
	zkr.MockInit()
	config.KeyRing = keyring.NewMemoryStore(nil)
	os.Exit(m.Run())
}
