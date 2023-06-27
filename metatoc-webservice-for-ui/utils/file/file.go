package file

import (
	"fmt"
	"os"
	"time"

	"gitlab.dev.21vianet.com/liu.hao8/metatoc-service/utils/common"
)

const (
	UPLOAD_FOLDER_BASE_NAME   = "./attachments/upload"
	DOWNLOAD_FOLDER_BASE_NAME = "./attachments/download"
)

// CreateUploadFolder 创建上传文件夹
func CreateUploadFolder(address string) (string, error) {
	folderName := generateFolderName(address)
	if err := createFolder(UPLOAD_FOLDER_BASE_NAME + "/" + folderName); err != nil {
		return "", err
	} else {
		return folderName, nil
	}
}

// CreateDownloadFolder 创建下载文件夹
func CreateDownloadFolder(folderName string) error {
	return createFolder(DOWNLOAD_FOLDER_BASE_NAME + "/" + folderName)
}

// CreateUploadFile 创建上传文件
func CreateUploadFile(folderName string, fileName string) (string, error) {
	fileFullName := UPLOAD_FOLDER_BASE_NAME + "/" + folderName + "/" + fileName

	err := createFile(fileFullName, "Hello, MinIO!")
	if err != nil {
		return "", err
	}

	return fileFullName, nil
}

// CreateDownloadFile 创建下载文件
func CreateDownloadFile(folderName string, fileName string) (string, error) {
	fileFullName := DOWNLOAD_FOLDER_BASE_NAME + "/" + folderName + "/" + fileName

	err := createFile(fileFullName, "")
	if err != nil {
		return "", err
	}

	return fileFullName, nil
}

// 生成文件夹名
func generateFolderName(address string) string {
	return fmt.Sprintf("%s/%s.%d", address, common.GenerateRandomString(8), time.Now().Unix())
}

// 创建文件夹
func createFolder(folderName string) error {
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		// 文件夹不存在，创建文件夹
		if err := os.MkdirAll(folderName, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

// 创建文件
func createFile(fileFullName string, content string) error {
	file, err := os.Create(fileFullName)
	if err != nil {
		return err
	}
	defer file.Close()

	if content != "" {
		_, err = file.WriteString(content)
		if err != nil {
			return err
		}
	}

	return nil
}
