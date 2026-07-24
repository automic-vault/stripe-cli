// Package keyring provides credential storage.
package keyring

import "github.com/stripe/stripe-cli/pkg/errorcategory"

// ErrKeyNotFound is returned when a key is not present in the secure store.
var ErrKeyNotFound = errorcategory.New(errorcategory.Filesystem, "secure store: key not found")

// SecureStore provides access to credential storage.
type SecureStore interface {
	Get(key string) ([]byte, error)
	Set(key string, data []byte, description string) error
	Remove(key string) error
}

// NewSecureStore returns the platform credential store.
func NewSecureStore(service, credentialsFilePath string) SecureStore {
	return newSecureStore(service, credentialsFilePath)
}
