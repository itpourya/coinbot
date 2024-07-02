package hamster

import (
	"CoinBot/internal/serilizers/tap"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"

	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	HamAPI  = "https://api.hamsterkombat.io/clicker/tap"
	TOKEN   = "Bearer 1719140009665GgQHMyGHsMFTBlZiYW8kPXEs0SIrGH0ottTN2MkypQD41QohXlIEGzUnJMc3pieO6282056902"
	ResChan = make(chan *http.Response)
	ErrChan = make(chan string)
)

func HamTap() (bool, error) {
	var wg sync.WaitGroup

	availableTaps := getTaps()

	jsonData, err := json.Marshal(
		map[string]int{
			"count":         5,
			"availableTaps": availableTaps,
			"timestamp":     int(timestamppb.Now().Seconds),
		})
	if err != nil {
		return false, err
	}

	for i := 0; i < availableTaps; i++ {
		wg.Add(1)
		go func() {
			status, err := processRequest(jsonData, &wg)
			if !status && err != "" {
				log.Fatal("HamTap: ", err)
			}
		}()
	}

	wg.Wait()

	return true, nil
}

func processRequest(jsonData []byte, wg *sync.WaitGroup) (bool, string) {
	defer wg.Done()

	req, err := http.NewRequest("POST", HamAPI, bytes.NewBuffer(jsonData))
	if err != nil {
		ErrChan <- err.Error()
		return false, err.Error()
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", TOKEN)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ErrChan <- err.Error()
		return false, err.Error()
	}

	ResChan <- resp
	err = resp.Body.Close()
	if err != nil {
		return false, ""
	}

	return true, ""
}

func getTaps() int {

	jsonData, _ := json.Marshal(map[string]int{
		"count":         1,
		"availableTaps": 1500,
		"timestamp":     int(timestamppb.Now().Seconds),
	})

	req, err := http.NewRequest("POST", HamAPI, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("getTaps: ", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", TOKEN)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("getTaps: ", err.Error())
	}

	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, resp.Body)
	if err != nil {
		log.Fatal("getTaps: ", err.Error())
	}

	// Convert the byte slice to a string
	str := buffer.String()

	var response struct {
		ClickerUser serilizers.ClickerUser `json:"clickerUser"`
	}

	err = json.Unmarshal([]byte(str), &response)
	if err != nil {
		log.Fatal("getTaps: ", err.Error())
	}
	err = resp.Body.Close()
	if err != nil {
		log.Fatal("getTaps: ", err.Error())
	}

	return response.ClickerUser.AvailableTaps
}
