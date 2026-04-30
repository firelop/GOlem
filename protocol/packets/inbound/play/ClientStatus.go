package play

import (
	"github.com/firelop/GOlem/logger"
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/server"
)

type ClientStatus struct {
	Status int32
}

func (e *ClientStatus) New() packets.InboundPacket {
	return &ClientStatus{}
}

func (e *ClientStatus) ID() int32 {
	return 0x01
}

func (e *ClientStatus) Read(pb *protocol.PacketBuffer) {
	e.Status, _ = pb.ReadVarInt()
}

func (e *ClientStatus) Handle(session *server.Session) {
	switch e.Status {
	case 0:
		logger.Debug(logger.Protocol, "Client status: Respawn")
	case 1:
		logger.Debug(logger.Protocol, "Client status: Request stats")
	}
}
