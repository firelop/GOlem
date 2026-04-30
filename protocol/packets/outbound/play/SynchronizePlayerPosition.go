package play

import (
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/datatypes"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/server"
)

type SynchronizePlayerPosition struct {
	TeleportID int32
	Position   *datatypes.DoubleVec3
	Velocity   *datatypes.DoubleVec3
	Angle      *datatypes.Angle
	Flags      int32
}

func (s *SynchronizePlayerPosition) New() packets.OutboundPacket {
	return &SynchronizePlayerPosition{}
}

func (s *SynchronizePlayerPosition) ID() int32 {
	return 0x46
}

func (s *SynchronizePlayerPosition) Write(pb *protocol.PacketBuffer) {
	pb.WriteVarInt(s.TeleportID)

	pb.WriteDouble(s.Position.X)
	pb.WriteDouble(s.Position.Y)
	pb.WriteDouble(s.Position.Z)

	pb.WriteDouble(s.Velocity.X)
	pb.WriteDouble(s.Velocity.Y)
	pb.WriteDouble(s.Velocity.Z)

	pb.WriteFloat(s.Angle.Yaw)
	pb.WriteFloat(s.Angle.Pitch)

	pb.WriteInt(s.Flags)
}

func (s *SynchronizePlayerPosition) Handle(session *server.Session) {
}
