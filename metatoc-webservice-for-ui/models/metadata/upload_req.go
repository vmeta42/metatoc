package modelMetadata

type UploadReq struct {
	Address    string `json:"address" valid:"Required"`
	PrivateKey string `json:"private_key" valid:"Required"`
}
