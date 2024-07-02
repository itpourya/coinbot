package main

import (
	"CoinBot/internal/database"
	"CoinBot/internal/repository"
	"CoinBot/pkg"
	"CoinBot/robot/hamster"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

func main() {
	rdp := database.NewCache()
	client := repository.NewCacheRepository(rdp)

	log.Println("BOT START")

	cr := cron.New()
	cr.AddFunc("@every 3s", func() {
		status, msg := client.POP()
		if msg == "command" && status {
			hamster.HamJsonUpdate()
			list := pkg.GetHamProfitId(1000000)
			if len(list) != 0 {
			loop:
				for {
					hamster.HamBuy(list)
					check := pkg.GetHamProfitId(1000000)
					if len(check) == 0 {
						break loop
					}
					time.Sleep(10 * time.Second)
				}
			}

			log.Println("improve profit ended!")

			log.Println("tap hamster start")
			_, err := hamster.HamTap()
			if err != nil {
				return
			}
		}
	})

	cr.Start()

	select {}
}
