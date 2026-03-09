package login

import (
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/server"
)

type EncryptionRequest struct {
	ServerID    string
	PublicKey   []byte
	VerifyToken []byte
}

func (e *EncryptionRequest) New() packets.OutboundPacket {
	return &EncryptionRequest{}
}

func (e *EncryptionRequest) ID() int32 {
	return 0x01
}

func (e *EncryptionRequest) Write(pb *protocol.PacketBuffer) {
	pb.WriteString(e.ServerID)
	pb.WriteByteArray(e.PublicKey)
	pb.WriteByteArray(e.VerifyToken)
}

func (e *EncryptionRequest) Handle(session *server.Session) {
}
