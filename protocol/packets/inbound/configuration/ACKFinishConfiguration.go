package configuration

import (
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/datatypes"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/protocol/packets/outbound/play"
	"github.com/firelop/GOlem/server"
)

type ACKFinishConfiguration struct {
}

func (a *ACKFinishConfiguration) New() packets.InboundPacket {
	return &ACKFinishConfiguration{}
}

func (a *ACKFinishConfiguration) ID() int32 {
	return 0x03
}

func (a *ACKFinishConfiguration) Read(pb *protocol.PacketBuffer) {
}

func (a *ACKFinishConfiguration) Handle(session *server.Session) {
	session.UpdateState(server.INTENT_PLAY)

	server := session.Server
	worlds := server.Worlds

	worldNames := make([]string, len(worlds))

	i := 0
	for k := range worlds {
		worldNames[i] = k
		i++
	}

	world := worlds[worldNames[0]]

	login := &play.Login{
		EntityID:            0,
		IsHardcore:          server.IsHardcore,
		DimensionNames:      worldNames,
		MaxPlayers:          server.MaxPlayers,
		ViewDistance:        server.ViewDistance,
		SimulationDistance:  server.SimulationDistance,
		ReducedDebugInfo:    server.ReducedDebugInfo,
		EnableRespawnScreen: server.EnableRespawnScreen,
		DoLimitedCrafting:   server.DoLimitedCrafting,
		DimensionType:       0,
		DimensionName:       "minecraft:overworld",
		HashedSeed:          0x1212121212121212,
		GameMode:            server.DefaultGameMode,
		PreviousGameMode:    int8(server.DefaultGameMode),
		IsDebug:             false,
		IsFlat:              world.IsFlat,
		HasDeathLocation:    false,
		DeathLocation:       *datatypes.NewPosition(0, 0, 0).ToWorldPosition(world),
		PortalCooldown:      0,
		SeaLevel:            64,
		EnforcesSecureChat:  false,
	}
	packets.SendPacket(session, login)

	packets.SendPacket(session, &play.SetDefaultSpawnPosition{
		WorldPosition: datatypes.NewPosition(0, 100, 0).ToWorldPosition(world),
		Angle:         datatypes.NewAngle(0, 0),
	})

	synch := &play.SynchronizePlayerPosition{
		TeleportID: 0,
		Position:   datatypes.NewDoubleVec3(0, 100, 0),
		Velocity:   datatypes.NewDoubleVec3(0, 0, 0),
		Angle:      datatypes.NewAngle(0, 0),
		Flags:      0,
	}
	packets.SendPacket(session, synch)

}
