package inbound

import (
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/protocol/packets/inbound/configuration"
	"github.com/firelop/GOlem/protocol/packets/inbound/handshake"
	"github.com/firelop/GOlem/protocol/packets/inbound/login"
	"github.com/firelop/GOlem/protocol/packets/inbound/play"
	"github.com/firelop/GOlem/protocol/packets/inbound/status"
	"github.com/firelop/GOlem/server"
)

var inboundPackets = map[uint8]map[int32]packets.InboundPacket{
	server.INTENT_HANDSHAKE: {
		0x00: &handshake.Handshake{},
	},
	server.INTENT_STATUS: {
		0x00: &status.StatusRequest{},
		0x01: &status.PingRequest{},
	},
	server.INTENT_LOGIN: {
		0x00: &login.LoginStart{},
		0x01: &login.EncryptionResponse{},
		0x02: &login.LoginPluginResponse{},
		0x03: &login.LoginAcknowledged{},
	},
	server.INTENT_CONFIGURATION: {
		0x00: &configuration.ClientInformation{},
		0x02: &configuration.PluginMessage{},
		0x03: &configuration.ACKFinishConfiguration{},
		0x07: &configuration.KnownPacks{},
	},
	server.INTENT_PLAY: {
		0x0B: &play.ClientStatus{},
		0x0C: &play.ClientTickEnd{},
	},
}

func GetNewInboundPacket(state uint8, id int32) packets.InboundPacket {
	// Check if the packet exists
	if _, ok := inboundPackets[state][id]; !ok {
		return nil
	}
	return inboundPackets[state][id].New()
}
