package configuration

import (
	"encoding/base64"
	"log"

	nbtstructs "github.com/firelop/GOlem/nbt"
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/protocol/packets/outbound/configuration"
	"github.com/firelop/GOlem/server"
	"github.com/sandertv/gophertunnel/minecraft/nbt"
)

type ClientInformation struct {
	Locale               string
	ViewDistance         byte
	ChatMode             int32
	ChatColors           bool
	SkinParts            byte
	MainHand             int32
	TextFilteringEnabled bool
	AllowServerListing   bool
	ParticleStatus       int32
}

func (c *ClientInformation) New() packets.InboundPacket {
	return &ClientInformation{}
}

func (c *ClientInformation) ID() int32 {
	return 0x00
}

func (c *ClientInformation) Read(pb *protocol.PacketBuffer) {
	c.Locale, _ = pb.ReadString()
	c.ViewDistance, _ = pb.ReadByte()
	c.ChatMode, _ = pb.ReadVarInt()
	c.ChatColors, _ = pb.ReadBool()
	c.SkinParts, _ = pb.ReadByte()
	c.MainHand, _ = pb.ReadVarInt()
	c.TextFilteringEnabled, _ = pb.ReadBool()
	c.AllowServerListing, _ = pb.ReadBool()
	c.ParticleStatus, _ = pb.ReadVarInt()
}

func (c *ClientInformation) Handle(session *server.Session) {
	// TODO: Save those settings in the session

	packets.SendPacket(session, &configuration.KnownPacks{
		Entries: []configuration.KnownPacksEntry{
			{
				Namespace: "minecraft",
				ID:        "core",
				Version:   "1.21.11",
			},
		},
	})

	// Send registry data

	data, err := nbt.MarshalEncoding(nbtstructs.NewOverworldDimensionType(), nbt.NetworkBigEndian)

	var dimensionType nbtstructs.DimensionType
	err = nbt.UnmarshalEncoding(data, &dimensionType, nbt.NetworkBigEndian)

	log.Println("NBT", dimensionType)

	if err != nil {
		panic(err)
	}

	// Print data as base64
	log.Println(base64.StdEncoding.EncodeToString(data))

	dimensionTypes := make([]configuration.RegistryDataEntries, len(session.Server.Worlds))

	i := 0
	for k := range session.Server.Worlds {
		// TODO: Customize dimension types

		dimensionTypes[i] = configuration.RegistryDataEntries{
			Name:      k,
			IsPresent: true,
			Value:     data,
		}
		i++
	}

	packets.SendPacket(session, &configuration.RegistryData{
		RegistryID: "minecraft:dimension_type",
		Entries:    dimensionTypes,
	})

	// Required mob variant registry
	packets.SendPacket(session, &configuration.RegistryData{
		RegistryID: "minecraft:cat_variant",
		Entries: []configuration.RegistryDataEntries{
			{
				Name:      "minecraft:tabby",
				IsPresent: false,
				Value:     []byte{},
			},
		},
	})

	packets.SendPacket(session, &configuration.RegistryData{
		RegistryID: "minecraft:cow_variant",
		Entries: []configuration.RegistryDataEntries{
			{
				Name:      "minecraft:cow",
				IsPresent: false,
				Value:     []byte{},
			},
		},
	})

	packets.SendPacket(session, &configuration.RegistryData{
		RegistryID: "minecraft:chicken_variant",
		Entries: []configuration.RegistryDataEntries{
			{
				Name:      "minecraft:chicken",
				IsPresent: false,
				Value:     []byte{},
			},
		},
	})

	packets.SendPacket(session, &configuration.RegistryData{
		RegistryID: "minecraft:frog_variant",
		Entries: []configuration.RegistryDataEntries{
			{
				Name:      "minecraft:frog",
				IsPresent: false,
				Value:     []byte{},
			},
		},
	})

	packets.SendPacket(session, &configuration.RegistryData{
		RegistryID: "minecraft:pig_variant",
		Entries: []configuration.RegistryDataEntries{
			{
				Name:      "minecraft:pig",
				IsPresent: false,
				Value:     []byte{},
			},
		},
	})

	packets.SendPacket(session, &configuration.RegistryData{
		RegistryID: "minecraft:wolf_variant",
		Entries: []configuration.RegistryDataEntries{
			{
				Name:      "minecraft:pale",
				IsPresent: false,
				Value:     []byte{},
			},
		},
	})

	packets.SendPacket(session, &configuration.RegistryData{
		RegistryID: "minecraft:zombie_nautilus_variant",
		Entries: []configuration.RegistryDataEntries{
			{
				Name:      "minecraft:zombie_nautilus",
				IsPresent: false,
				Value:     []byte{},
			},
		},
	})

	packets.SendPacket(session, &configuration.RegistryData{
		RegistryID: "minecraft:painting_variant",
		Entries: []configuration.RegistryDataEntries{
			{
				Name:      "minecraft:kebab",
				IsPresent: false,
				Value:     []byte{},
			},
		},
	})

	// Send finish configuration
	packets.SendPacket(session, &configuration.FinishConfiguration{})

}
