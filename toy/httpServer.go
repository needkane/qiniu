package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"time"
)

type IdcInfo struct {
	T time.Time
	Idc   string   `json:"idc"`
	Nodes []string `json:"nodes"`
}
type AllIdc struct {
	Idcs []IdcInfo `json:"idcs"`
}

func main() {
	net.Listen("tcp", ":81")
	//define handler
	http.HandleFunc("/listnodes", serveUfop)

	//bind and listen

	ufopServer := &http.Server{
		Addr: "0.0.0.0:8881",
	}

	listenErr := ufopServer.ListenAndServe()
	if listenErr != nil {
		log.Println(listenErr)
	}
}
func get(f io.Writer, client *http.Client)(n int64,err error){
	
	r, err := http.NewRequest("GET", "http://127.0.0.1:3808/13", nil) 
 if err != nil {
	         return
			     
 }    
	     res, err := client.Do(r)
		 if err != nil {
			         return
					     
		 }    
		     defer res.Body.Close()

n,err = io.Copy(f,res.Body)
return

}
func xx() (an int64){
	    f, err := os.OpenFile("allts", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0600)
		    defer f.Close()
			    log.Println(f,"------",err)
				client:= &http.Client{}
				for i:=0;i<40;i++{
					 
//					    f1,err1:=os.Open("/home/qboxtest/Desktop/13")
//						    log.Println(f1,"===",err1)
//							    defer f1.Close()
			n,err:=get(f,client)
			if err!=nil{
				log.Println("get failed: ",err)
				return
			}	
			an =an+n
	}
										      return  
}
func serveUfop(w http.ResponseWriter, req *http.Request) {
	log.Println("in-----------")
	i1 := IdcInfo{T: time.Now(),Idc: "1h8AAKGcouQBYzcU:ctc-lz-1-1-c24c1", Nodes: []string{"xA8AABFvqdLGxi8U:EACNBGPBJJBJJ01C002", "1h8AAKGcouQBYzcU:ctc-lz-1-1-c24c1"}}
	i2 := IdcInfo{T:time.Now(),Idc: "x1cAAGiWbL176zAU:EACNCMCFJMFZM02C001", Nodes: []string{"x1cAAGiWbL176zAU:EACNCMCFJMFZM02C001"}}
	all := AllIdc{Idcs: []IdcInfo{i1, i2}}
	bs, _ := json.Marshal(all)
//	time.Sleep(50 * time.Second)
an:=xx()
log.Println(an)
WriteJsonError(w, 200, bs)
}

func WriteJsonError(w http.ResponseWriter, code int, content []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	w.Write(content)
}

func ReplyImageStream(w http.ResponseWriter) {
	f, _ := os.Open("/home/qbox5ddtest/qbox/devtools/bin/a.m3u8")
	defer f.Close()
	fi, _ := f.Stat()

	fsize := fi.Size()

	w.Header().Set("Content-Type", "")
	w.Header().Set("Content-Length", strconv.FormatInt(fsize, 10))
	w.WriteHeader(200)
	io.CopyN(w, f, fsize)
}
