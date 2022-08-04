package cmd

import (
	"fmt"
	"github.com/desertbit/grumble"
	"github.com/vmeta42/metatoc/metatoc-client/config"
	"github.com/vmeta42/metatoc/metatoc-client/utils"
	"net/http"
)

type Detail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Data string `json:"data"`
	} `json:"data"`
}

func init() {
	App.AddCommand(&grumble.Command{
		Name: "detail",
		Help: "return the detail of data related to HDFS path",
		Flags: func(f *grumble.Flags) {
			f.String("p", "hdfs_path", "", "HDFS resource path id")
			f.String("a", "address", "", "wallet address")
			f.String("k", "key", "", "wallet private key")
		},
		Args: func(a *grumble.Args) {

		},
		Run: func(c *grumble.Context) error {
			hdfsPath := c.Flags.String("hdfs_path")
			if hdfsPath == "" {
				return fmt.Errorf("HDFS path cannot be empty")
			}
			newHttp := utils.NewHttp()
			newHttp.Method = http.MethodGet
			newHttp.Url = fmt.Sprintf("%s/paths/%s", config.MetatocWebServiceAddress, hdfsPath)
			newHttp.Params = []map[string]interface{}{
				{
					"key":   "hdfs_path",
					"value": c.Flags.String("hdfs_path"),
				},
				{
					"key":   "address",
					"value": c.Flags.String("address"),
				},
			}
			newHttp.Data = map[string]interface{}{
				"private_key": c.Flags.String("key"),
			}
			result := &Detail{}
			if err := newHttp.SendRequest(result); err != nil {
				return err
			}
			if result.Code != 0 {
				return fmt.Errorf(result.Message)
			} else {
				fmt.Println(result.Message)
				fmt.Printf("data is [%s]\n", result.Data.Data)
			}
			return nil
		},
	})
}
