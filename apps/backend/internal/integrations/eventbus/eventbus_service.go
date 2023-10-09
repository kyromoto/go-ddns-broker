package eventbus

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kyromoto/go-ddns/internal/services/clientmanager"
	"github.com/mustafaturan/bus/v3"
	"github.com/mustafaturan/monoton/v2"
	"github.com/mustafaturan/monoton/v2/sequencer"
	"inet.af/netaddr"
)

func NewInMemory() clientmanager.Eventbus {

	node := uint64(1)
	initTime := uint64(time.Now().Nanosecond())

	m, err := monoton.New(sequencer.NewNanosecond(), node, initTime)

	if err != nil {
		panic(err)
	}

	var idGenerator bus.Next = m.Next

	bus, err := bus.NewBus(idGenerator)

	if err != nil {
		panic(err)
	}

	bus.RegisterTopics("client.ipupdate")

	return &eventbusInMemory{
		bus: bus,
	}
}

type eventbusInMemory struct {
	bus *bus.Bus
}

func (b *eventbusInMemory) PublishClientIpUpdate(clientid uuid.UUID, ip netaddr.IP) error {
	ctx := context.Background()
	data := eventDataClientIpUpdate{
		Clientid: clientid,
		Ip:       ip,
	}

	return b.bus.Emit(ctx, "client.ipupdate", data)
}

func (b *eventbusInMemory) SubscribeClientIpUpdate(subscriber clientmanager.SubscriberClientIpUpdate) error {
	handler := bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			data := e.Data.(eventDataClientIpUpdate)
			subscriber.OnClientIpUpdate(data.Clientid, data.Ip)
		},
		Matcher: "client.ipupdate",
	}

	b.bus.RegisterHandler(uuid.New().String(), handler)

	return nil
}

type eventDataClientIpUpdate struct {
	Clientid uuid.UUID
	Ip       netaddr.IP
}
