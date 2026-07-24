//go:build !darwin

package keyring

func newSecureStore(service, credentialsFilePath string) SecureStore {
	return &fallbackStore{
		primary:  newZalandoStore(service),
		fallback: &fileStore{path: credentialsFilePath},
	}
}

// ProtectsAllAPIKeys reports whether API keys must be stored in SecureStore.
func ProtectsAllAPIKeys() bool {
	return false
}
