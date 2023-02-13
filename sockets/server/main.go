package main

import(
	"fmt"
	"net"
	
)
const (
	SERVER_HOST ="localhost"
	SERVER_PORT="8080"
	SERVER_TYPE="tcp"
)

func main(){
	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil{
		panic(err)
	}
	_, err = connection.Write([]byte("Hello Server! greetings."))
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil{
		fmt.Println("error")
	}
	fmt.Println("Received: ", string(buffer[:mLen]))
	defer connection.Close()

}
