package factory

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInstanceApplicationAggregate(t *testing.T) {
	rootID := ""
	require.Equal(t, InstanceApplicationAggregate(rootID).Root.RootID, rootID, "no equal")
}

func TestFactory_InstanceOf(t *testing.T) {
	rootID := ""
	application := InstanceApplicationAggregate(rootID)
	require.True(t, application.InstanceOf(), "application instance failed")
}