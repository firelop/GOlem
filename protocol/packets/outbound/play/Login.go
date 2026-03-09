package play

import (
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/datatypes"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/server"
)

// Login (play) packet implementation for Minecraft 1.20.5+ (ID 0x2B)
// See: https://minecraft.wiki/w/Java_Edition_protocol/Packets#Login_(play)
type Login struct {
	EntityID            int32
	IsHardcore          bool
	DimensionNames      []string
	MaxPlayers          int32
	ViewDistance        int32
	SimulationDistance  int32
	ReducedDebugInfo    bool
	EnableRespawnScreen bool
	DoLimitedCrafting   bool
	DimensionType       string // Identifier
	DimensionName       string // Identifier
	HashedSeed          int64
	GameMode            byte
	PreviousGameMode    int8 // -1 for none
	IsDebug             bool
	IsFlat              bool
	HasDeathLocation    bool
	DeathLocation       datatypes.WorldPosition
	PortalCooldown      int32
	EnforcesSecureChat  bool
}

func (l *Login) New() packets.OutboundPacket {
	return &Login{}
}

func (l *Login) ID() int32 {
	return 0x30
}

func (l *Login) Write(pb *protocol.PacketBuffer) {
	pb.WriteInt32(l.EntityID)
	pb.WriteBool(l.IsHardcore)

	// Dimension Names: Array of Identifiers
	protocol.WritePrefixedArray(pb, l.DimensionNames, (*protocol.PacketBuffer).WriteString)

	pb.WriteVarInt(l.MaxPlayers)
	pb.WriteVarInt(l.ViewDistance)
	pb.WriteVarInt(l.SimulationDistance)
	pb.WriteBool(l.ReducedDebugInfo)
	pb.WriteBool(l.EnableRespawnScreen)
	pb.WriteBool(l.DoLimitedCrafting)

	pb.WriteString(l.DimensionType)
	pb.WriteString(l.DimensionName)
	pb.WriteInt64(l.HashedSeed)
	pb.WriteByte(l.GameMode)
	pb.WriteByte(byte(l.PreviousGameMode))
	pb.WriteBool(l.IsDebug)
	pb.WriteBool(l.IsFlat)

	pb.WriteBool(l.HasDeathLocation)
	if l.HasDeathLocation {
		pb.WriteString(l.DeathLocation.World.Name)
		pb.WritePosition(l.DeathLocation.Position)
	}

	pb.WriteVarInt(l.PortalCooldown)
	pb.WriteBool(l.EnforcesSecureChat)
}

func (l *Login) Handle(session *server.Session) {
	// Handle is not used for outbound packets in this implementation
}
