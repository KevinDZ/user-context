package factory

import (
	"github.com/stretchr/testify/require"
	"testing"
	"user-context/rhombic/acl/adapters/clients"
	"user-context/rhombic/acl/adapters/repositories"
)

func TestFactory_InstanceOf(t *testing.T) {
	rootID := ""
	require.Error(t, InstanceAccountAggregate(rootID).InstanceOf(), "")
	return
}

func TestFactory_Registered(t *testing.T) {
	rootID := ""
	instance := InstanceAccountAggregate(rootID)
	instance.Service.Client = clients.NewUUIDAdapter()
	instance.Service.Repository = repositories.NewAccountAdapter()
	require.Error(t, instance.Registered(), "")
	return
}
