package main

import (
	"bytes"
	//"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
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
	ACCESS_KEY = "fDidaMyecyd7b9Vtc1Yqzl3unHcfAli4EC4iSisR"
	SECRET_KEY = "tFYXaoOXdVtdmnujRMHgFGfmNdmUgJNdzvV-fdXy"
	UP_HOST    = "http://up-z1.qiniu.com"
)

func main() {
	upload2()
}

//GET upload access token
func uptoken(bucketName string) string {
	putPolicy := rs.PutPolicy{
		Scope: bucketName,
		//CallbackUrl: callbackUrl,
		//CallbackBody:callbackBody,
		//ReturnUrl: returnUrl,
		//ReturnBody: returnBody,
		//AsyncOps: asyncOps,
		//EndUser: endUser,
		//Expires: expires,
		PersistentOps: "vframe/jpg/offset/1|saveas/bmVlZGthbmU6MTN2Mg==",
	}
	return putPolicy.Token(nil)
}

func upload() {
	uptoken := uptoken("needkane")
	fmt.Printf("uptoken:%s\n", uptoken)

	var err error
	var ret qio.PutRet
	var extra = &qio.PutExtra{
	//Params: params,
	//MimeType: mieType,
	//Crc32: crc32,
	//CheckCrc: CheckCrc,
	}

	var key = "face-36123"
	var localFile = "/home/qboxtest/Desktop/13.mkv"

	// ret 变量用于存取返回的信息，详情见 io.PutRet
	// uptoken 为业务服务器生成的上传口令
	// key 为文件存储的标识
	// localFile 为本地文件名
	// extra 为上传文件的额外信息，详情见 io.PutExtra，可选
	err = qio.PutFile(nil, &ret, uptoken, key, localFile, extra)

	if err != nil {
		//上传产生错误
		log.Print("io.PutFile failed:", err)
		return
	}

	//上传成功，处理返回值
	log.Print(ret.Hash, ret.Key)

}

func upload2() {
	file, err := os.Open("/home/qboxtest/Desktop/13.mkv")
	if err != nil {
		log.Println("os.Open failed")
	}
	defer file.Close()
	policy := up.AuthPolicy{
		Scope:      "needbc",
		InsertOnly: 0, //如果设置为非0值, 则无论scope设置为什么形式, 仅能以新增模式上传文件: 1,
		Deadline:   uint32(time.Now().Unix()) + 3600,
		//DetectMime: detectMime,
		PersistentOps:      "vframe/jpg/offset/1|saveas/bmVlZGJjOjEzdjI=",
		PersistentPipeline: "kane",
	}
	uptoken := up.MakeAuthTokenString(ACCESS_KEY, SECRET_KEY, &policy)
	url := UP_HOST
	//bucket:filename
	entry := "needbc:13v1"
	//encodeEntryUri := base64.URLEncoding.EncodeToString([]byte(entry))
	//action := "/rs-put/" + encodeEntryUri
	extraParams := map[string]string{
		"token": uptoken,
		//"action": action,
		"key": "13v",
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
