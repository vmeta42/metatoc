package controllers

import (
	"encoding/json"
	"strings"

	"gitlab.dev.21vianet.com/liu.hao8/metatoc-service/consts"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"gitlab.dev.21vianet.com/liu.hao8/metatoc-service/errors"
	modelBase "gitlab.dev.21vianet.com/liu.hao8/metatoc-service/models/base"
	modelBlockchain "gitlab.dev.21vianet.com/liu.hao8/metatoc-service/models/blockchain"
	"gitlab.dev.21vianet.com/liu.hao8/metatoc-service/utils/blockchain"
)

type BlockchainController struct {
	beego.Controller
}

// @router /signup [post]
func (c *BlockchainController) Signup() {
	var resp modelBlockchain.SignupResp

	result, err := blockchain.Signup()
	if err != nil {
		logs.Error("blockchain.Signup error: %v", err)
		resp.Error = &modelBase.Error{
			Message: errors.SIGNUP_WALLET_ADDRESS_ERROR_MESSAGE,
		}
	} else {
		resp.Data.Address = result.Data.Address
		resp.Data.PrivateKey = result.Data.PrivateKey
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

// @router /share [post]
func (c *BlockchainController) Share() {
	var resp modelBlockchain.ShareRecordResp

	// 获取请求体的原始字节数据
	requestBody := c.Ctx.Input.RequestBody

	// 解析请求体为RequestBody结构体
	var shareRecordReq modelBlockchain.ShareRecordReq
	err := json.Unmarshal(requestBody, &shareRecordReq)
	if err != nil {
		logs.Error("json.Unmarshal error: %v", err)
		resp.Error = &modelBase.Error{
			Message: errors.SHARE_RECORD_PARAMS_READ_ERROR_MESSAGE,
		}
	} else {
		result, err := blockchain.ShareRecord(shareRecordReq.Address, shareRecordReq.ToAddress, shareRecordReq.PrivateKey, shareRecordReq.Path)
		if err != nil {
			logs.Error("blockchain.ShareRecord error: %v", err)
			resp.Error = &modelBase.Error{
				Message: errors.SHARE_RECORD_ERROR_MESSAGE,
			}
		} else {
			resp.Data = result.Data
		}
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

// @router /list [post]
func (c *BlockchainController) List() {
	var resp modelBlockchain.ListRecordResp

	// 获取请求体的原始字节数据
	requestBody := c.Ctx.Input.RequestBody

	// 解析请求体为RequestBody结构体
	var listRecordReq modelBlockchain.ListRecordReq
	err := json.Unmarshal(requestBody, &listRecordReq)
	if err != nil {
		logs.Error("json.Unmarshal error: %v", err)
		resp.Error = &modelBase.Error{
			Message: errors.LIST_RECORD_PARAMS_READ_ERROR_MESSAGE,
		}
	} else {
		result, err := blockchain.ListRecord(listRecordReq.Address)
		if err != nil {
			logs.Error("blockchain.ListRecord error: %v", err)
			resp.Error = &modelBase.Error{
				Message: errors.LIST_RECORD_ERROR_MESSAGE,
			}
		} else {
			for _, path := range result.Data.Paths {
				if strings.Index(path, consts.HDFS_PATH_PREFIX) == 0 {
					resp.Data.Paths = append(resp.Data.Paths, strings.Replace(path, consts.HDFS_PATH_PREFIX, "", 1))
				}
			}
			//resp.Data.Paths = result.Data.Paths
		}
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

// @router /view [post]
func (c *BlockchainController) View() {
	var resp modelBlockchain.ViewRecordResp

	// 获取请求体的原始字节数据
	requestBody := c.Ctx.Input.RequestBody

	// 解析请求体为RequestBody结构体
	var viewRecordReq modelBlockchain.ViewRecordReq
	err := json.Unmarshal(requestBody, &viewRecordReq)
	if err != nil {
		logs.Error("json.Unmarshal error: %v", err)
		resp.Error = &modelBase.Error{
			Message: errors.VIEW_RECORD_PARAMS_READ_ERROR_MESSAGE,
		}
	} else {
		result, err := blockchain.ViewRecord(viewRecordReq.Address, viewRecordReq.PrivateKey, viewRecordReq.Path)
		if err != nil {
			logs.Error("blockchain.ViewRecord error: %v", err)
			resp.Error = &modelBase.Error{
				Message: errors.VIEW_RECORD_ERROR_MESSAGE,
			}
		} else {
			resp.Data = result.Data.Data
		}
	}

	c.Data["json"] = resp
	c.ServeJSON()
}
