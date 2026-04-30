package configuration

import (
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/server"
)

type RegistryDataEntry struct {
	Name      string
	IsPresent bool
	Value     []byte
}

type RegistryData struct {
	RegistryID string
	Entries    []RegistryDataEntry
}

func (r *RegistryData) New() packets.OutboundPacket {
	return &RegistryData{}
}

func (r *RegistryData) ID() int32 {
	return 0x07
}

func (r *RegistryData) Write(pb *protocol.PacketBuffer) {
	pb.WriteString(r.RegistryID)
	protocol.WritePrefixedArray(pb, r.Entries, func(pb *protocol.PacketBuffer, entry RegistryDataEntry) {
		pb.WriteString(entry.Name)
		pb.WriteBool(entry.IsPresent)
		if entry.IsPresent {
			pb.Write(entry.Value)
		}
	})
}

func (r *RegistryData) Handle(session *server.Session) {
}
