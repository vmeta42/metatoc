package cmd

import (
	"fmt"
	"github.com/desertbit/grumble"
	"github.com/vmeta42/metatoc/metatoc-client/config"
	"github.com/vmeta42/metatoc/metatoc-client/utils"
	"net/http"
)

type Signup struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Address    string `json:"address"`
		PrivateKey string `json:"private_key"`
	} `json:"data"`
}

func init() {
	App.AddCommand(&grumble.Command{
		Name: "signup",
		Help: "create a new wallet",
		Flags: func(f *grumble.Flags) {

		},
		Args: func(a *grumble.Args) {

		},
		Run: func(c *grumble.Context) error {
			newHttp := utils.NewHttp()
			newHttp.Method = http.MethodGet
			newHttp.Url = fmt.Sprintf("%s/signup", config.MetatocWebServiceAddress)
			result := &Signup{}
			if err := newHttp.SendRequest(result); err != nil {
				return err
			}
			if result.Code != 0 {
				return fmt.Errorf(result.Message)
			} else {
				fmt.Println(result.Message)
				fmt.Printf("address is [%s]\n", result.Data.Address)
				fmt.Printf("private key is [%s]\n", result.Data.PrivateKey)
			}
			return nil
		},
	})
}
