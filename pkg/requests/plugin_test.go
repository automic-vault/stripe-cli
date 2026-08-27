package requests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stripe/stripe-cli/pkg/config"
	"github.com/stripe/stripe-cli/pkg/keyring"
)

type recordingKeyring struct {
	gets int
}

func (r *recordingKeyring) Get(string) ([]byte, error) {
	r.gets++
	return nil, keyring.ErrKeyNotFound
}

func (*recordingKeyring) Set(string, []byte, string) error { return nil }
func (*recordingKeyring) Remove(string) error              { return nil }

func TestAnonymousPluginMetadataDoesNotReadCredentials(t *testing.T) {
	ring := &recordingKeyring{}
	originalKeyRing := config.KeyRing
	config.KeyRing = ring
	t.Cleanup(func() { config.KeyRing = originalKeyRing })

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/ajax/stripecli/plugins_metadata", r.URL.Path)
		require.Empty(t, r.Header.Get("Authorization"))
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	_, err := GetPluginMetadata(context.Background(), "https://api.example.test", server.URL, "2026-08-27", "", &config.Profile{}, "docs", "", "darwin", "arm64", "machine")
	require.NoError(t, err)
	require.Zero(t, ring.gets)
}
