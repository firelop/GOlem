package configuration

import (
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/server"
)

type UpdateTags struct {
	Registries []RegistryTagsEntry
}

type RegistryTagsEntry struct {
	Identifier string
	Tags       []TagEntry
}

type TagEntry struct {
	Identifier string
	Tags       []int32
}

func (u *UpdateTags) New() packets.OutboundPacket {
	return &UpdateTags{}
}

func (u *UpdateTags) ID() int32 {
	return 0x0D
}

func (u *UpdateTags) Write(pb *protocol.PacketBuffer) {
	protocol.WritePrefixedArray(pb, u.Registries, func(pb *protocol.PacketBuffer, entry RegistryTagsEntry) {
		pb.WriteString(entry.Identifier)
		protocol.WritePrefixedArray(pb, entry.Tags, func(pb *protocol.PacketBuffer, tag TagEntry) {
			pb.WriteString(tag.Identifier)
			protocol.WritePrefixedArray(pb, tag.Tags, func(pb *protocol.PacketBuffer, tag int32) {
				pb.WriteVarInt(tag)
			})
		})
	})
}

func (u *UpdateTags) Handle(session *server.Session) {
}
