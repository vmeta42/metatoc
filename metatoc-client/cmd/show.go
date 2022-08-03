package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/desertbit/grumble"
	"os"
	"strings"
)

func init() {
	App.AddCommand(&grumble.Command{
		Name: "show",
		Help: "show wallet of created",
		Flags: func(f *grumble.Flags) {

		},
		Args: func(a *grumble.Args) {

		},
		Run: func(c *grumble.Context) error {
			if content, err := os.ReadFile("./json/wallet.json"); err == nil {
				noData := true
				contentSlice := strings.Split(string(content), "\r\n")
				for index, value := range contentSlice {
					signup := Signup{}
					if err = json.Unmarshal([]byte(value), &signup); err == nil {
						if signup.Message == "ok" {
							noData = false
							if index == 0 {
								fmt.Printf("There are %d pieces of data\n", len(contentSlice))
								fmt.Printf("The %dst data: address is [%s], private key is [%s]\n", index+1, signup.Data.Address, signup.Data.PrivateKey)
							} else if index == 1 {
								fmt.Printf("The %dnd data: address is [%s], private key is [%s]\n", index+1, signup.Data.Address, signup.Data.PrivateKey)
							} else if index == 2 {
								fmt.Printf("The %drd data: address is [%s], private key is [%s]\n", index+1, signup.Data.Address, signup.Data.PrivateKey)
							} else {
								fmt.Printf("The %dth data: address is [%s], private key is [%s]\n", index+1, signup.Data.Address, signup.Data.PrivateKey)
							}
						}
					}
				}
				if noData == true {
					fmt.Println("no data")
				}
			} else {
				fmt.Println("no data")
			}
			return nil
		},
	})
}
