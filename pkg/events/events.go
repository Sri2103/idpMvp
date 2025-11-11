// pkg/events/events.go
package events

import (
	"log"

	"github.com/nats-io/nats.go"
)

type Bus struct {
	conn *nats.Conn
}

func New(url string) *Bus {
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	return &Bus{conn: nc}
}

func (b *Bus) Publish(subject string, msg []byte) {
	_ = b.conn.Publish(subject, msg)
}
