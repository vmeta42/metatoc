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
		DataPath string `json:"data_path"`
	} `json:"data"`
}

func init() {
	App.AddCommand(&grumble.Command{
		Name: "detail",
		Help: "return the detail of data related to HDFS path",
		Flags: func(f *grumble.Flags) {
			f.String("p", "hdfs_path", "", "HDFS resource path id")
			f.String("a", "address", "", "wallet address")
		},
		Args: func(a *grumble.Args) {

		},
		Run: func(c *grumble.Context) error {
			hdfsPath := c.Flags.String("hdfs_path")
			if hdfsPath == "" {
				return fmt.Errorf("hdfs_path cannot be emppty")
			}
			newHttp := utils.NewHttp()
			newHttp.Method = http.MethodGet
			newHttp.Url = fmt.Sprintf("%s/paths/%s", config.MetatocWebServiceAddress, hdfsPath)
			newHttp.Params = []map[string]interface{}{
				{
					"key":   "address",
					"value": c.Flags.String("address"),
				},
			}
			result := &Detail{}
			if err := newHttp.SendRequest(result); err != nil {
				return err
			}
			if result.Code != 0 {
				return fmt.Errorf(result.Message)
			} else {
				fmt.Println(result.Message)
				fmt.Printf("data path is [%s]\n", result.Data.DataPath)
			}
			return nil
		},
	})
}
