// Package keyring provides credential storage.
package keyring

import "errors"

// ErrKeyNotFound is returned when a key is not present in the secure store.
var ErrKeyNotFound = errors.New("secure store: key not found")

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
