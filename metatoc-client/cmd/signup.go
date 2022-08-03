package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/desertbit/grumble"
	"github.com/vmeta42/metatoc/metatoc-client/config"
	"github.com/vmeta42/metatoc/metatoc-client/utils"
	"net/http"
	"os"
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

				if _, err := os.Stat("./json"); err != nil {
					os.Mkdir("./json", os.ModePerm)
				}
				file, _ := os.OpenFile("./json/wallet.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
				defer file.Close()
				writer := bufio.NewWriter(file)
				resultJson, _ := json.Marshal(result)
				writer.WriteString(fmt.Sprintf("%s\r\n", string(resultJson)))
				writer.Flush()
			}

			return nil
		},
	})
}
