package modelBlockchain

import (
	modelBase "gitlab.dev.21vianet.com/liu.hao8/metatoc-service/models/base"
)

type SignupResp struct {
	modelBase.Resp
	Data SignupInventory `json:"data,omitempty"`
}

type SignupBlockchainResp struct {
	modelBase.BlockchainResp
	Data SignupInventory `json:"data,omitempty"`
}

type SignupInventory struct {
	Address    string `json:"address"`
	PrivateKey string `json:"private_key"`
}
