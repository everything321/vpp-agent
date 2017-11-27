// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vppcalls

import (
	"testing"

	govppapi "git.fd.io/govpp.git/api"
	"github.com/ligato/cn-infra/logging/logrus"
	l2ba "github.com/ligato/vpp-agent/plugins/defaultplugins/l2plugin/bin_api/l2"
	"github.com/ligato/vpp-agent/plugins/defaultplugins/l2plugin/vppcalls/test/impl"
	"github.com/onsi/gomega"
)

const (
	dummyBridgeDomain uint32 = 4
	dummyMACAddress          = "FF:FF:FF:FF:FF:FF"
	dummyIPAddress           = "192.168.4.4"
	dummyLoggerName          = "dummyLogger"
)

//TestVppAddArpTerminationTableEntry tests VppAddArpTerminationTableEntry
func TestVppAddArpTerminationTableEntry(t *testing.T) {
	gomega.RegisterTestingT(t)
	mockedChannel := &impl.MockedChannel{Channel: govppapi.Channel{}}
	VppAddArpTerminationTableEntry(dummyBridgeDomain, dummyMACAddress, dummyIPAddress, logrus.NewLogger(dummyLoggerName),
		mockedChannel, nil)

	var vppMsg l2ba.BdIPMacAddDel
	switch t := mockedChannel.Msg.(type) {
	case *l2ba.BdIPMacAddDel:
		vppMsg = *t
	}

	gomega.Expect(vppMsg).NotTo(gomega.BeNil())
	//check values which will be send to VPP
	gomega.Expect(vppMsg.BdID).To(gomega.Equal(dummyBridgeDomain))
	gomega.Expect(vppMsg.MacAddress).To(gomega.Equal([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}))
	gomega.Expect(vppMsg.IPAddress).To(gomega.Equal([]byte{192, 168, 4, 4}))
	gomega.Expect(vppMsg.IsIpv6).To(gomega.Equal(uint8(0)))
	gomega.Expect(vppMsg.IsAdd).To(gomega.Equal(uint8(1)))
}

// TestVppRemoveArpTerminationTableEntry tests VppRemoveArpTerminationTableEntry method
func TestVppRemoveArpTerminationTableEntry(t *testing.T) {
	gomega.RegisterTestingT(t)
	mockedChannel := &impl.MockedChannel{Channel: govppapi.Channel{}}
	VppRemoveArpTerminationTableEntry(dummyBridgeDomain, dummyMACAddress, dummyIPAddress, logrus.NewLogger(dummyLoggerName),
		mockedChannel, nil)
	var vppMsg l2ba.BdIPMacAddDel
	switch t := mockedChannel.Msg.(type) {
	case *l2ba.BdIPMacAddDel:
		vppMsg = *t
	}

	gomega.Expect(vppMsg).NotTo(gomega.BeNil())
	//check values which will be send to VPP
	gomega.Expect(vppMsg.BdID).To(gomega.Equal(dummyBridgeDomain))
	gomega.Expect(vppMsg.MacAddress).To(gomega.Equal([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}))
	gomega.Expect(vppMsg.IPAddress).To(gomega.Equal([]byte{192, 168, 4, 4}))
	gomega.Expect(vppMsg.IsIpv6).To(gomega.Equal(uint8(0)))
	gomega.Expect(vppMsg.IsAdd).To(gomega.Equal(uint8(0)))
}
