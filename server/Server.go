package server

import (
	"net"

	"github.com/firelop/GOlem/protocol/datatypes"
)

type Server struct {
	Address string
	Port    uint16
	Conn    net.Conn

	ViewDistance        int32
	SimulationDistance  int32
	MaxPlayers          int32
	EnableRespawnScreen bool
	ReducedDebugInfo    bool
	IsHardcore          bool
	DoLimitedCrafting   bool
	EnforcesSecureChat  bool

	DefaultGameMode byte

	Worlds map[string]*datatypes.World

	Debug bool
}

func NewServer(address string, port uint16) *Server {
	worlds := make(map[string]*datatypes.World)
	worlds["world"] = datatypes.NewWorld("world", "minecraft:overworld", false, 0, *datatypes.NewPosition(0, 64, 0))

	return &Server{
		Address: address,
		Port:    port,
		Debug:   true,

		ViewDistance:        10,
		SimulationDistance:  10,
		MaxPlayers:          2026,
		EnableRespawnScreen: true,
		ReducedDebugInfo:    false,
		IsHardcore:          false,
		DoLimitedCrafting:   false,
		EnforcesSecureChat:  false,

		DefaultGameMode: 0,

		Worlds: worlds,
	}
}
