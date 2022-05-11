package factory

import (
	"github.com/stretchr/testify/require"
	"testing"
	"user-context/diamond/acl/adapters/clients"
	"user-context/diamond/acl/adapters/repositories"
)

func TestFactory_InstanceOf(t *testing.T) {
	rootID := ""
	instance := InstanceAccountAggregate(rootID)
	require.Error(t, instance.InstanceOf(), "")
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
