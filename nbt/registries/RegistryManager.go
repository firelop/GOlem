package registries

import (
	"github.com/firelop/GOlem/logger"
	"github.com/firelop/GOlem/protocol/packets/outbound/configuration"
	"github.com/sandertv/gophertunnel/minecraft/nbt"
)

func RegistryStructToRegistryData(registryId string, registryData map[string]any) configuration.RegistryData {
	entries := make([]configuration.RegistryDataEntry, len(registryData))

	i := 0
	for k, v := range registryData {
		data, err := nbt.MarshalEncoding(v, nbt.NetworkBigEndian)
		if err != nil {
			logger.Error("Error marshalling registry data: ", err)
			continue
		}
		entries[i] = configuration.RegistryDataEntry{
			Name:      k,
			IsPresent: true,
			Value:     data,
		}
		i++
	}

	return configuration.RegistryData{
		RegistryID: registryId,
		Entries:    entries,
	}
}
