package status

import (
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/protocol/packets/outbound"
	"github.com/firelop/GOlem/server"
)

type StatusRequest struct {
}

func (s *StatusRequest) New() packets.InboundPacket {
	return &StatusRequest{}
}

func (s *StatusRequest) ID() int32 {
	return 0x00
}

func (s *StatusRequest) Read(pb *protocol.PacketBuffer) {
}

func (s *StatusRequest) Handle(session *server.Session) {
	statusResponse := outbound.NewStatusResponse("\u00a72GOlem runtime\u00a76 v0.0.1 (alpha)\u00a7r\n  \u00a7e\u1405 Lightweight golang minecraft server", 774, "1.21.11", int(session.Server.MaxPlayers), 0)
	packets.SendPacket(session, statusResponse)
}
