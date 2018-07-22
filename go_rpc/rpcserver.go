package main

import (
	"fmt"
	"net"
	"net/rpc"
)


type Calc int

type Args struct {
	A,B int
}

type Reply struct {
	C int
}


func(c *Calc) Sum(args Args, reply *Reply) error{

	reply.C = args.A +args.B

	return nil
}

func main() {
	rpc.Register(new(Calc)) //Calc  타입의 인스턴스를 생성하여 서버에 등록

	ln,err := net.Listen("tcp",":6000")

	if err != nil {
		fmt.Println(err)
		return 
	}

	defer ln.Close()


	for {
		conn, err := ln.Accept()

		if err != nil{
			continue
			
		}

		defer conn.Close()


		go rpc.ServeConn(conn)
	}
	
}
