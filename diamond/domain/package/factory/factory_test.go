package factory

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInstancePackageAggregate(t *testing.T) {
	rootID := ""
	require.Equal(t, InstancePackageAggregate(rootID).Root.RootID, rootID, "no equal")
}

func TestFactory_InstanceOf(t *testing.T) {
	rootID := ""
	packages := InstancePackageAggregate(rootID)
	require.True(t, packages.InstanceOf(), "package instance failed")
}
