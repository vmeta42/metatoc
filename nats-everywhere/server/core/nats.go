package core

import (
	"github.com/nats-io/nats.go"
	"strings"
	"time"
)

func (server *NATSBridge) natsError(nc *nats.Conn, sub *nats.Subscription, err error) {
}

func (server *NATSBridge) natsDiscoveredServers(nc *nats.Conn) {
}

func (server *NATSBridge) natsDisconnected(nc *nats.Conn) {
}

func (server *NATSBridge) natsReconnected(nc *nats.Conn) {
}

func (server *NATSBridge) natsClosed(nc *nats.Conn) {
}

func (server *NATSBridge) connectToNATS() error {
	server.natsLock.Lock()
	defer server.natsLock.Unlock()

	server.logger.Noticef("connecting to NATS core")

	config := server.config.NATS
	options := []nats.Option{
		nats.MaxReconnects(config.MaxReconnects),
		nats.ReconnectWait(time.Duration(config.ReconnectWait) * time.Millisecond),
		nats.Timeout(time.Duration(config.ConnectTimeout) * time.Millisecond),
		nats.ErrorHandler(server.natsError),
		nats.DiscoveredServersHandler(server.natsDiscoveredServers),
		nats.DisconnectHandler(server.natsDisconnected),
		nats.ReconnectHandler(server.natsReconnected),
		nats.ClosedHandler(server.natsClosed),
		nats.NoCallbacksAfterClientClose(),
	}

	nc, err := nats.Connect(strings.Join(config.Servers, ","),
		options...,
	)
	if err != nil {
		return err
	}
	server.nats = nc

	return nil
}

func (server *NATSBridge) connectToJetStream() error {
	server.natsLock.Lock()
	defer server.natsLock.Unlock()

	server.logger.Noticef("connecting to JetStream")

	config := server.config.JetStream
	options := []nats.JSOpt{
		nats.MaxWait(time.Duration(config.MaxWait) * time.Millisecond),
	}

	js, err := server.nats.JetStream(options...)
	if err != nil {
		return err
	}
	server.js = js

	return nil
}
