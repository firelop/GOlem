package status

import (
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/protocol/packets/outbound"
	"github.com/firelop/GOlem/server"
)

type PingRequest struct {
	Timestamp int64
}

func (p *PingRequest) New() packets.InboundPacket {
	return &PingRequest{}
}

func (p *PingRequest) ID() int32 {
	return 0x01
}

func (p *PingRequest) Read(pb *protocol.PacketBuffer) {
	p.Timestamp, _ = pb.ReadInt64()
}

func (p *PingRequest) Handle(session *server.Session) {
	pongResponse := outbound.NewPongResponse(p.Timestamp)
	packets.SendPacket(session, pongResponse)
}
