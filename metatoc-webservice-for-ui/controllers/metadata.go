package controllers

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"gitlab.dev.21vianet.com/liu.hao8/metatoc-service/utils/blockchain"

	"gitlab.dev.21vianet.com/liu.hao8/metatoc-service/utils/file"

	"gitlab.dev.21vianet.com/liu.hao8/metatoc-service/utils/minio"

	"github.com/beego/beego/v2/core/logs"

	"gitlab.dev.21vianet.com/liu.hao8/metatoc-service/errors"

	modelBase "gitlab.dev.21vianet.com/liu.hao8/metatoc-service/models/base"

	modelMetadata "gitlab.dev.21vianet.com/liu.hao8/metatoc-service/models/metadata"

	beego "github.com/beego/beego/v2/server/web"
)

type MetadataController struct {
	beego.Controller
}

// @router /upload [post]
func (c *MetadataController) Upload() {
	var resp modelMetadata.UploadResp

	// 获取请求体的原始字节数据
	requestBody := c.Ctx.Input.RequestBody

	// 解析请求体为RequestBody结构体
	var uploadReq modelMetadata.UploadReq
	err := json.Unmarshal(requestBody, &uploadReq)
	if err != nil {
		logs.Error("json.Unmarshal error: %v", err)
		resp.Error = &modelBase.Error{
			Message: errors.UPLOAD_FILE_PARAMS_READ_ERROR_MESSAGE,
		}
	} else {
		folderName, err := file.CreateUploadFolder(uploadReq.Address)
		if err != nil {
			logs.Error("file.CreateUploadFolder error: %v", err)
			resp.Error = &modelBase.Error{
				Message: errors.CREATE_UPLOAD_FOLDER_ERROR_MESSAGE,
			}
			c.Data["json"] = resp
			c.ServeJSON()
			return
		}

		for i := 1; i <= 9; i++ {
			fileName := fmt.Sprintf("%d.tmp", i)
			err = minio.Upload(folderName, fileName)
			if err != nil {
				logs.Error("minio.Upload error: %v", err)
				resp.Error = &modelBase.Error{
					Message: errors.UPLOAD_FILE_ERROR_MESSAGE,
				}
				c.Data["json"] = resp
				c.ServeJSON()
				return
			}
		}

		err = blockchain.CreateRecord(uploadReq.Address, uploadReq.PrivateKey, folderName)
		if err != nil {
			logs.Error("blockchain.CreateRecord error: %v", err)
			resp.Error = &modelBase.Error{
				Message: errors.CREATE_BLOCKCHAIN_RECORD_ERROR_MESSAGE,
			}
			c.Data["json"] = resp
			c.ServeJSON()
			return
		}

		//resp.ObjectUUID = folderName
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

// @router /download [get]
func (c *MetadataController) Download() {
	var resp modelBase.Resp

	objectName := c.Ctx.Input.Query("object_name")
	if objectName == "" {
		resp.Error = &modelBase.Error{
			Message: errors.DOWNLOAD_FILE_PARAMS_IS_EMPTY_ERROR_MESSAGE,
		}
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}

	localPath, err := minio.Download(objectName)
	if err != nil {
		logs.Error("minio.Download error: %v", err)
		resp.Error = &modelBase.Error{
			Message: errors.DOWNLOAD_FILE_ERROR_MESSAGE,
		}
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}

	// 返回下载的文件给浏览器
	c.Ctx.Output.Download(localPath)
}

// @router /listFolders [post]
func (c *MetadataController) ListFolders() {
	var resp modelMetadata.ListResp

	// 获取请求体的原始字节数据
	requestBody := c.Ctx.Input.RequestBody

	// 解析请求体为RequestBody结构体
	var listFoldersReq modelMetadata.ListFoldersReq
	err := json.Unmarshal(requestBody, &listFoldersReq)
	if err != nil {
		logs.Error("json.Unmarshal error: %v", err)
		resp.Error = &modelBase.Error{
			Message: errors.LIST_OBJECTS_PARAMS_READ_ERROR_MESSAGE,
		}
	} else {
		address := listFoldersReq.Address
		if address == "" {
			resp.Error = &modelBase.Error{
				Message: errors.LIST_OBJECTS_PARAMS_IS_EMPTY_ERROR_MESSAGE,
			}
		} else {
			objectInfo, err := minio.ListFolders(address)
			if err != nil {
				logs.Error("minio.ListFolders error: %v", err)
				resp.Error = &modelBase.Error{
					Message: errors.LIST_OBJECTS_GET_ERROR_MESSAGE,
				}
			} else {
				for _, object := range objectInfo {
					// 字符串转int64
					createAtString := strings.Split(object.Key, ".")[1]
					createAt, _ := strconv.ParseInt(createAtString[:len(createAtString)-1], 10, 64)
					resp.Data = append(resp.Data, modelMetadata.ListData{
						FileName: object.Key,
						FileSize: object.Size,
						CreateAt: createAt,
					})
				}

				// 按照创建时间排序
				sort.Slice(resp.Data, func(i, j int) bool {
					return resp.Data[i].CreateAt > resp.Data[j].CreateAt
				})
			}
		}
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

// @router /listFiles [post]
func (c *MetadataController) ListFiles() {
	var resp modelMetadata.ListResp

	// 获取请求体的原始字节数据
	requestBody := c.Ctx.Input.RequestBody

	// 解析请求体为RequestBody结构体
	var listFilesReq modelMetadata.ListFilesReq
	err := json.Unmarshal(requestBody, &listFilesReq)
	if err != nil {
		logs.Error("json.Unmarshal error: %v", err)
		resp.Error = &modelBase.Error{
			Message: errors.LIST_OBJECTS_PARAMS_READ_ERROR_MESSAGE,
		}
	} else {
		objectInfo, err := minio.ListFiles(listFilesReq.ObjectUUID)
		if err != nil {
			logs.Error("minio.ListFiles error: %v", err)
			resp.Error = &modelBase.Error{
				Message: errors.LIST_OBJECTS_GET_ERROR_MESSAGE,
			}
		} else {
			for _, object := range objectInfo {
				resp.Data = append(resp.Data, modelMetadata.ListData{
					FileName: object.Key,
					FileSize: object.Size,
					CreateAt: object.LastModified.Unix(),
				})
			}

			////按照创建时间排序
			//sort.Slice(resp.Data, func(i, j int) bool {
			//	return resp.Data[i].CreateAt > resp.Data[j].CreateAt
			//})
		}
	}

	c.Data["json"] = resp
	c.ServeJSON()
}
