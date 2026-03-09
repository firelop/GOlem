package configuration

import (
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/server"
)

type KnownPacksEntry struct {
	Namespace string
	ID        string
	Version   string
}

type KnownPacks struct {
	Entries []KnownPacksEntry
}

func (k *KnownPacks) New() packets.OutboundPacket {
	return &KnownPacks{}
}

func (k *KnownPacks) ID() int32 {
	return 0x0E
}

func (k *KnownPacks) Write(pb *protocol.PacketBuffer) {
	protocol.WritePrefixedArray(pb, k.Entries, func(pb *protocol.PacketBuffer, entry KnownPacksEntry) {
		pb.WriteString(entry.Namespace)
		pb.WriteString(entry.ID)
		pb.WriteString(entry.Version)
	})
}

func (k *KnownPacks) Handle(session *server.Session) {
}
