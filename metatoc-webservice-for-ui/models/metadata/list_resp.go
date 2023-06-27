package modelMetadata

import (
	modelBase "gitlab.dev.21vianet.com/liu.hao8/metatoc-service/models/base"
)

type ListResp struct {
	modelBase.Resp
	Data []ListData `json:"data,omitempty"`
}

type ListData struct {
	FileName string `json:"file_name"`
	FileSize int64  `json:"file_size"`
	CreateAt int64  `json:"create_at"`
}
