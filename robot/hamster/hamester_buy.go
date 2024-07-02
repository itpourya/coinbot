package hamster

import (
	"CoinBot/internal/serilizers/buy"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	BuyApi           = "https://api.hamsterkombat.io/clicker/buy-upgrade"
	UpgradeForBuyAPI = "https://api.hamsterkombat.io/clicker/upgrades-for-buy"
)

func HamBuy(profitId []buy.UpgradeId) {
	var wg sync.WaitGroup

	for _, pID := range profitId {
		wg.Add(1)
		go func() {
			buyProcess(pID, &wg)
		}()
	}

	wg.Wait()
}

func HamJsonUpdate() {
	log.Println("Start improve profit")

	req, err := http.NewRequest("POST", UpgradeForBuyAPI, nil)
	if err != nil {
		log.Println("HamJsonUpdate: ", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", TOKEN)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("HamJsonUpdate: ", err)
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Println("HamJsonUpdate: ", err)
		}
	}(resp.Body)

	newData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("HamJsonUpdate: ", err)
	}
	err = ioutil.WriteFile("./../pkg/data.json", newData, 0644)
	if err != nil {
		log.Println("HamJsonUpdate: ", err)
	}

	log.Println("HamJsonUpdate : Json File Updated")
}

func buyProcess(pID buy.UpgradeId, wg *sync.WaitGroup) {
	defer wg.Done()

	jsonData, err := json.Marshal(map[string]interface{}{
		"upgradeId": string(pID.Id),
		"timestamp": timestamppb.Now().Seconds,
	})
	if err != nil {
		log.Fatal("HamBuy: ", err)
	}

	req, err := http.NewRequest("POST", BuyApi, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("HamBuy: ", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", TOKEN)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("HamBuy: ", err)
	}

	log.Println("HamBuy : ", pID.Id, pID.Level)

	err = resp.Body.Close()
	if err != nil {
		log.Fatal("HamBuy: ", err)
	}
}
