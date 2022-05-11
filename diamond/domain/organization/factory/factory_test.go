package factory

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInstanceOrganizationAggregate(t *testing.T) {
	rootID := ""
	require.Equal(t, InstanceOrganizationAggregate(rootID).Aggregate.RootID, rootID, "no equal")
}

func TestFactory_InstanceOf(t *testing.T) {
	rootID := ""
	organization := InstanceOrganizationAggregate(rootID)
	require.True(t, organization.InstanceOf(), "organization instance failed")
}
