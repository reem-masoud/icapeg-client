package main

import (
	"fmt"
	"log"
	"net/http"

	//"os"
	"time"

	ic "github.com/egirna/icap-client"
)

func main() {

	httpReq, err := http.NewRequest(http.MethodGet, "http://www.eicar.org/download/eicar_com.zip", nil)

	if err != nil {
		log.Fatal(err)
	}
	httpClient := &http.Client{}
	httpResp, err := httpClient.Do(httpReq)

	if err != nil {
		log.Fatal(err)
	}
	//Making a simple RESPMOD call
	req, err := ic.NewRequest(ic.MethodRESPMOD, "icap://127.0.0.1:1344/respmod", httpReq, httpResp)

	if err != nil {
		log.Fatal(err)
	}

	client := &ic.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	//Setting preview obtained from OPTIONS call
	optReq, err := ic.NewRequest(ic.MethodOPTIONS, "icap://127.0.0.1:1344/respmod", nil, nil)

	if err != nil {
		log.Fatal(err)
		return
	}

	optResp, err := client.Do(optReq)

	if err != nil {
		log.Fatal(err)
		return
	}

	req.SetPreview(optResp.PreviewBytes)
	fmt.Println(resp.StatusCode)

}
