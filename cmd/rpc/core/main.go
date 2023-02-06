package main

import (
	core "douyin/kitex_gen/core/coreservice"
	"log"
)

func main() {
	svr := core.NewServer(new(CoreServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
