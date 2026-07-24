//go:build darwin && cgo

package keyring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVaultKey(t *testing.T) {
	require.Equal(t, "STRIPE_CLI_64656661756C742E746573745F6D6F64655F6170695F6B6579", vaultKey("default.test_mode_api_key"))
	require.NotEqual(t, vaultKey("a-b"), vaultKey("a b"))
}

func TestCredentialRequestDetail(t *testing.T) {
	require.Equal(t, "stripe needs credential account.acct_123.test_mode_api_key", credentialRequestDetail("account.acct_123.test_mode_api_key"))
}

func TestApprovalServiceSigningRequirementPinsTeamAndIdentifier(t *testing.T) {
	require.Contains(t, approvalServiceSigningRequirement, `certificate leaf[subject.OU] = ZU76A67LGU`)
	require.Contains(t, approvalServiceSigningRequirement, `identifier "com.automicvault"`)
}

func TestApprovalEventNotice(t *testing.T) {
	require.Equal(t, humanApprovalRequiredNotice, approvalEventNotice(humanApprovalRequiredEvent))
	require.Empty(t, approvalEventNotice("other-event"))
}
