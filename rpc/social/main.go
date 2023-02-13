package main

import (
	social "douyin/kitex_gen/social/socialservice"
	social2 "douyin/rpc/social"
	"log"
)

func main() {
	svr := social.NewServer(new(social2.SocialServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
