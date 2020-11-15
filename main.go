package main

import (
	"icapeg/client"
	"icapeg/server"
)

func main() {

	go server.StartServer()
	 client.IcapClient()
}
