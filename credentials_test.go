package epitome_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/arturovm/epitome"
)

func TestGenerateCredentials(t *testing.T) {
	password := "testpassword"
	creds, err := epitome.GenerateCredentials(password)
	require.NoError(t, err)
	require.NotNil(t, creds)
	require.NotEqual(t, password, string(creds.Password))
}

func TestMatchPassword(t *testing.T) {
	password := "this is a test password"
	creds, _ := epitome.GenerateCredentials(password)

	match := creds.MatchPassword(password)
	require.True(t, match)
}
