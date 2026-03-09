package outbound

import (
	"encoding/json"

	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/server"
)

type StatusResponse struct {
	Content struct {
		Version struct {
			Name     string `json:"name"`
			Protocol int    `json:"protocol"`
		} `json:"version"`
		Players struct {
			Max    int   `json:"max"`
			Online int   `json:"online"`
			Sample []int `json:"sample"`
		} `json:"players"`
		Description struct {
			Text string `json:"text"`
		} `json:"description"`
	}
}

func NewStatusResponse(motd string, protocolVersion int, versionName string, maxPlayers int, onlinePlayers int) *StatusResponse {
	s := &StatusResponse{}
	s.Content.Description.Text = motd
	s.Content.Version.Protocol = protocolVersion
	s.Content.Version.Name = versionName
	s.Content.Players.Max = maxPlayers
	s.Content.Players.Online = onlinePlayers
	s.Content.Players.Sample = []int{}
	return s
}

func (p *StatusResponse) ID() int32 {
	return 0x00
}

func (p *StatusResponse) New() packets.OutboundPacket {
	return &StatusResponse{}
}

func (p *StatusResponse) Write(pb *protocol.PacketBuffer) {
	jsonData, _ := json.Marshal(p.Content)
	pb.WriteString(string(jsonData))
}

func (p *StatusResponse) Handle(session *server.Session) {
}
