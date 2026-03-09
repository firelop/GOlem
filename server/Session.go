package server

import (
	"log"
	"net"

	"github.com/firelop/GOlem/protocol"
)

const (
	INTENT_HANDSHAKE     = 0
	INTENT_STATUS        = 1
	INTENT_LOGIN         = 2
	INTENT_PLAY          = 3
	INTENT_CONFIGURATION = 4
)

type Session struct {
	Server *Server

	State uint8
	Conn  net.Conn

	ReadPacketBuffer   *protocol.PacketBuffer
	WritePacketBuffer  *protocol.PacketBuffer
	HeaderPacketBuffer *protocol.PacketBuffer
}

func (s *Session) UpdateState(state uint8) {
	if s.Server.Debug {
		log.Println("State updated from", s.State, "to", state)
	}
	s.State = state
}

func NewSession(Server *Server, conn net.Conn) *Session {
	return &Session{
		Server:             Server,
		State:              INTENT_HANDSHAKE,
		Conn:               conn,
		ReadPacketBuffer:   protocol.NewPacketBuffer(make([]byte, 0, 128)),
		WritePacketBuffer:  protocol.NewPacketBuffer(make([]byte, 0, 128)),
		HeaderPacketBuffer: protocol.NewPacketBuffer(make([]byte, 0, 5)),
	}
}
