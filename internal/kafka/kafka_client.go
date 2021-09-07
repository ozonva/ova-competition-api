package kafka

import (
	"context"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
)

type Event interface {
	toMessage() kafka.Message
}

type Client interface {
	Send(ctx context.Context, messages ...Event) error
	Close() error
}

type client struct {
	writer *kafka.Writer
}

func NewClient(address, topic string) Client {
	return &client{
		writer: &kafka.Writer{
			Addr:  kafka.TCP(address),
			Topic: topic,
		},
	}
}

func (c *client) Send(ctx context.Context, messages ...Event) error {
	err := c.writer.WriteMessages(ctx, eventsToMessages(messages)...)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to send messages to kafka: %v", err))
	}

	return nil
}

func (c *client) Close() error {
	return c.writer.Close()
}

func eventsToMessages(events []Event) []kafka.Message {
	out := make([]kafka.Message, 0, len(events))
	for _, event := range events {
		out = append(out, event.toMessage())
	}

	return out
}
