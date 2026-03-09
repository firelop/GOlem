package login

import (
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/server"
)

type PluginRequest struct {
	MessageID int32
	Channel   string
	Data      []byte
}

func (p *PluginRequest) New() packets.OutboundPacket {
	return &PluginRequest{}
}

func (p *PluginRequest) ID() int32 {
	return 0x04
}

func (p *PluginRequest) Write(pb *protocol.PacketBuffer) {
	pb.WriteVarInt(p.MessageID)
	pb.WriteString(p.Channel)
	pb.WriteByteArray(p.Data)
}

func (p *PluginRequest) Handle(session *server.Session) {
}
