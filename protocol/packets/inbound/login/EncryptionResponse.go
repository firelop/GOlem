package login

import (
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/server"
)

type EncryptionResponse struct {
	SharedSecret []byte
	VerifyToken  []byte
}

func (e *EncryptionResponse) New() packets.InboundPacket {
	return &EncryptionResponse{}
}

func (e *EncryptionResponse) ID() int32 {
	return 0x01
}

func (e *EncryptionResponse) Read(pb *protocol.PacketBuffer) {
	e.SharedSecret, _ = pb.ReadByteArray()
	e.VerifyToken, _ = pb.ReadByteArray()
}

func (e *EncryptionResponse) Handle(session *server.Session) {
	// TODO: Handle encryption
}
