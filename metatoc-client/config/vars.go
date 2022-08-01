package config

import (
	"fmt"
	"os"
)

var (
	MetatocWebServiceAddress = "172.22.50.211:5000"
	MetatocNatsAddress       = "172.22.50.211:4222"
)

func init() {
	if os.Getenv("METATOC_WEBSERVICE_ADDRESS") != "" {
		MetatocWebServiceAddress = fmt.Sprintf("http://%s", os.Getenv("METATOC_WEBSERVICE_ADDRESS"))
	} else {
		MetatocWebServiceAddress = fmt.Sprintf("http://%s", MetatocWebServiceAddress)
	}
	if os.Getenv("METATOC_NATS_ADDRESS") != "" {
		MetatocNatsAddress = fmt.Sprintf("nats://%s", os.Getenv("METATOC_NATS_ADDRESS"))
	} else {
		MetatocNatsAddress = fmt.Sprintf("nats://%s", MetatocNatsAddress)
	}
}
