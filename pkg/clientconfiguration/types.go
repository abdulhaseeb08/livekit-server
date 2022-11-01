package clientconfiguration

import (
	"github.com/abdulhaseeb08/protocol/livekit"
)

type ClientConfigurationManager interface {
	GetConfiguration(clientInfo *livekit.ClientInfo) *livekit.ClientConfiguration
}
