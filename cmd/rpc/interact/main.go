package main

import (
	interact "douyin/kitex_gen/interact/interactservice"
	"log"
)

func main() {
	svr := interact.NewServer(new(InteractServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
