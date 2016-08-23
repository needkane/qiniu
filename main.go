package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	//"encoding/base64"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"qbox.us/api/up"

	//. "github.com/qiniu/api/conf"
	//"github.com/qiniu/api/fop" //这个地方是坑，网上的api。如果下载下来，请改本地地址，这里是网络地址，而且官网没提示要引用下面几个
	qio "github.com/qiniu/api/io"
	"github.com/qiniu/api/rs"
)

var (
	ACCESS_KEY = ""
	SECRET_KEY = ""
	UP_HOST    = "http://up-z1.qiniu.com"
)

func main() {
	//	upload3()
	newUrl := makeSaveasUrl("http://needkane.qiniudn.com/lufei.jpg?imageMogr2/thumbnail/200x", ACCESS_KEY, []byte(SECRET_KEY), "needkane", "k45")
	log.Println(newUrl)
}

//GET upload access token
func uptoken(bucketName string) string {
	putPolicy := rs.PutPolicy{
		Scope: bucketName,
		//CallbackUrl: "127.0.0.1",
		//CallbackBody:callbackBody,
		//ReturnUrl: "127.0.0.1",
		//ReturnBody: returnBody,
		//AsyncOps: asyncOps,
		//EndUser: endUser,
		//Expires: expires,
		PersistentOps: "vframe/jpg/offset/1|saveas/bmVlZGthbmU6MTN2Mg==",
	}
	return putPolicy.Token(nil)
}

func upload() {
	policy := up.AuthPolicy{
		Scope:               "needkane",
		InsertOnly:          0, //如果设置为非0值, 则无论scope设置为什么形式, 仅能以新增模式上传文件: 1,
		Deadline:            time.Now().Unix() + 3600,
		CallbackUrl:         "125.64.9.132:8090/callback",
		CallbackBody:        "callbackBody",
		PersistentNotifyUrl: "125.64.9.132:8090/callback",
		//CallbackBody:callbackBody,
		//ReturnUrl: "http://101.71.89.171:6666",
		//DetectMime: detectMime,
		PersistentOps:      "vframe/jpg/offset/2",
		PersistentPipeline: "kane",
	}
	uptoken := up.MakeAuthTokenString(ACCESS_KEY, SECRET_KEY, &policy)
	var ret qio.PutRet
	var extra = &qio.PutExtra{
	//Params: params,
	//MimeType: mieType,
	//Crc32: crc32,
	//CheckCrc: CheckCrc,
	}
	err := qio.PutFile(nil, &ret, uptoken, "key123", "/home/qboxtest/Desktop/123.mp4", extra)

	if err != nil {
		//上传产生错误
		log.Print("io.PutFile failed:", err)
		return
	}

	//上传成功，处理返回值
	log.Print(ret.Hash, ret.Key)
}

func upload2() {
	file, err := os.Open("/home/qboxtest/Desktop/123.mp4")
	if err != nil {
		log.Println("os.Open failed")
	}
	defer file.Close()
	policy := up.AuthPolicy{
		Scope:        "needbc",
		InsertOnly:   0, //如果设置为非0值, 则无论scope设置为什么形式, 仅能以新增模式上传文件: 1,
		Deadline:     time.Now().Unix() + 3600,
		CallbackUrl:  "125.64.9.132:8090/callback",
		CallbackBody: "callbackBody",
		//PersistentNotifyUrl: "125.64.9.132:8090/callback",
		//CallbackBody:callbackBody,
		//DetectMime: detectMime,
		//PersistentOps:      "vframe/jpg/offset/2",
		//PersistentPipeline: "kane2",
	}
	uptoken := up.MakeAuthTokenString(ACCESS_KEY, SECRET_KEY, &policy)

	url := UP_HOST
	//bucket:filename
	entry := "needbc:132"
	//encodeEntryUri := base64.URLEncoding.EncodeToString([]byte(entry))
	//action := "/rs-put/" + encodeEntryUri
	extraParams := map[string]string{
		"token": uptoken,
		//"action": action,
		"key": "13v2",
	}
	filename := entry
	idx := strings.Index(entry, ":")
	if idx != -1 {
		filename = entry[idx+1:]
	}
	req, err := newUploadRequest(url, "file", filename, extraParams, file)
	if err != nil {
		log.Println("---------err:", err)
	}
	hc := http.Client{}
	resp, err := hc.Do(req)
	log.Println("--------create resp")
	data, _ := ioutil.ReadAll(resp.Body)
	log.Println(resp.StatusCode, "-----------", err, "---------", string(data))
}

func upload3() {
	ACCESS_KEY = ""
	SECRET_KEY = ""
	UP_HOST = "http://127.0.0.1:11200"
	file, err := os.Open("/home/qboxtest/Desktop/123.mp4")
	if err != nil {
		log.Println("os.Open failed")
	}
	defer file.Close()
	policy := up.AuthPolicy{
		Scope:               "need10",
		InsertOnly:          0, //如果设置为非0值, 则无论scope设置为什么形式, 仅能以新增模式上传文件: 1,
		Deadline:            time.Now().Unix() + 3600,
		CallbackUrl:         "125.64.9.132:8090/callback",
		CallbackBody:        "callbackBody",
		PersistentNotifyUrl: "125.64.9.132:8090/callback",
		//CallbackBody:callbackBody,
		//DetectMime: detectMime,
		PersistentOps:      "vframe/jpg/offset/2|saveas/bmVlZDEwOjEzMw==",
		PersistentPipeline: "need",
	}
	uptoken := up.MakeAuthTokenString(ACCESS_KEY, SECRET_KEY, &policy)

	url := UP_HOST
	//bucket:filename
	entry := "needbc:132"
	//encodeEntryUri := base64.URLEncoding.EncodeToString([]byte(entry))
	//action := "/rs-put/" + encodeEntryUri
	extraParams := map[string]string{
		"token": uptoken,
		//"action": action,
		"key": "13v3",
	}
	filename := entry
	idx := strings.Index(entry, ":")
	if idx != -1 {
		filename = entry[idx+1:]
	}
	req, err := newUploadRequest(url, "file", filename, extraParams, file)
	if err != nil {
		log.Println("---------err:", err)
	}
	hc := http.Client{}
	resp, err := hc.Do(req)
	log.Println("--------create resp")
	data, _ := ioutil.ReadAll(resp.Body)
	log.Println(resp.StatusCode, "-----------", err, "---------", string(data))
}

func newUploadRequest(url, paramName, fileName string, params map[string]string, file io.Reader) (req *http.Request, err error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base("/home/qboxtest/Desktop/lufei"))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	for key, val := range params {
		writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		return
	}
	req, err = http.NewRequest("POST", url, body)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return
}

func makeSaveasUrl(URL, accessKey string, secretKey []byte, saveBucket, saveKey string) string {
	accessKey = "v7VBTrPY-M28Yzz2Bq4gb0fW_yBVRQahuQQSj3B2"
	secretKey = []byte("rjJ51h_-4dwxWC2r1u_hFvg-4Vfok_ejdF4EEiuQ")
	encodedEntryURI := base64.URLEncoding.EncodeToString([]byte(saveBucket + ":" + saveKey))

	URL += "|saveas/" + encodedEntryURI

	h := hmac.New(sha1.New, secretKey)

	// 签名内容不包括Scheme，仅含如下部分：
	// <Domain>/<Path>?<Query>

	u, _ := url.Parse(URL)
	io.WriteString(h, u.Host+u.RequestURI())

	d := h.Sum(nil)
	sign := accessKey + ":" + base64.URLEncoding.EncodeToString(d)

	return URL + "/sign/" + sign

}
