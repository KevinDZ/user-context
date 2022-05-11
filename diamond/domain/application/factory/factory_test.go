package factory

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInstanceApplicationAggregate(t *testing.T) {
	rootID := ""
	instance := InstanceApplicationAggregate(rootID)
	require.Equal(t, instance.Aggregate.RootID, rootID, "no equal")
}

func TestFactory_InstanceOf(t *testing.T) {
	rootID := ""
	require.True(t, InstanceApplicationAggregate(rootID).InstanceOf(), "application instance failed")
}
