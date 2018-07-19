package main


import (
	"fmt"
	"net"
)

func requestHandler(c net.Conn){
	data := make([]byte,4096) //슬라이스 생성

	for { // 무한 루프
		n,err := c.Read(data)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(data[:n])) //byte 데이터 문자열 출력

		_,err = c.Wirte(data[:n])

		if err != nil {
			fmt.Println(err)

			return
		}
	}

	
}

func main(){
	ln,err := net.Listen("tcp",":8000") // TCP Protocol 8000 open

	if err != nil {
		fmt.Println(err)
		return 
	}

	defer ln.Close() //Main 함수 종료 전에 연결을 종료 

	for{

		conn,err := ln.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		defer conn.Close()

		go requestHandler(conn) //패킷 처리 함수 실행 

		
	}
	
}
