// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Copyright 2024 Oxide Computer Company
package main

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/oxidecomputer/oxide.go/oxide"
	"github.com/oxidecomputer/rancher-machine-driver-oxide/mock"
	"github.com/rancher/machine/commands/commandstest"
	"github.com/rancher/machine/libmachine/state"
	"go.uber.org/mock/gomock"
)

//go:generate mockgen -destination mock/oxide_client.go -package mock . oxideClient
//go:generate mockgen -destination mock/driver_options.go -package mock github.com/rancher/machine/libmachine/drivers DriverOptions

var _ = Describe("Driver", func() {
	var SUT *Driver
	var ctrl *gomock.Controller
	var opts *mock.MockDriverOptions

	BeforeEach(func() {
		SUT = newDriver("bob", "path")
		ctrl = gomock.NewController(GinkgoT())
		opts = mock.NewMockDriverOptions(ctrl)
	})

	Describe("SetConfigFromFlags", func() {
		It("should succeed when all required fields are given", func() {
			defaultDriverOptions(opts)
			Expect(SUT.SetConfigFromFlags(opts)).To(Succeed())
		})

		Describe("errors", func() {
			DescribeTable("should fail when a required string field is missing",
				func(fields []string) {
					for _, field := range fields {
						opts.EXPECT().String(gomock.Eq(field)).Return("").AnyTimes()
					}
					defaultDriverOptions(opts)
					err := SUT.SetConfigFromFlags(opts)
					Expect(err).To(HaveOccurred())
				},
				Entry("host", []string{flagHost}),
				Entry("token", []string{flagToken}),
				Entry("project", []string{flagProject}),
				Entry("diskImageId", []string{flagBootDiskImageID}),
			)

			It("should fail when nothing is given", func() {
				err := SUT.SetConfigFromFlags(&commandstest.FakeFlagger{
					Data: map[string]any{},
				})
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("required option \"oxide-host\" not set"))
				Expect(err.Error()).To(ContainSubstring("required option \"oxide-token\" not set"))
				Expect(err.Error()).To(ContainSubstring("required option \"oxide-project\" not set"))
				Expect(err.Error()).To(ContainSubstring("required option \"oxide-boot-disk-image-id\" not set"))
			})
		})
	})

	DescribeTable("RancherMachineState mapping is correct",
		func(instanceState oxide.InstanceState, expectedState state.State) {
			Expect(toRancherMachineState(instanceState)).To(Equal(expectedState))
		},
		Entry("creating", oxide.InstanceStateCreating, state.Starting),
		Entry("starting", oxide.InstanceStateStarting, state.Starting),
		Entry("running", oxide.InstanceStateRunning, state.Running),
		Entry("stopping", oxide.InstanceStateStopping, state.Stopping),
		Entry("stopped", oxide.InstanceStateStopped, state.Stopped),
		Entry("repairing", oxide.InstanceStateRepairing, state.Starting),
		Entry("rebooting", oxide.InstanceStateRebooting, state.Starting),
		Entry("migrating", oxide.InstanceStateMigrating, state.Running),
		Entry("failed", oxide.InstanceStateFailed, state.Error),
		Entry("destroyed", oxide.InstanceStateDestroyed, state.NotFound),
		Entry("unknown", oxide.InstanceState("unknown"), state.None),
	)

	Describe("ParseAdditionalDisk", func() {
		DescribeTable("Success",
			func(s string, expected AdditionalDisk) {
				Expect(ParseAdditionalDisk(s)).To(Equal(expected))
			},
			Entry("parses integer without suffix", "21474836480", AdditionalDisk{Size: 21474836480, Label: "additional"}),
			Entry("parses integer with suffix", "10GiB", AdditionalDisk{Size: 10737418240, Label: "additional"}),
			Entry("parses integer with space suffix", "10 GiB", AdditionalDisk{Size: 10737418240, Label: "additional"}),
			Entry("parses integer without suffix and label", "21474836480,data", AdditionalDisk{Size: 21474836480, Label: "data"}),
			Entry("parses integer with suffix and label", "10GiB,data", AdditionalDisk{Size: 10737418240, Label: "data"}),
			Entry("parses integer with space suffix and label", "10 GiB,data", AdditionalDisk{Size: 10737418240, Label: "data"}),
			Entry("parses integer without suffix trailing comma", "21474836480,", AdditionalDisk{Size: 21474836480, Label: "additional"}),
			Entry("parses integer with suffix trailing comma", "10GiB,", AdditionalDisk{Size: 10737418240, Label: "additional"}),
		)

		DescribeTable("Error",
			func(s string) {
				_, err := ParseAdditionalDisk(s)
				Expect(err).To(HaveOccurred())
			},
			Entry("errors with empty string", ""),
			Entry("errors with empty invalid format", ","),
			Entry("errors with no size", ",foo"),
			Entry("errors with invalid size unit suffix", "20 ABC,"),
		)
	})

	Describe("Workflow", func() {
		var mockClient *mock.MockoxideClient
		BeforeEach(func() {
			mockClient = mock.NewMockoxideClient(ctrl)
			SUT.oxideClient = mockClient
		})

		Describe("Start", func() {
			It("starts successfully", func() {
				mockClient.EXPECT().InstanceStart(gomock.Any(), gomock.Any()).Return(nil, nil)
				Expect(SUT.Start()).To(Succeed())
			})

			It("fails when the machine is not running", func() {
				mockClient.EXPECT().InstanceStart(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("fake error"))
				err := SUT.Start()
				Expect(err).To(HaveOccurred())
			})
		})

		Describe("Stop", func() {
			It("starts successfully", func() {
				mockClient.EXPECT().InstanceStop(gomock.Any(), gomock.Any()).Return(nil, nil)
				Expect(SUT.Stop()).To(Succeed())
			})

			It("fails when the machine is not running", func() {
				mockClient.EXPECT().InstanceStop(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("fake error"))
				err := SUT.Stop()
				Expect(err).To(HaveOccurred())
			})
		})
	})
})

func defaultDriverOptions(m *mock.MockDriverOptions) {
	m.EXPECT().String(gomock.Eq(flagHost)).Return("host").AnyTimes()
	m.EXPECT().String(gomock.Eq(flagToken)).Return("token").AnyTimes()
	m.EXPECT().String(gomock.Eq(flagProject)).Return("project").AnyTimes()
	m.EXPECT().String(gomock.Eq(flagBootDiskImageID)).Return("image").AnyTimes()
	m.EXPECT().String(gomock.Any()).Return("").AnyTimes()
	m.EXPECT().StringSlice(gomock.Any()).Return(nil).AnyTimes()
	m.EXPECT().Int(gomock.Any()).Return(0).AnyTimes()
	m.EXPECT().Bool(gomock.Any()).Return(false).AnyTimes()
}
