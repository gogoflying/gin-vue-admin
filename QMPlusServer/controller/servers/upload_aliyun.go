package servers

import (
	"fmt"
	"gin-vue-admin/config"
	"io/ioutil"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var endPoint string = config.GinVueAdminconfig.AliyunOss.EndPoint
var accessKeyId string = config.GinVueAdminconfig.AliyunOss.AccessKeyId
var accessKeySecret string = config.GinVueAdminconfig.AliyunOss.AccessKeySecret

// 接收两个参数 一个文件流 一个 bucket 你的七牛云标准空间的名字
func UploadFileOss(file *multipart.FileHeader, bucketName string, urlPath string) (err error, path string, key string) {
	client, err := oss.New(endPoint, accessKeyId, accessKeySecret)
	if err != nil {
		return
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	fmt.Println("bucket over")
	if err != nil {
		return
	}

	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)
	f, err := file.Open()
	fmt.Println("open over")
	if err != nil {
		return
	}
	//var respHeader http.Header
	err = bucket.PutObject(fileKey, f)
	fmt.Println("putobject over")
	if err != nil {
		return
	}

	singUrl := "https://" + bucketName + "." + endPoint + "/" + fileKey
	return err, singUrl, fileKey
}

func DownloadFileOss(bucketName string, key string) (data []byte, err error) {
	client, err := oss.New(endPoint, accessKeyId, accessKeySecret)
	if err != nil {
		return
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return
	}

	singUrl := "https://" + bucketName + "." + endPoint + "/" + key
	body, err := bucket.GetObjectWithURL(singUrl)
	if err != nil {
		return
	}
	defer body.Close()

	data, err = ioutil.ReadAll(body)
	if err != nil {
		return
	}
	return
}

func DeleteFileOss(bucketName string, key string) (err error) {
	client, err := oss.New(endPoint, accessKeyId, accessKeySecret)
	if err != nil {
		return
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return
	}
	return bucket.DeleteObject(key)
}

func UploadLocalFile(fileName string, bucketName string, urlPath string) (err error, path string, key string) {
	client, err := oss.New(endPoint, accessKeyId, accessKeySecret)
	if err != nil {
		return
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return
	}

	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), fileName)
	//var respHeader http.Header
	err = bucket.PutObjectFromFile(fileKey, fileName)
	if err != nil {
		return
	}

	singUrl := "https://" + bucketName + "." + endPoint + "/" + fileKey
	return err, singUrl, fileKey
}

func DownLoadLocalFile(bucketName string, urlPath string, localPath string) (err error, path string) {
	client, err := oss.New(endPoint, accessKeyId, accessKeySecret)
	if err != nil {
		return err, ""
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err, ""
	}

	fileName := filepath.Base(urlPath)
	localPath = localPath + "/" + fileName
	//var respHeader http.Header
	err = bucket.GetObjectToFile(getDownloadFileName(urlPath), localPath)
	if err != nil {
		return err, ""
	}
	return err, localPath
}

func getDownloadFileName(source string) string {
	//source := "http://bda-edu-hr.oss-cn-beijing.aliyuncs.com/1588560066tmp//zzz_6.jpg"
	npos := strings.LastIndex(source, "aliyuncs.com/")
	return source[npos+len("aliyuncs.com/") : len(source)]
}
