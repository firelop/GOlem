package login

import (
	"github.com/firelop/GOlem/logger"

	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/packets"
	outbound "github.com/firelop/GOlem/protocol/packets/outbound/login"
	"github.com/firelop/GOlem/server"
	"github.com/google/uuid"
)

type LoginStart struct {
	Username string
	UUID     uuid.UUID
}

func (l *LoginStart) New() packets.InboundPacket {
	return &LoginStart{}
}

func (l *LoginStart) ID() int32 {
	return 0x00
}

func (l *LoginStart) Read(pb *protocol.PacketBuffer) {
	l.Username, _ = pb.ReadString()
	l.UUID, _ = pb.ReadUUID()
}

func (l *LoginStart) Handle(session *server.Session) {
	logger.Debug(logger.Protocol, "Login start for ", l.Username, " with UUID ", l.UUID)

	successPacket := &outbound.Success{
		UUID:     l.UUID,
		Username: l.Username,
	}
	packets.SendPacket(session, successPacket) // Offline mode: just accept TODO: Handle online mode

	// TODO: Save those settings in the session
}
