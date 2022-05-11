package factory

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInstancePermissionAggregate(t *testing.T) {
	rootID := ""
	require.Equal(t, InstancePermissionAggregate(rootID).Root.RootID, rootID, "no equal")
}

func TestFactory_InstanceOf(t *testing.T) {
	rootID := ""
	space := InstancePermissionAggregate(rootID)
	require.True(t, space.InstanceOf(), "permission instance failed")
}
