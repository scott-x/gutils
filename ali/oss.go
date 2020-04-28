package ali

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/scott-x/gutils/model"
	"os"
	"strings"
)

func UploadImageToAliYunOSS(base64String string, remote_destination string, m *model.OSS) string {
	if base64String[:10] == "data:image" {
		//the image comes from front end: input.src
		index := strings.Index(base64String, ",") //data:image/jpeg;base64,/9j/4AAQSkZJ

	}
	//解压
	dist, _ := base64.StdEncoding.DecodeString(base64String)
	// fmt.Println(string(dist))
	//写入新文件
	f, _ := os.OpenFile(remote_destination, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	f.Write(dist)
	// data:image/png;base64,
	// http: //image-layer.oss-cn-shenzhen.aliyuncs.com/blog/imgs/xxx.jpeg
	// 创建OSSClient实例。
	client, err := oss.New(m.Endpoint, m.AccessKeyId, m.AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
		return ""
	}

	// 获取存储空间。
	bucket, err := client.Bucket(m.Bucket)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
		return ""
	}

	// 上传Byte数组。
	err = bucket.PutObject(remote_destination, bytes.NewReader(dist))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
		return ""
	}
	return "https://" + m.Bucket + "." + m.Endpoint + "/" + remote_destination
}
