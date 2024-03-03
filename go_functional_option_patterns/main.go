package main

import (
	"log"
)

type ClientOptions struct {
	Url    string
	Port   int
	Method string
}

type Option func(*ClientOptions) error

func WithUrl(url string) Option {
	return func(co *ClientOptions) error {
		co.Url = url
		return nil
	}
}

func WithPort(port int) Option {
	return func(co *ClientOptions) error {
		co.Port = port
		return nil
	}
}

func WithMethod(method string) Option {
	return func(co *ClientOptions) error {
		co.Method = method
		return nil
	}
}

func NewClient(opts ...Option) (*ClientOptions, error) {
	var co ClientOptions
	for _, opt := range opts {
		err := opt(&co)
		if err != nil {
			return nil, err
		}
	}
	return &co, nil
}

func main() {
	client, err := NewClient(WithUrl("http://localhost"), WithPort(8080), WithMethod("GET"))
	if err != nil {
		panic(err)
	}

	log.Println(client.Port)
	log.Println(client.Url)
	log.Println(client)

}
