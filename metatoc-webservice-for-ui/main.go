package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "gitlab.dev.21vianet.com/liu.hao8/metatoc-service/routers"
	"gitlab.dev.21vianet.com/liu.hao8/metatoc-service/utils/minio"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	err := minio.CreateBucket(minio.InitVarBucketName())
	if err != nil {
		panic(err)
	}

	beego.Run()
}
