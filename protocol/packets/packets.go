package packets

import (
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/server"
)

type InboundPacket interface {
	ID() int32
	Read(pb *protocol.PacketBuffer)
	New() InboundPacket
	Handle(session *server.Session)
}

type OutboundPacket interface {
	ID() int32
	Write(pb *protocol.PacketBuffer)
	New() OutboundPacket
	Handle(session *server.Session) // Not used or neither called yet
}

func SendPacket(s *server.Session, packet OutboundPacket) {
	s.WritePacketBuffer.Reset()
	s.WritePacketBuffer.WriteVarInt(packet.ID())
	packet.Write(s.WritePacketBuffer)

	s.HeaderPacketBuffer.Reset()
	s.HeaderPacketBuffer.WriteVarInt(int32(len(s.WritePacketBuffer.Bytes())))

	s.Conn.Write(s.HeaderPacketBuffer.Bytes())
	s.Conn.Write(s.WritePacketBuffer.Bytes())
}
