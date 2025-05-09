package main

import (
	"context"

	"github.com/oxidecomputer/oxide.go/oxide"
)

type oxideClient interface {
	CurrentUserSshKeyCreate(ctx context.Context, params oxide.CurrentUserSshKeyCreateParams) (*oxide.SshKey, error)
	CurrentUserSshKeyDelete(ctx context.Context, params oxide.CurrentUserSshKeyDeleteParams) error
	DiskDelete(ctx context.Context, params oxide.DiskDeleteParams) error
	InstanceCreate(ctx context.Context, params oxide.InstanceCreateParams) (*oxide.Instance, error)
	InstanceDelete(ctx context.Context, params oxide.InstanceDeleteParams) error
	InstanceDiskListAllPages(ctx context.Context, params oxide.InstanceDiskListParams) ([]oxide.Disk, error)
	InstanceNetworkInterfaceListAllPages(ctx context.Context, params oxide.InstanceNetworkInterfaceListParams) ([]oxide.InstanceNetworkInterface, error)
	InstanceReboot(ctx context.Context, params oxide.InstanceRebootParams) (*oxide.Instance, error)
	InstanceStart(ctx context.Context, params oxide.InstanceStartParams) (*oxide.Instance, error)
	InstanceStop(ctx context.Context, params oxide.InstanceStopParams) (*oxide.Instance, error)
	InstanceView(ctx context.Context, params oxide.InstanceViewParams) (*oxide.Instance, error)
}

func (d *Driver) assertClient() {
	if d.oxideClient == nil {
		client, err := d.createOxideClient()
		if err != nil {
			panic(err)
		}
		d.oxideClient = client
	}
}

// createOxideClient creates an Oxide oxideClient from the machine driver
// configuration.
func (d *Driver) createOxideClient() (oxideClient, error) {
	return oxide.NewClient(&oxide.Config{
		Host:      d.Host,
		Token:     d.Token,
		UserAgent: "Oxide Rancher Machine Driver/0.0.1 (Go; Linux) [Environment: Development]",
	})
}
