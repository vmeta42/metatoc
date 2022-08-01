package main

import (
	"gitlab.dev.21vianet.com/liu.hao8/nats-everywhere/server/conf"
	"gitlab.dev.21vianet.com/liu.hao8/nats-everywhere/server/core"
	"log"
	"os"
)

var (
	natsUrl     string
	natsSubject string
	natsDurable string
)

func init() {
	if os.Getenv("NATS_URL") != "" {
		natsUrl = os.Getenv("NATS_URL")
	}

	if os.Getenv("NATS_SUBJECT") != "" {
		natsSubject = os.Getenv("NATS_SUBJECT")
	}

	if os.Getenv("NATS_DURABLE") != "" {
		natsDurable = os.Getenv("NATS_DURABLE")
	}
}

func main() {
	if natsUrl == "" {
		panic("NATS URL cannot be empty")
	}
	if natsSubject == "" {
		panic("NATS subject cannot be empty")
	}
	if natsDurable == "" {
		panic("NATS durable cannot be empty")
	}

	log.Println("natsUrl:", natsUrl)
	log.Println("natsSubject:", natsSubject)
	log.Println("natsDurable:", natsDurable)

	config := conf.DefaultNatsBridgeConfig()

	log.Println("config.Logging.Time:", config.Logging.Time)
	log.Println("config.Logging.Debug:", config.Logging.Debug)
	log.Println("config.Logging.Trace:", config.Logging.Trace)
	log.Println("config.Logging.Colors:", config.Logging.Colors)
	log.Println("config.Logging.PID:", config.Logging.PID)
	log.Println("config.NATS.ConnectTimeout:", config.NATS.ConnectTimeout)
	log.Println("config.NATS.ReconnectWait: ", config.NATS.ReconnectWait)
	log.Println("config.NATS.MaxReconnects:", config.NATS.MaxReconnects)
	log.Println("config.JetStream.MaxWait:", config.JetStream.MaxWait)

	config.NATS = conf.NATSConfig{
		Servers: []string{natsUrl},

		ConnectTimeout: config.NATS.ConnectTimeout,
		ReconnectWait:  config.NATS.ReconnectWait,
		MaxReconnects:  config.NATS.MaxReconnects,
	}
	config.JetStream = conf.JetStreamConfig{
		Subject: natsSubject,
		Durable: natsDurable,

		MaxWait: config.JetStream.MaxWait,
	}
	bridge := core.NewNATSBridge()
	bridge.InitializeFromConfig(config)
	bridge.Start()
}
