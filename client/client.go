package client

import (
	"fmt"
	"log"
	"net/http"

	"os"
	"time"
	"io"

	"github.com/icapeg-client/config"

	ic "github.com/egirna/icap-client"
)

//Clienticap is
func Clienticap() {
	var requestHeader http.Header
	file, host, port, service,timeout,filepath := config.Configtoml()
	httpReq, err := http.NewRequest(http.MethodGet, file, nil)
	
	if err != nil {
		log.Fatalln(err)
	}
	httpClient := &http.Client{}
	httpResp, err := httpClient.Do(httpReq)

	if err != nil {
		log.Fatal(err)
	}
	//var httpRequest *http.Request
	//var httpResponse *http.Response

	icap := "icap://" + host + ":" + port + "/" + service
	req, err := ic.NewRequest(ic.MethodRESPMOD, icap, httpReq, httpResp)

	if err != nil {
		panic(err)
	}

	req.ExtendHeader(requestHeader)

	client := &ic.Client{
		Timeout: timeout * time.Millisecond,
	}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	//	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//We Read the response body on the line below.
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//Convert the body to type string
	//sb := string(body)
	//log.Printf(sb)
	// }
	/*
		var w http.Response
		var r *http.Request

		httpReq, err := http.NewRequest(http.MethodGet, "http://www.eicar.org/download/eicar_com.zip", nil)
		//resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
		if err != nil {
			log.Fatalln(err)
		}
		//body, err := ioutil.ReadAll(httpReq.Body)
		//if err != nil {
		//	log.Fatalln(err)
		//}
		//sb := string(body)
		//log.Printf(sb)
		//httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
		//client :=   &http.Client{}
		//dat, err := ioutil.ReadFile("/config.json")
		//var dat = []byte(`{"key1":"...Z-DAaFpGT0t...","key2":"..."}`)
		//httpReq, err := http.NewRequest("PUT", "http://www.yahoo.com", nil)
		//resp, err := client.Do(req)
		//fmt.Println(file)
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		httpClient := &http.Client{}
		httpResp, err := httpClient.Do(httpReq)

		if err != nil {
			log.Fatal(err)
		}
		//	body, err := ioutil.ReadAll(httpResp.Body)
		//	if err != nil {
		//		log.Fatalln(err)
		//	}
		//	sb := string(body)
		//	log.Printf(sb)
		w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(filename))
		w.Header().Set("Content-Type", "application/octet-stream")
		http.ServeFile(w, r, filePath)

		icap := "icap://" + host + ":" + port + "/" + service
		//Making a simple RESPMOD call
		req, err := ic.NewRequest(ic.MethodRESPMOD, icap, httpReq, httpResp)
		fmt.Println(icap)

		if err != nil {
			log.Fatal(err)
		}
		//requestHeader   http.Header
		//req.ExtendHeader(http.Header)

		client := &ic.Client{
			Timeout: 5 * time.Second,
		}

		resp, err := client.Do(req)

		if err != nil {
			log.Fatal(err)
		}
		//Setting preview obtained from OPTIONS call
		*/
			optReq, err := ic.NewRequest(ic.MethodOPTIONS, icap, nil, nil)

			if err != nil {
				log.Fatal(err)
				return
			}
		//	client := &ic.Client{
		//		Timeout: 5 * time.Second,
		//	}
			optResp, err := client.Do(optReq)

			if err != nil {
				log.Fatal(err)
				return
			}
			//Making simple response call
		/*	req, err := ic.NewRequest(ic.MethodRESPMOD, icap, httpReq, httpResp)
			fmt.Println(icap)

			if err != nil {
				log.Fatal(err)
			}
*/
			req.SetPreview(optResp.PreviewBytes)
/*
			resp, err := client.Do(req)

			if err != nil {
				log.Fatal(err)
			}
*/
fmt.Println(resp.StatusCode)
//fmt.Println(resp.ContentResponse)

samplefile, err := os.Create(filepath)
if err != nil {
 fmt.Println(err)
 os.Exit(1)
}
 defer samplefile.Close()
//x.Write(samplefile)
io.Copy(samplefile, resp.ContentResponse.Body)	

}
