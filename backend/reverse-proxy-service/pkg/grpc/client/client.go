package client

import "context"

type Client interface {
	// Connect подключается к внешнему сервису по grpc
	Connect() (interface{}, error)
	// Close закрывает соединение с внешним сервисом
	Close()
	// Ping запускает отдельную горутину и отслеживает состояние подключения
	Ping(ctx context.Context)
}
