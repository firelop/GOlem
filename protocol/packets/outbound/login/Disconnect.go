package login

import (
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/server"
)

type Disconnect struct {
	Reason string // JSON Chat Component
}

func (d *Disconnect) New() packets.OutboundPacket {
	return &Disconnect{}
}

func (d *Disconnect) ID() int32 {
	return 0x00
}

func (d *Disconnect) Write(pb *protocol.PacketBuffer) {
	pb.WriteString(d.Reason)
}

func (d *Disconnect) Handle(session *server.Session) {
}
