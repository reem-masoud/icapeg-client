package main

import (
"github.com/icapeg-client/client"
	
)

func main() {

	go client.Clienticap()
	client.SetupRoutes()
//	go client.Clienticap()
	

}
