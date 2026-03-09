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
		DimensionType:       "",
		DimensionName:       "",
		HashedSeed:          0,
		GameMode:            server.DefaultGameMode,
		PreviousGameMode:    -1,
		IsDebug:             false,
		IsFlat:              world.IsFlat,
		HasDeathLocation:    false,
		DeathLocation:       *datatypes.NewPosition(0, 0, 0).ToWorldPosition(world),
		PortalCooldown:      0,
		EnforcesSecureChat:  false,
	}
	packets.SendPacket(session, login)

}
