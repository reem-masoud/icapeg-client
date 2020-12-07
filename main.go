package main

import (
"github.com/icapeg-client/client"
	
)

func main() {
 //  go client.SetupRoutes()
	 //client.Clienticap()

	 go func(){
	 client.SetupRoutes()
	 }()
	 client.Clienticap()


}
