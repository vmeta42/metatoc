package blockchain

import (
	"fmt"
	"net/http"
	"os"

	"gitlab.dev.21vianet.com/liu.hao8/metatoc-service/utils/common"

	"gitlab.dev.21vianet.com/liu.hao8/metatoc-service/consts"

	modelBase "gitlab.dev.21vianet.com/liu.hao8/metatoc-service/models/base"

	modelBlockchain "gitlab.dev.21vianet.com/liu.hao8/metatoc-service/models/blockchain"

	myHttp "gitlab.dev.21vianet.com/liu.hao8/metatoc-service/utils/http"
)

const (
	DEFAULT_BLOCKCHAIN_WEBSERVICE_ADDRESS = "http://172.22.50.211:5000"
)

func InitBlockchainWebserviceAddress() string {
	if os.Getenv("BLOCKCHAIN_WEBSERVICE_ADDRESS") != "" {
		return os.Getenv("BLOCKCHAIN_WEBSERVICE_ADDRESS")
	} else {
		return DEFAULT_BLOCKCHAIN_WEBSERVICE_ADDRESS
	}
}

var (
	blockchainWebserviceAddress = InitBlockchainWebserviceAddress()
)

func Signup() (*modelBlockchain.SignupBlockchainResp, error) {
	newHttp := myHttp.New()
	newHttp.Method = http.MethodGet
	newHttp.Url = fmt.Sprintf("%s/signup", blockchainWebserviceAddress)
	result := &modelBlockchain.SignupBlockchainResp{}
	if err := newHttp.SendRequest(result); err != nil {
		return nil, err
	}
	if result.Code != 0 {
		return nil, fmt.Errorf(result.Message)
	} else {
		return result, nil
	}
}

// CreateRecord 创建区块链记录
func CreateRecord(address string, privateKey string, objectUUID string) error {
	newHttp := myHttp.New()
	newHttp.Method = http.MethodPost
	newHttp.Url = fmt.Sprintf("%s/paths", blockchainWebserviceAddress)
	newHttp.Data = map[string]interface{}{
		"address":     address,
		"private_key": privateKey,
		"path":        fmt.Sprintf("%s%s", consts.HDFS_PATH_PREFIX, common.GenerateRandomString(8)),
		//"path":        fmt.Sprintf("%s%s", consts.HDFS_PATH_PREFIX, objectUUID),
		//"path":    fmt.Sprintf("%s", objectUUID),
		"content": objectUUID,
	}
	result := &modelBase.BlockchainResp{}
	if err := newHttp.SendRequest(result); err != nil {
		return err
	}
	if result.Code != 0 {
		return fmt.Errorf(result.Message)
	} else {
		return nil
	}
}

func ShareRecord(address string, toAddress string, privateKey string, path string) (*modelBlockchain.ShareRecordBlockchainResp, error) {
	newHttp := myHttp.New()
	newHttp.Method = http.MethodPut
	newHttp.Url = fmt.Sprintf("%s/paths", blockchainWebserviceAddress)
	newHttp.Data = map[string]interface{}{
		"from_address": address,
		"to_address":   toAddress,
		"private_key":  privateKey,
		"token_name":   fmt.Sprintf("%s%s", consts.HDFS_PATH_PREFIX, path),
	}
	result := &modelBlockchain.ShareRecordBlockchainResp{}
	if err := newHttp.SendRequest(result); err != nil {
		return nil, err
	}
	if result.Code != 0 {
		return nil, fmt.Errorf(result.Message)
	} else {
		return result, nil
	}
}

func ListRecord(address string) (*modelBlockchain.ListRecordBlockchainResp, error) {
	newHttp := myHttp.New()
	newHttp.Method = http.MethodGet
	newHttp.Url = fmt.Sprintf("%s/paths", blockchainWebserviceAddress)
	newHttp.Params = []map[string]interface{}{
		{
			"key":   "address",
			"value": address,
		},
	}
	result := &modelBlockchain.ListRecordBlockchainResp{}
	if err := newHttp.SendRequest(result); err != nil {
		return nil, err
	}
	if result.Code != 0 {
		return nil, fmt.Errorf(result.Message)
	} else {
		return result, nil
	}
}

func ViewRecord(address string, privateKey string, path string) (*modelBlockchain.ViewRecordBlockchainResp, error) {
	newHttp := myHttp.New()
	newHttp.Method = http.MethodGet
	newHttp.Url = fmt.Sprintf("%s/paths/%s", blockchainWebserviceAddress, fmt.Sprintf("%s%s", consts.HDFS_PATH_PREFIX, path))
	newHttp.Params = []map[string]interface{}{
		{
			"key":   "address",
			"value": address,
		},
		//{
		//	"key":   "hdfs_path",
		//	"value": fmt.Sprintf("%s%s", consts.HDFS_PATH_PREFIX, path),
		//},
	}
	newHttp.Data = map[string]interface{}{
		"private_key": privateKey,
	}
	result := &modelBlockchain.ViewRecordBlockchainResp{}
	if err := newHttp.SendRequest(result); err != nil {
		return nil, err
	}
	if result.Code != 0 {
		return nil, fmt.Errorf(result.Message)
	} else {
		return result, nil
	}
}
