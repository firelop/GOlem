package play

import (
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/datatypes"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/server"
)

type SetDefaultSpawnPosition struct {
	WorldPosition *datatypes.WorldPosition
	Angle         *datatypes.Angle
}

func (s *SetDefaultSpawnPosition) New() packets.OutboundPacket {
	return &SetDefaultSpawnPosition{}
}

func (s *SetDefaultSpawnPosition) ID() int32 {
	return 0x5F
}

func (s *SetDefaultSpawnPosition) Write(pb *protocol.PacketBuffer) {
	pb.WriteString(s.WorldPosition.World.DimensionType)
	pb.WritePosition(s.WorldPosition.Position)
	pb.WriteFloat(s.Angle.Yaw)
	pb.WriteFloat(s.Angle.Pitch)
}

func (s *SetDefaultSpawnPosition) Handle(session *server.Session) {
}
