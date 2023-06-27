package modelBlockchain

import modelBase "gitlab.dev.21vianet.com/liu.hao8/metatoc-service/models/base"

type ListRecordReq struct {
	Address string `json:"address" valid:"Required"`
}

type ListRecordResp struct {
	modelBase.Resp
	Data ListRecordInventory `json:"data,omitempty"`
}

type ListRecordBlockchainResp struct {
	modelBase.BlockchainResp
	Data ListRecordInventory `json:"data,omitempty"`
}

type ListRecordInventory struct {
	Paths []string `json:"paths"`
}
