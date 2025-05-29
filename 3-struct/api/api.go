package api

import (
	"fmt"

	"pferdina.com/3-struct/config"
)

func GetEnv() {
	config := config.NewConfig()
	fmt.Println(config.Key)
}
