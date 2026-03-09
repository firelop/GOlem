package handshake

import (
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/server"
)

type Handshake struct {
	ProtocolVersion int32
	Address         string
	Port            uint16
	Intent          int32
}

func (h *Handshake) New() packets.InboundPacket {
	return &Handshake{}
}

func (h *Handshake) Handle(session *server.Session) {
	session.UpdateState(uint8(h.Intent))
}

func (h *Handshake) ID() int32 {
	return 0x00
}

func (h *Handshake) Read(pb *protocol.PacketBuffer) {
	var err error
	h.ProtocolVersion, err = pb.ReadVarInt()
	if err != nil {
		return
	}
	h.Address, err = pb.ReadString()
	if err != nil {
		return
	}
	h.Port, err = pb.ReadUint16()
	if err != nil {
		return
	}

	h.Intent, err = pb.ReadVarInt()
	if err != nil {
		return
	}
}
