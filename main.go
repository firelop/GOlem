package main

import (
	"bufio"
	"errors"
	"io"
	"net"
	"os"
	"strconv"

	"github.com/firelop/GOlem/logger"
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/packets/inbound"
	"github.com/firelop/GOlem/server"
)

func main() {
	s := server.NewServer("0.0.0.0", 25565)
	start(s)
}

func start(s *server.Server) {
	listener, err := net.Listen("tcp", s.Address+":"+strconv.Itoa(int(s.Port)))
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	defer listener.Close()

	logger.Info("Server started on ", s.Address, ":", s.Port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error(err)
			continue
		}
		go handleConnection(s, conn)
	}

}

func handleConnection(s *server.Server, conn net.Conn) {
	defer conn.Close()

	session := server.NewSession(s, conn)

	reader := bufio.NewReader(conn)
	pb := protocol.NewPacketBuffer(make([]byte, 0, 128))
	for {
		packetLength, err := readVarIntFromConn(reader)

		if err != nil {
			logger.Debug(logger.Network, "Error reading packet length:", err)
			return
		}

		packetData := make([]byte, packetLength)
		_, err = io.ReadFull(reader, packetData)
		if err != nil {
			logger.Error("Error reading packet data:", err)
			return
		}
		pb.Reset()
		pb.Write(packetData)
		packetID, _ := pb.ReadVarInt()

		logger.Debugf(logger.Network, "[->] ID: 0x%X, Size: %d\n", packetID, packetLength)

		packet := inbound.GetNewInboundPacket(session.State, packetID)
		if packet == nil {
			logger.Warn("Packet not found with id: ", packetID, " in state: ", session.State)
			logger.Warn(packetData)
			continue
		}

		packet.Read(pb)
		packet.Handle(session)

	}
}

func readVarIntFromConn(reader io.Reader) (int32, error) {
	var result int32
	var shift uint

	b := make([]byte, 1)

	for {
		_, err := reader.Read(b)
		if err != nil {
			return 0, err
		}

		currentByte := b[0]
		result |= int32(currentByte&0x7F) << shift

		if (currentByte & 0x80) == 0 {
			return result, nil
		}

		shift += 7

		if shift >= 35 {
			return 0, errors.New("VarInt is too large")
		}
	}
}
