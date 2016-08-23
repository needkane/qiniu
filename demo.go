package main

import (
	. "github.com/qiniu/api/conf"
	"github.com/qiniu/api/fop"
	//"github.com/qiniu/api/io"
	"github.com/qiniu/api/rs"
	"github.com/qiniu/api/rsf"
	"github.com/qiniu/rpc.v1"
	"io"
	"log"
)

func init() {

	ACCESS_KEY = "-9qVmPlRdq3CJgiyRBjnKX4ihYz4AHpFjYVkiR2A"
	SECRET_KEY = "bdiREwcdRTa8ORqjim2HGPpp3EA7nNxN02nDf4ur"
}

func main() {
	delFile2()
	/*imageAttr2()
	copyFile("needkane", "need", "needcrystal", "needtoo")*/
	//log.Println(downloadUrl("7te1kx.com1.z0.glb.clouddn.com", "HMRZnzCmSjWx6L3JhegtAg78Ais=/lmEUf_B9rb2bCwQLFc-K7EO6O8yw"))
	//getBatchInfo()
	/*highBatch()
	println(makeViewUrl("http://needkane.qiniudn.com/need"))
	var rc = rsf.New(nil)

	listAll(nil, &rc, "need10", "13")*/

	//log.Println(downloadUrl("needkane", "13.mkv"))
}

//5.2delete file
func delFile2() {
	bucket := "need10"
	key := "132"
	var rsCli = rs.New(nil)
	err := rsCli.Delete(nil, bucket, key)
	if err != nil {
		//generate error
		log.Println("rs.Copy failed:", err)
		return
	} else {
		log.Println("success delete:", key)
	}

}

//6.1.1 query image attribute
func imageAttr2() {
	var imageUrl = "http://needkane.qiniudn.com/kane2.jpg"
	ii := fop.ImageInfo{}
	infoRet, err := ii.Call(nil, imageUrl)
	if err != nil {
		log.Println("fop getImageInfo failed:", err)
		return
	}

	log.Println(infoRet.Height, infoRet.Width, infoRet.ColorModel, infoRet.Format)
}

//6.3 copy file
func copyFile(bucketSrc string, keySrc string, bucketDest string, keyDest string) {
	var rsCli = rs.New(nil)
	err := rsCli.Copy(nil, bucketSrc, keySrc, bucketDest, keyDest)
	if err != nil {
		log.Println("rs.Copy failed :", err)
		return
	} else {
		log.Println("copy success to:", bucketDest)
	}
}

//download
func downloadUrl(domain, key string) string {
	baseUrl := rs.MakeBaseUrl(domain, key)
	//baseUrl = baseUrl + "?pm3u8/0/e/1438162783"
	baseUrl = baseUrl + "?avinfo"
	policy := rs.GetPolicy{}
	return policy.MakeRequest(baseUrl, nil)
}
func getBatchInfo() {
	entryPathes := []rs.EntryPath{
		rs.EntryPath{
			Bucket: "needkane",
			Key:    "kane3.jpg",
		},
		rs.EntryPath{
			Bucket: "needcrystal",
			Key:    "needtoo",
		},
	}
	var batchStatRests []rs.BatchStatItemRet
	var rsCli = rs.New(nil)

	batchStatRests, err := rsCli.BatchStat(nil, entryPathes)
	if err != nil {
		log.Println("rs.BatchStat failed", err)
		return
	}
	for time, item := range batchStatRests {
		log.Println(time, ";;;;", item)
	}

}

//高级批处理功能
func highBatch() {
	ops := []string{
		rs.URIStat("needkane", "need"),
		rs.URICopy("needkane", "kane3.jpg", "needcrystal", "kane"),
		rs.URIDelete("needkane", "kane3.jpg"),
	}
	rets := new([]rs.BatchItemRet)
	var rsCli = rs.New(nil)
	err := rsCli.Batch(nil, rets, ops)
	if err != nil {
		log.Println("rs.Batch failed:", err)
		return
	}
	for time, ret := range *rets {
		log.Println(time, "::", ret.Code, ret.Error)
	}
}

//preview
func makeViewUrl(imageUrl string) string {
	/*	var view = fop.ImageView{
		//Mode int      // 缩略模式
		Width:   1000,  // Width = 0 表示不限定宽度
		Height:  1200,  // Height = 0 表示不限定高度
		Quality: 100,   // 质量, 1-100
		Format:  "gif", // 输出格式，如jpg, gif, png, tif等等
	}*/
	var view = fop.ImageView{}
	view.Height = 100
	view.Width = 500
	view.Quality = 100
	return view.MakeRequest(imageUrl)
}
func listAll(l rpc.Logger, rc *rsf.Client, bucketName string, prefix string) {
	var entries []rsf.ListItem
	var marker = ""
	var err error
	var limit = 1000
	for err == nil {

		entries, marker, err = rc.ListPrefix(nil, bucketName, prefix, marker, limit)
		for _, item := range entries {
			log.Print("item:", item)
		}
	}

	if err != io.EOF {
		//非预期的错误
		log.Print("listAll failed:", err)
	}
}
