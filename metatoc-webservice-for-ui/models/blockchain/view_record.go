package modelBlockchain

import modelBase "gitlab.dev.21vianet.com/liu.hao8/metatoc-service/models/base"

type ViewRecordReq struct {
	Address    string `json:"address" valid:"Required"`
	PrivateKey string `json:"private_key" valid:"Required"`
	Path       string `json:"path" valid:"Required"`
}

type ViewRecordResp struct {
	modelBase.Resp
	Data string `json:"data,omitempty"`
}

type ViewRecordBlockchainResp struct {
	modelBase.BlockchainResp
	Data ViewRecordInventory `json:"data,omitempty"`
}

type ViewRecordInventory struct {
	Data string `json:"data"`
}
