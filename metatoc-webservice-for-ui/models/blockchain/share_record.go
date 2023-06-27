package modelBlockchain

import modelBase "gitlab.dev.21vianet.com/liu.hao8/metatoc-service/models/base"

type ShareRecordReq struct {
	Address    string `json:"address" valid:"Required"`
	ToAddress  string `json:"to_address" valid:"Required"`
	PrivateKey string `json:"private_key" valid:"Required"`
	Path       string `json:"path" valid:"Required"`
}

type ShareRecordResp struct {
	modelBase.Resp
	//Data ListRecordInventory `json:"data"`
}

type ShareRecordBlockchainResp struct {
	modelBase.BlockchainResp
	//Data ListRecordInventory `json:"data"`
}
