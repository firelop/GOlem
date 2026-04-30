package configuration

import (
	"github.com/firelop/GOlem/logger"
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/server"
)

type PluginMessage struct {
	Channel string
	Data    []byte
}

func (p *PluginMessage) New() packets.InboundPacket {
	return &PluginMessage{}
}

func (p *PluginMessage) ID() int32 {
	return 0x02
}

func (p *PluginMessage) Read(pb *protocol.PacketBuffer) {
	p.Channel, _ = pb.ReadString()
	p.Data = pb.Bytes()
}

func (p *PluginMessage) Handle(session *server.Session) {
	logger.Debug(logger.Protocol, "Plugin message received: ", p.Channel)
	logger.Debug(logger.Protocol, "Data: ", string(p.Data))
	// TODO: Handle plugin message
}
