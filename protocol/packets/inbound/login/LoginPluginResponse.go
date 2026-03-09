package login

import (
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/server"
)

type LoginPluginResponse struct {
	MessageID  int32
	Successful bool
	Data       []byte
}

func (l *LoginPluginResponse) New() packets.InboundPacket {
	return &LoginPluginResponse{}
}

func (l *LoginPluginResponse) ID() int32 {
	return 0x02
}

func (l *LoginPluginResponse) Read(pb *protocol.PacketBuffer) {
	l.MessageID, _ = pb.ReadVarInt()
	l.Successful, _ = pb.ReadBool()
	if l.Successful {
		// Reading the rest of the packet as data if successful?
		// Note: Wiki says "Optional", but usually implies "until end of packet" or length prefixed if it's a byte array field type.
		// However, in our PacketBuffer ReadByteArray expects a length prefix.
		// For LoginPluginResponse: "Data: Byte Array. The data to send to the server. Only present if Successful is true."
		// Protocol usually uses standard types. "Byte Array" = VarInt Length + Bytes.
		// So we use ReadByteArray.
		l.Data, _ = pb.ReadByteArray()
	}
}

func (l *LoginPluginResponse) Handle(session *server.Session) {
	// TODO: Handle plugin response
}
