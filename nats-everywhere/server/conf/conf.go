package conf

import "gitlab.dev.21vianet.com/liu.hao8/nats-everywhere/server/logging"

// NATSConfig configuration for a NATS connection.
type NATSConfig struct {
	Servers []string

	ConnectTimeout int // milliseconds
	ReconnectWait  int // milliseconds
	MaxReconnects  int
}

// JetStreamConfig configuration for a JetStream connection.
type JetStreamConfig struct {
	Subject string
	Durable string

	MaxWait int // milliseconds
}

// NATSBridgeConfig is the root structure for a bridge configuration file.
type NATSBridgeConfig struct {
	Logging   logging.Config
	NATS      NATSConfig
	JetStream JetStreamConfig
}

// DefaultBridgeConfig generates a default configuration with
// logging set to colors, time, debug and trace
func DefaultNatsBridgeConfig() NATSBridgeConfig {
	return NATSBridgeConfig{
		Logging: logging.Config{
			Time:   true,
			Debug:  false,
			Trace:  false,
			Colors: true,
			PID:    false,
		},
		NATS: NATSConfig{
			ConnectTimeout: 5000,
			ReconnectWait:  1000,
			MaxReconnects:  0,
		},
		JetStream: JetStreamConfig{
			MaxWait: 5000,
		},
	}
}
