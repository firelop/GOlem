package configuration

import (
	"github.com/firelop/GOlem/logger"
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/protocol/packets/outbound/configuration"
	"github.com/firelop/GOlem/server"
)

type KnownPacks struct {
	Entries []configuration.KnownPacksEntry
}

func (k *KnownPacks) New() packets.InboundPacket {
	return &KnownPacks{}
}

func (k *KnownPacks) ID() int32 {
	return 0x07
}

func (k *KnownPacks) Read(pb *protocol.PacketBuffer) {
	k.Entries, _ = protocol.ReadPrefixedArray(pb, func(pb *protocol.PacketBuffer) (configuration.KnownPacksEntry, error) {
		namespace, _ := pb.ReadString()
		id, _ := pb.ReadString()
		version, _ := pb.ReadString()
		return configuration.KnownPacksEntry{
			Namespace: namespace,
			ID:        id,
			Version:   version,
		}, nil
	})
}

func (k *KnownPacks) Handle(session *server.Session) {
	for _, entry := range k.Entries {
		logger.Debug(logger.Protocol, "(Client) Known pack: ", entry.Namespace, entry.ID, entry.Version)
	}
}
