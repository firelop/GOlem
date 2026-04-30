package play

import (
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/server"
)

type ClientTickEnd struct {
}

func (c *ClientTickEnd) New() packets.InboundPacket {
	return &ClientTickEnd{}
}

func (c *ClientTickEnd) ID() int32 {
	return 0x00
}

func (c *ClientTickEnd) Read(pb *protocol.PacketBuffer) {

}

func (c *ClientTickEnd) Handle(session *server.Session) {
}
