package cmd

import (
	"fmt"
	"github.com/desertbit/grumble"
	"github.com/vmeta42/metatoc/metatoc-client/config"
	"github.com/vmeta42/metatoc/metatoc-client/utils"
	"net/http"
)

type List struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Paths []string `json:"paths"`
	} `json:"data"`
}

func init() {
	App.AddCommand(&grumble.Command{
		Name: "list",
		Help: "return a list of HDFS resource paths",
		Flags: func(f *grumble.Flags) {
			f.String("a", "address", "", "wallet address")
			f.Int("o", "offset", 0, "specifies the page number of the hdfs paths to be display")
			f.Int("l", "limit", 10, "limits the number of items on a page")
		},
		Args: func(a *grumble.Args) {

		},
		Run: func(c *grumble.Context) error {
			newHttp := utils.NewHttp()
			newHttp.Method = http.MethodGet
			newHttp.Url = fmt.Sprintf("%s/paths", config.MetatocWebServiceAddress)
			newHttp.Params = []map[string]interface{}{
				{
					"key":   "address",
					"value": c.Flags.String("address"),
				},
				{
					"key":   "offset",
					"value": c.Flags.Int("offset"),
				},
				{
					"key":   "limit",
					"value": c.Flags.Int("limit"),
				},
			}
			result := &List{}
			if err := newHttp.SendRequest(result); err != nil {
				return err
			}
			if result.Code != 0 {
				return fmt.Errorf(result.Message)
			} else {
				fmt.Println(result.Message)
				if len(result.Data.Paths) > 0 {
					fmt.Printf("There are %d pieces of data\n", len(result.Data.Paths))
					for index, value := range result.Data.Paths {
						if index == 0 {
							fmt.Printf("The %dst data is [%s]\n", index+1, value)
						} else if index == 1 {
							fmt.Printf("The %dnd data is [%s]\n", index+1, value)
						} else if index == 2 {
							fmt.Printf("The %drd data is [%s]\n", index+1, value)
						} else {
							fmt.Printf("The %dth data is [%s]\n", index+1, value)
						}
					}
				} else {
					fmt.Println("no data")
				}
			}
			return nil
		},
	})
}
