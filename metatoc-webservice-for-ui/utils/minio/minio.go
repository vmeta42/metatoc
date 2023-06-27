package minio

import (
	"context"
	"fmt"
	"os"
	"strings"

	"gitlab.dev.21vianet.com/liu.hao8/metatoc-service/utils/common"

	"gitlab.dev.21vianet.com/liu.hao8/metatoc-service/utils/file"

	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/beego/beego/v2/core/logs"

	"github.com/minio/minio-go/v7"
)

const (
	DEFAULT_MINIO_ENDPOINT      = "minio.qa.meta42.indc.vnet.com"
	DEFAULT_MINIO_ACCESS_KEY    = "minioadmin"
	DEFAULT_MINIO_SECRET_KEY    = "minioadmin"
	DEFAULT_MINIO_USE_SSL_TRUE  = true
	DEFAULT_MINIO_USE_SSL_FALSE = false
	DEFAULT_MINIO_BUCKET_NAME   = "test"
)

// InitVarEndpoint 初始化环境变量MINIO_ENDPOINT
func InitVarEndpoint() string {
	if os.Getenv("MINIO_ENDPOINT") != "" {
		return os.Getenv("MINIO_ENDPOINT")
	} else {
		return DEFAULT_MINIO_ENDPOINT
	}
}

// InitVarAccessKey 初始化环境变量MINIO_ACCESS_KEY
func InitVarAccessKey() string {
	if os.Getenv("MINIO_ACCESS_KEY") != "" {
		return os.Getenv("MINIO_ACCESS_KEY")
	} else {
		return DEFAULT_MINIO_ACCESS_KEY
	}
}

// InitVarSecretKey 初始化环境变量MINIO_SECRET_KEY
func InitVarSecretKey() string {
	if os.Getenv("MINIO_SECRET_KEY") != "" {
		return os.Getenv("MINIO_SECRET_KEY")
	} else {
		return DEFAULT_MINIO_SECRET_KEY
	}
}

// InitVarUseSSL 初始化环境变量MINIO_USE_SSL
func InitVarUseSSL() bool {
	if os.Getenv("MINIO_USE_SSL") == "true" {
		return DEFAULT_MINIO_USE_SSL_TRUE
	} else {
		return DEFAULT_MINIO_USE_SSL_FALSE
	}
}

// InitVarBucketName 初始化环境变量MINIO_BUCKET_NAME
func InitVarBucketName() string {
	if os.Getenv("MINIO_BUCKET_NAME") != "" {
		return os.Getenv("MINIO_BUCKET_NAME")
	} else {
		return DEFAULT_MINIO_BUCKET_NAME
	}
}

var (
	client     *minio.Client
	endpoint   = InitVarEndpoint()
	accessKey  = InitVarAccessKey()
	secretKey  = InitVarSecretKey()
	useSSL     = InitVarUseSSL()
	bucketName = InitVarBucketName()
)

// CreateBucket 创建存储桶
func CreateBucket(bucketName string) error {
	// 创建MinIO客户端
	client, err := newClient()
	if err != nil {
		return err
	}

	// 检查Bucket是否已存在
	exists, err := client.BucketExists(context.Background(), bucketName)
	if err != nil {
		return err
	}

	if !exists {
		// 创建Bucket
		err = client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}

// Upload 上传文件
func Upload(folderName string, fileName string) error {
	// 创建MinIO客户端
	client, err := newClient()
	if err != nil {
		return err
	}

	// 创建文件
	fileFullName, err := file.CreateUploadFile(folderName, fileName)
	if err != nil {
		return err
	}

	file, err := os.Open(fileFullName)
	if err != nil {
		return err
	}
	defer file.Close()

	// 获取文件信息
	fileStat, err := file.Stat()
	if err != nil {
		return err
	}
	fileSize := fileStat.Size()

	// 设置上传选项
	opts := minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	}

	// 执行上传操作
	_, err = client.PutObject(context.TODO(), bucketName, folderName+"/"+fileName, file, fileSize, opts)
	if err != nil {
		return err
	}

	return nil
}

// Download 下载文件
func Download(objectName string) (string, error) {
	// 创建MinIO客户端
	client, err := newClient()
	if err != nil {
		return "", err
	}

	objectNameSlice := strings.Split(objectName, "/")
	folderName := objectNameSlice[0]
	fileName := objectNameSlice[1]
	fileNewName := strings.Replace(fileName, ".tmp", fmt.Sprintf(".%s.tmp", common.GenerateRandomString(8)), -1)

	err = file.CreateDownloadFolder(folderName)
	if err != nil {
		return "", err
	}

	fileFullName, err := file.CreateDownloadFile(folderName, fileNewName)
	if err != nil {
		return "", err
	}

	// 下载文件
	err = client.FGetObject(context.TODO(), bucketName, objectName, fileFullName, minio.GetObjectOptions{})
	if err != nil {
		return "", err
	}

	return fileFullName, nil
}

// ListFolders 列出文件夹
func ListFolders(address string) ([]minio.ObjectInfo, error) {
	return listObjects(address, false)
}

// ListFiles 列出文件
func ListFiles(folderName string) ([]minio.ObjectInfo, error) {
	return listObjects(folderName, false)
}

// NewMinioClient 创建MinIO客户端
func newClient() (*minio.Client, error) {
	if client != nil {
		return client, nil
	} else {
		logs.Info("endpoint: %s, accessKey: %s, secretKey: %s, useSSL: %v, bucketName: %s\n", endpoint, accessKey, secretKey, useSSL, bucketName)
		client, err := minio.New(endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
			Secure: useSSL,
		})
		if err != nil {
			return nil, err
		}
		return client, nil
	}
}

// 列出存储桶中的对象
func listObjects(prefix string, recursive bool) ([]minio.ObjectInfo, error) {
	// 判断prefix最后一个字符是否为'/'，如果不是则添加
	if prefix[len(prefix)-1:] != "/" {
		prefix = prefix + "/"
	}

	// 创建MinIO客户端
	client, err := newClient()
	if err != nil {
		return nil, err
	}

	// 设置列出文件的选项
	opts := minio.ListObjectsOptions{
		Prefix:    prefix,    // 文件夹路径
		Recursive: recursive, // 递归列出文件夹下的文件
	}

	// 执行列出文件操作
	objectInfoSlice := make([]minio.ObjectInfo, 0)
	for object := range client.ListObjects(context.TODO(), bucketName, opts) {
		if object.Err != nil {
			return nil, object.Err
		}
		objectInfoSlice = append(objectInfoSlice, object)
	}

	logs.Info("objectInfoSlice: %v\n", objectInfoSlice)

	return objectInfoSlice, nil
}
