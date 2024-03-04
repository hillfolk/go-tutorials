package main

import (
	"log"
)

/*
함수형 옵션 패턴 (Functional Option Pattern)
함수형 옵션 패턴은 함수가 선택적 인수를 허용하는 함수형 프로그래밍 에서 유래했다. 함수형 옵션 패턴을 사용하여 기존 함수 구조를 손상시키지 않고 확장할 수 있는 유여한 인터페이스를 제공한다.
Go에서는 구조체를 단순화 하기 위해서 사용하며 서로 다른 매개변수를 가진 많은 생성자를 정의하는 대신 다양한 함수 옵션을 허용하여 단일 생성자를 정의할 수 있다.
*/
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
