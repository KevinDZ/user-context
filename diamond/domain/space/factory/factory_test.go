package factory

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInstanceSpaceAggregate(t *testing.T) {
	rootID := ""
	require.Equal(t, InstanceSpaceAggregate(rootID).Aggregate.RootID, rootID, "no equal")
}

func TestFactory_InstanceOf(t *testing.T) {
	rootID := ""
	space := InstanceSpaceAggregate(rootID)
	require.True(t, space.InstanceOf(), "space instance failed")
}
