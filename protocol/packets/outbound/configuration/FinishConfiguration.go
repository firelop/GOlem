package configuration

import (
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/server"
)

type FinishConfiguration struct {
}

func (f *FinishConfiguration) New() packets.OutboundPacket {
	return &FinishConfiguration{}
}

func (f *FinishConfiguration) ID() int32 {
	return 0x03
}

func (f *FinishConfiguration) Write(pb *protocol.PacketBuffer) {
	// Empty body
}

func (f *FinishConfiguration) Handle(session *server.Session) {
	session.State = server.INTENT_PLAY
}
