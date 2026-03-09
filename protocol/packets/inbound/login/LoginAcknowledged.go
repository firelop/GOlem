package login

import (
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/server"
)

type LoginAcknowledged struct {
}

func (l *LoginAcknowledged) New() packets.InboundPacket {
	return &LoginAcknowledged{}
}

func (l *LoginAcknowledged) ID() int32 {
	return 0x03
}

func (l *LoginAcknowledged) Read(pb *protocol.PacketBuffer) {
	// Empty packet
}

func (l *LoginAcknowledged) Handle(session *server.Session) {
	// Switch to Configuration state
	session.UpdateState(server.INTENT_CONFIGURATION)
}
