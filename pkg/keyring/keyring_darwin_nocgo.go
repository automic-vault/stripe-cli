//go:build darwin && !cgo

package keyring

import "errors"

type unavailableStore struct{}

func newSecureStore(_, _ string) SecureStore {
	return &unavailableStore{}
}

// ProtectsAllAPIKeys reports whether API keys must be stored in SecureStore.
func ProtectsAllAPIKeys() bool {
	return true
}

func (s *unavailableStore) Get(string) ([]byte, error) {
	return nil, errors.New("Automic Vault credential storage requires a CGO-enabled Stripe CLI build")
}

func (s *unavailableStore) Set(string, []byte, string) error {
	return errors.New("Automic Vault credential storage requires a CGO-enabled Stripe CLI build")
}

func (s *unavailableStore) Remove(string) error {
	return errors.New("Automic Vault credential storage requires a CGO-enabled Stripe CLI build")
}
