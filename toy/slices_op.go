package main

import (
	"encoding/json"
	"fmt"
	//"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/qiniu/rpc.v3"
)

type Http_API_Op struct {
	strBase string
	aParams map[string]string
	uMode   uint
}

type device_info struct {
	algo           float64
	area           string
	channel_num    float64
	channel_online string
	connect_status string
	dev_type       float64
	device_status  float64
	did            string
	expire         float64
	from           string
	max_rate       float64
	min_rate       float64
	paddr          string
	rtsp_url       string
	stream_serv    string
	tcp_flag       float64
	version        string
	video_height   float64
	video_type     float64
	video_width    float64
}

type device_status struct {
	devices_status []device_info
}

func main() {

	var http_trans_cfg rpc.TransportConfig
	var resp_ins *http.Response
	var err error

	http_trans_cfg.DialTimeout = 50000000
	http_trans_cfg.ResponseHeaderTimeout = 50000000
	http_trans_cfg.MaxIdleConnsPerHost = 2
	var http_rip_ins http.RoundTripper
	p := fmt.Println

	http_rip_ins = rpc.NewTransport(&http_trans_cfg)

	p(http_rip_ins)

	http_client_ins := rpc.NewClientWithTransport(http_rip_ins)
	fmt.Println(http_client_ins)

	resp_ins, err = http_client_ins.DoRequest(nil, "GET", "http://115.159.192.233:9010/query_device_status?device_id=GF151228ULUCU00062AB")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(resp_ins)

	fmt.Println(http_client_ins)

	var http_get_device Http_API_Op
	http_get_device.aParams = make(map[string]string)

	http_get_device.strBase = "http://115.159.192.233:9010/query_device_status"

	var all_device_infos device_status
	var all_device_any interface{}
	fi, err := os.Open("query_device_status.json")
	if err != nil {
		panic(err)
	}

	defer fi.Close()

	fd, err := ioutil.ReadAll(fi)
	fmt.Println(string(fd))
	err = json.Unmarshal(fd, &all_device_infos)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(fd, &all_device_any)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(all_device_any)
	fmt.Println(all_device_infos)
}
