package client

import (
	"fmt"
	"log"
	"net/http"

	//"os"
	"time"

	"github.com/icapeg-client/config"

	ic "github.com/egirna/icap-client"
)

//Clienticap is
func Clienticap() {
	file, host, port, service := config.Configtoml()
	httpReq, err := http.NewRequest(http.MethodGet, file, nil)
	fmt.Println(file)
	if err != nil {
		log.Fatal(err)
	}
	httpClient := &http.Client{}
	httpResp, err := httpClient.Do(httpReq)

	if err != nil {
		log.Fatal(err)
	}
	icap := "icap://" + host + ":" + port + "/" + service
	//Making a simple RESPMOD call
	/*req, err := ic.NewRequest(ic.MethodRESPMOD, icap, httpReq, httpResp)
	fmt.Println(icap)

	if err != nil {
		log.Fatal(err)
	}

	client := &ic.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}*/
	//Setting preview obtained from OPTIONS call
	optReq, err := ic.NewRequest(ic.MethodOPTIONS, icap, nil, nil)

	if err != nil {
		log.Fatal(err)
		return
	}
	client := &ic.Client{
		Timeout: 5 * time.Second,
	}
	optResp, err := client.Do(optReq)

	if err != nil {
		log.Fatal(err)
		return
	}
	//Making simple response call
	req, err := ic.NewRequest(ic.MethodRESPMOD, icap, httpReq, httpResp)
	fmt.Println(icap)

	if err != nil {
		log.Fatal(err)
	}

	req.SetPreview(optResp.PreviewBytes)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.StatusCode)

}
