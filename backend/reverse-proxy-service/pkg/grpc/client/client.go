package client

type Client interface {
	Connect() (interface{}, error)
	Close()
}
