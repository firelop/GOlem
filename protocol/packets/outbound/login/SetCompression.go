package login

import (
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/server"
)

type SetCompression struct {
	Threshold int32
}

func (s *SetCompression) New() packets.OutboundPacket {
	return &SetCompression{}
}

func (s *SetCompression) ID() int32 {
	return 0x03
}

func (s *SetCompression) Write(pb *protocol.PacketBuffer) {
	pb.WriteVarInt(s.Threshold)
}

func (s *SetCompression) Handle(session *server.Session) {
}
