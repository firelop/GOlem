package outbound

import (
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/server"
)

type PongResponse struct {
	Timestamp int64
}

func NewPongResponse(timestamp int64) *PongResponse {
	return &PongResponse{Timestamp: timestamp}
}

func (p *PongResponse) ID() int32 {
	return 0x01
}

func (p *PongResponse) New() packets.OutboundPacket {
	return &PongResponse{}
}

func (p *PongResponse) Write(pb *protocol.PacketBuffer) {
	pb.WriteInt64(p.Timestamp)
}

func (p *PongResponse) Handle(session *server.Session) {
}
