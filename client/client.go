package client

import (
//	"bytes"
	"fmt"
//	"io"
	"log"
//	"io/ioutil"
	//"mime/multipart"
	"net/http"
	"os"
	"time"
	//"path"
	//"path/filepath"
	"github.com/gorilla/mux"
	"github.com/icapeg-client/config"

	ic "github.com/egirna/icap-client"
)
/*
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Create("./result.pdf")
	if err != nil {
		panic(err)
	}
	n, err := io.Copy(file, r.Body)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(fmt.Sprintf("%d bytes are recieved.\n", n)))
}*/
/*
func main() {
	http.HandleFunc("/upload", uploadHandler)

	go func() {
		time.Sleep(time.Second * 1)
		upload()
	}()

	http.ListenAndServe(":5050", nil)
}*/
/*
func upload() {
	file, err := os.Open("./example.pdf")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	res, err := http.NewRequest("POST","http://127.0.0.1:5050/upload", file)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	message, _ := ioutil.ReadAll(res.Body)
	fmt.Printf(string(message))
}
//IndexHandler is
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/pdf")
	http.ServeFile(w, r, "example.pdf")
}*/
//UploadHandler is
/*
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}*/

func createHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
//	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
w.Header().Set("Content-Type", "application/pdf")
//	w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, vars["example.pdf"])
//	name := vars["example.pdf"]
//	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
   // body, _ := ioutil.ReadAll(r.Body)
 //   fmt.Fprint(w, name)
	//ioutil.WriteFile(name, []byte(body), 0644)
	//w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	
}
func createHandlerg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, vars["example.pdf"])
//	name := vars["example.pdf"]
//	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
   // body, _ := ioutil.ReadAll(r.Body)
 //   fmt.Fprint(w, name)
	//ioutil.WriteFile(name, []byte(body), 0644)
	//w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	
}
//SetupRoutes is
func SetupRoutes() {
	/*
    http.HandleFunc("/upload", UploadFile)
	http.ListenAndServe(":8188", nil)
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", IndexHandler)
	//mux.HandleFunc("/upload", UploadHandler)

	if err := http.ListenAndServe(":4503", mux); err != nil {
		log.Fatal(err)
	}
	/*
	http.HandleFunc("/upload", uploadHandler)

	go func() {
		time.Sleep(time.Second * 1)
		upload()
	}()

	http.ListenAndServe(":5050", nil)*/
	r := mux.NewRouter()
	r.HandleFunc("/p/{example.pdf}", createHandler).Methods("POST")
	r.HandleFunc("/g/{example.pdf}", createHandler).Methods("GET")
//	r.HandleFunc("/a/{example.pdf}", createHandlerg).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", r))
}

//Clienticap is
func Clienticap() {
	//var requestHeader http.Header
	_, host, port, service, timeout, _ := config.Configtoml()
	//httpReq, err := http.NewRequest(http.MethodGet, file, nil)
/*	fileDir, _ := os.Getwd()
  fileName := "example.pdf"
  filepathhh := path.Join(fileDir, fileName)
	file, _ := os.Open(filepathhh)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()*/
	clientt := &http.Client{}
	f, err := os.Open("example.pdf")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	
	//postData := make([]byte, 1000)
	//req, err := http.NewRequest("POST", "http://example.com", bytes.NewReader(postData))
	reqq, err := http.NewRequest("POST", "http://localhost:8080/p/example.pdf",f)
	if err != nil {
		os.Exit(1)
	}
	reqq.Header.Add("Content-Type", "application/pdf; charset=utf-8")
	//reqq.Header.Add("User-Agent", "myClient")
	respp, err := clientt.Do(reqq)
	defer respp.Body.Close()
	fmt.Println(respp)


//	httpReq, _ := http.NewRequest("POST", "http://www.yahoo.com", body)
//httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:4503/upload", nil)
httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/g/example.pdf", nil)

	  //fmt.Println("Contents of file:", string(data))
		if err != nil {
		log.Fatalln(err)
		}
		httpReq.Header.Add("Content-Type", "application/pdf; charset=utf-8")
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

	//req.ExtendHeader(requestHeader)

	client := &ic.Client{
		Timeout: timeout * time.Millisecond,
	}

	resp, err := client.Do(req)

	if err != nil {
	log.Fatal(err)
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
	fmt.Println(resp.StatusCode)
	/*
		resp, err := client.Do(req)

		if err != nil {
			log.Fatal(err)
		}
	
	fmt.Println(resp.StatusCode)*/
	fmt.Println(resp.ContentResponse)
		
/*	samplefile, err := os.Create(filepathh)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer samplefile.Close()
	//x.Write(samplefile)
	io.Copy(samplefile, resp.ContentResponse.Body)*/

}