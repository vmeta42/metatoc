package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/desertbit/grumble"
	"github.com/nats-io/nats.go"
	"github.com/vmeta42/metatoc/metatoc-client/config"
	"net/http"
)

type Create struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func init() {
	App.AddCommand(&grumble.Command{
		Name: "create",
		Help: "create a new block on chain with new HDFS path",
		Flags: func(f *grumble.Flags) {
			f.String("k", "key", "", "wallet private key")
			f.String("a", "address", "", "wallet address")
			f.String("p", "path", "", "the stored HDFS path")
			f.String("c", "content", "", "contents of the stored HDFS path")
		},
		Args: func(a *grumble.Args) {

		},
		Run: func(c *grumble.Context) error {
			//newHttp := utils.NewHttp()
			//newHttp.Method = http.MethodPost
			//newHttp.Url = fmt.Sprintf("%s/paths", config.MetatocWebServiceAddress)
			//newHttp.Data = map[string]interface{}{
			//	"private_key": c.Flags.String("key"),
			//	"address":     c.Flags.String("address"),
			//	"path":        c.Flags.String("path"),
			//	"content":     c.Flags.String("content"),
			//}
			//result := &Create{}
			//if err := newHttp.SendRequest(result); err != nil {
			//	return err
			//}
			//if result.Code != 0 {
			//	return fmt.Errorf(result.Message)
			//} else {
			//	fmt.Println(result.Message)
			//}

			publishDataMap := make(map[string]interface{})
			publishDataMap["consumer"] = "sendRequest"
			publishDataMap["type"] = http.MethodPost
			publishDataMap["url"] = fmt.Sprintf("%s/paths", config.MetatocWebServiceAddress)
			publishDataMap["data"] = make(map[string]interface{})
			publishDataMap["data"] = map[string]interface{}{
				"private_key": c.Flags.String("key"),
				"address":     c.Flags.String("address"),
				"path":        c.Flags.String("path"),
				"content":     c.Flags.String("content"),
			}
			publishDataString, _ := json.Marshal(publishDataMap)

			nc, _ := nats.Connect(config.MetatocNatsAddress)
			defer nc.Close()

			js, _ := nc.JetStream(nats.PublishAsyncMaxPending(256))
			js.Publish("CONSUMER_METATOC.PATHS", publishDataString)

			fmt.Println("SUCCESSFUL")
			return nil
		},
	})
}
