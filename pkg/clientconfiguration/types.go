package clientconfiguration

import (
	"github.com/carbonteq/protocol/livekit"
)

type ClientConfigurationManager interface {
	GetConfiguration(clientInfo *livekit.ClientInfo) *livekit.ClientConfiguration
}
