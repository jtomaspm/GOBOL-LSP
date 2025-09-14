package application

import (
	"io"

	"github.com/jtomaspm/GOBOL-LSP/pkg/rpc"
)

type NotificationCenter struct {
	input  chan any
	writer io.Writer
	logger *Logger
}

func NewNotificationCenter(writer io.Writer, logger *Logger) *NotificationCenter {
	return &NotificationCenter{
		input:  make(chan any),
		writer: writer,
		logger: logger,
	}
}

func (nc *NotificationCenter) SendNotification(notification any) {
	nc.input <- notification
}

func (nc *NotificationCenter) Run() {
	go func() {
		for val := range nc.input {
			notification := []byte(rpc.EncodeMessage(val))
			if _, err := nc.writer.Write(notification); err != nil {
				nc.logger.Printf("failed to write notification: %v", err)
			}
		}
	}()
}
