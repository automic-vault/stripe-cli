//go:build darwin

package keyring

// IsUsingInsecureStorage is always false because the macOS store fails closed.
func IsUsingInsecureStorage(SecureStore) bool {
	return false
}

// FallbackStoragePath is empty because the macOS store has no file fallback.
func FallbackStoragePath(SecureStore) string {
	return ""
}
