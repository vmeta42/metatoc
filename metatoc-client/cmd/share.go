package cmd

import (
	"fmt"
	"github.com/desertbit/grumble"
	"github.com/vmeta42/metatoc/metatoc-client/config"
	"github.com/vmeta42/metatoc/metatoc-client/utils"
	"net/http"
)

type Share struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func init() {
	App.AddCommand(&grumble.Command{
		Name: "share",
		Help: "share block with other wallet",
		Flags: func(f *grumble.Flags) {
			f.String("k", "key", "", "wallet private key")
			f.String("f", "from_address", "", "from which wallet address to share")
			f.String("o", "to_address", "", "where is the wallet address shared")
			f.String("t", "token", "", "token when sharing wallet address")
		},
		Args: func(a *grumble.Args) {

		},
		Run: func(c *grumble.Context) error {
			newHttp := utils.NewHttp()
			newHttp.Method = http.MethodPut
			newHttp.Url = fmt.Sprintf("%s/paths", config.MetatocWebServiceAddress)
			newHttp.Data = map[string]interface{}{
				"private_key":  c.Flags.String("key"),
				"from_address": c.Flags.String("from_address"),
				"to_address":   c.Flags.String("to_address"),
				"token_name":   c.Flags.String("token"),
			}
			result := &Share{}
			if err := newHttp.SendRequest(result); err != nil {
				return err
			}
			if result.Code != 0 {
				return fmt.Errorf(result.Message)
			} else {
				fmt.Println(result.Message)
			}
			return nil
		},
	})
}
