package filesync

import (
	"errors"
	"fmt"
	"time"

	"github.com/nats-io/jsm.go"
	jsmapi "github.com/nats-io/jsm.go/api"
	"github.com/nats-io/nats.go"
)

type JsmClient struct {
	nc *nats.Conn
	jm *jsm.Manager
}

func (c *JsmClient) Close() {
	c.nc.Drain()
}

func (c *JsmClient) LoadStream(name string) (*jsm.Stream, error) {
	return c.jm.LoadStream(name)
}

func (c *JsmClient) NewStream(name string, opts []jsm.StreamOption) (*jsm.Stream, error) {
	return c.jm.NewStream(name, opts...)
}

func (c *JsmClient) LoadConsumer(stream, consumer string) (*jsm.Consumer, error) {
	return c.jm.LoadConsumer(stream, consumer)
}

func (c *JsmClient) NewConsumer(stream string, opts []jsm.ConsumerOption) (*jsm.Consumer, error) {
	return c.jm.NewConsumer(stream, opts...)
}

func (c *JsmClient) CreateDefaultStreamIfNotExist(streamName string, subjects ...string) (created bool, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("failed to create stream %q: %w", streamName, err)
		}
	}()
	var apierr jsmapi.ApiError
	_, err = c.LoadStream(streamName)
	if errors.As(err, &apierr) && apierr.NotFoundError() {
		opts := []jsm.StreamOption{
			jsm.Subjects(subjects...),
			jsm.MaxConsumers(-1),
			jsm.MaxMessageSize(-1),
			jsm.MaxMessages(-1),
			jsm.Replicas(1),
			jsm.DuplicateWindow(2 * time.Second),
			jsm.MaxAge(0),
			jsm.WorkQueueRetention(),
			jsm.MemoryStorage(),
			jsm.DiscardOld(),
		}
		_, err = c.NewStream(streamName, opts)
		return true, err
	} else if err != nil {
		return false, err
	}
	return false, nil
}

func (c *JsmClient) CreateDefaultConsumerIfNotExist(streamName, consumerName, filterSubject string) (created bool, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("failed to create consumer %q on stream %q: %w", consumerName, streamName, err)
		}
	}()

	var apierr jsmapi.ApiError
	_, err = c.LoadConsumer(streamName, consumerName)
	if errors.As(err, &apierr) && apierr.NotFoundError() {
		opts := []jsm.ConsumerOption{
			jsm.DurableName(consumerName),
			jsm.DeliverAllAvailable(),
			jsm.ReplayInstantly(),
			jsm.FilterStreamBySubject(filterSubject),
			jsm.MaxAckPending(uint(1)),
			jsm.MaxDeliveryAttempts(-1),
		}
		_, err = c.NewConsumer(streamName, opts)
		return true, err
	} else if err != nil {
		return false, err
	}

	return false, nil
}

func NewJsmClient(servers string, opts ...nats.Option) (*JsmClient, error) {
	jc := &JsmClient{}
	nc, err := nats.Connect(servers, opts...)
	if err != nil {
		return nil, err
	}
	jc.nc = nc

	m, err := jsm.New(nc)
	if err != nil {
		return nil, err
	}
	jc.jm = m

	return jc, nil
}
