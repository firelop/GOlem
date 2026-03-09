package login

import (
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/server"
	"github.com/google/uuid"
)

type Success struct {
	UUID     uuid.UUID
	Username string
}

func (s *Success) New() packets.OutboundPacket {
	return &Success{}
}

func (s *Success) ID() int32 {
	return 0x02
}

func (s *Success) Write(pb *protocol.PacketBuffer) {
	pb.WriteUUID(s.UUID)
	pb.WriteString(s.Username)
	pb.WriteVarInt(0) // Number of properties (0 for now)
}

func (s *Success) Handle(session *server.Session) {
}
