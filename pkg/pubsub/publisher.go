package pubsub

import (
	"context"
	"fmt"
	"github.com/RacoonMediaServer/rms-packages/pkg/events"
	"github.com/RacoonMediaServer/rms-packages/pkg/service/servicemgr"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
)

type publisher struct {
	notificationPub micro.Event
	malfunctionPub  micro.Event
	alertPub        micro.Event
}

// NewPublisher creates entity to events publishing
func NewPublisher(s servicemgr.ClientFactory) micro.Event {
	return &publisher{
		notificationPub: micro.NewEvent(NotificationTopic, s.Client()),
		malfunctionPub:  micro.NewEvent(MalfunctionTopic, s.Client()),
		alertPub:        micro.NewEvent(AlertTopic, s.Client()),
	}
}

func (p *publisher) Publish(ctx context.Context, msg interface{}, opts ...client.PublishOption) error {
	switch msg.(type) {
	case *events.Notification:
		return p.notificationPub.Publish(ctx, msg, opts...)
	case *events.Malfunction:
		return p.malfunctionPub.Publish(ctx, msg, opts...)
	case *events.Alert:
		return p.alertPub.Publish(ctx, msg, opts...)
	default:
		return fmt.Errorf("unknown event type: %T", msg)
	}
}
