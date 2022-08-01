package core

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"gitlab.dev.21vianet.com/liu.hao8/nats-everywhere/server/conf"
	"gitlab.dev.21vianet.com/liu.hao8/nats-everywhere/server/logging"
	"sync"
	"time"
)

// Version specifies the command version. This should be set at compile time.
var Version = "0.1-dev"

type NATSBridge struct {
	sync.Mutex
	natsLock sync.Mutex

	startTime time.Time

	config conf.NATSBridgeConfig

	logger logging.Logger
	nats   *nats.Conn
	js     nats.JetStreamContext
}

// NewNATSBridge creates a new account server with a default logger
func NewNATSBridge() *NATSBridge {
	return &NATSBridge{}
}

// InitializeFromConfig initialize the server's configuration to an existing config object, useful for tests
// Does not change the config at all, use DefaultServerConfig() to create a default config
func (server *NATSBridge) InitializeFromConfig(config conf.NATSBridgeConfig) error {
	server.config = config
	return nil
}

// Start the server, will lock the server, assumes the config is loaded
func (server *NATSBridge) Start() error {
	server.Lock()
	defer server.Unlock()

	server.startTime = time.Now()

	server.logger = logging.NewNATSLogger(server.config.Logging)
	server.logger.Noticef("starting NATS Bridge, version %s", Version)
	server.logger.Noticef("server time is %s", server.startTime.Format(time.UnixDate))

	if err := server.connectToNATS(); err != nil {
		return err
	}
	if err := server.connectToJetStream(); err != nil {
		return err
	}

	server.logger.Noticef("pull subscribe")
	sub, err := server.js.PullSubscribe(server.config.JetStream.Subject, server.config.JetStream.Durable)
	if err != nil {
		server.logger.Errorf("pull subscribe failure, %s", err.Error())
	} else {
		server.logger.Noticef("pull subscribe success")
		server.logger.Noticef("wait for next message")
		for {
			message, _ := sub.Fetch(1)
			if len(message) > 0 {
				server.logger.Noticef("fet a new message, %s", string(message[0].Data))
				server.logger.Noticef("start consumption")
				var messageDataMap map[string]interface{}
				err = json.Unmarshal([]byte(message[0].Data), &messageDataMap)
				if err != nil {
					server.logger.Errorf("consumption failure, %s", err.Error())
				} else {
					switch messageDataMap["consumer"] {
					case "sendRequest":
						if err = server.sendRequest(messageDataMap); err != nil {
							server.logger.Errorf("consumption failure, %s", err.Error())
						} else {
							server.logger.Noticef("consumption success")
						}
					}
				}
				if err = message[0].AckSync(); err != nil {
					server.logger.Errorf("ack sync failure, %s", err.Error())
				} else {
					server.logger.Noticef("ack sync success")
				}
				server.logger.Noticef("wait for next message")
			}
		}
	}

	return nil
}
