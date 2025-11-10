package xmemory_test

import (
	"context"
	"testing"

	"github.com/agussyahrilmubarok/gox/xdiscovery"
	"github.com/agussyahrilmubarok/gox/xdiscovery/xmemory"
	"github.com/stretchr/testify/assert"
)

func TestMemoryRegistry(t *testing.T) {
	ctx := context.Background()
	reg := xmemory.NewRegistry()
	serviceName := "test-service"
	instanceID := "instance-1"
	hostPort := "127.0.0.1:8080"

	err := reg.Register(ctx, instanceID, serviceName, hostPort)
	assert.NoError(t, err)

	addrs, err := reg.ServiceAddresses(ctx, serviceName)
	assert.NoError(t, err)
	assert.Len(t, addrs, 1)
	assert.Equal(t, hostPort, addrs[0])

	err = reg.ReportHealthyState(instanceID, serviceName)
	assert.NoError(t, err)

	err = reg.Deregister(ctx, instanceID, serviceName)
	assert.NoError(t, err)

	addrs, err = reg.ServiceAddresses(ctx, serviceName)
	assert.ErrorIs(t, err, xdiscovery.ErrNotFound)
	assert.Len(t, addrs, 0)

	err = reg.ReportHealthyState(instanceID, serviceName)
	assert.Error(t, err)
}
