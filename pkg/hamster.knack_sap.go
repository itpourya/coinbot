package pkg

import (
	"CoinBot/internal/serilizers/buy"
	"encoding/json"
	"log"
	"os"
)

func GetHamProfitId(maxBudget int) []buy.UpgradeId {
	file, err := os.Open("./../pkg/data.json")
	if err != nil {
		log.Fatal("GetHamProfitId: Error: File not found or cannot be opened.")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("GetHamProfitId: Error: Can not close file.")
		}
	}(file)

	var data map[string]interface{}
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		log.Fatal("GetHamProfitId: Error: Invalid JSON file.")
	}

	upgrades := make([]buy.Upgrade, 0)
	for _, item := range data["upgradesForBuy"].([]interface{}) {
		itemMap := item.(map[string]interface{})
		upgrade := buy.Upgrade{
			Id:                 itemMap["id"].(string),
			Level:              itemMap["level"].(float64),
			Name:               itemMap["name"].(string),
			Price:              int(itemMap["price"].(float64)),
			ProfitPerHourDelta: int(itemMap["profitPerHourDelta"].(float64)),
			IsExpired:          itemMap["isExpired"].(bool),
			IsAvailable:        itemMap["isAvailable"].(bool),
		}
		if !upgrade.IsExpired && upgrade.IsAvailable {
			upgrades = append(upgrades, upgrade)
		}
	}

	list := handleUpgrades(upgrades, maxBudget)

	return list
}

func handleUpgrades(upgrades []buy.Upgrade, maxBudget int) []buy.UpgradeId {

	upgradeIds := make([]buy.UpgradeId, 0)
	countBuy := 0

	for _, item := range upgrades {
		price := item.Price - 200
		perHour := item.ProfitPerHourDelta * 20

		if countBuy <= maxBudget {
			if perHour >= price {
				upgradeId := buy.UpgradeId{
					Id:    item.Id,
					Level: float64(item.Level),
				}
				upgradeIds = append(upgradeIds, upgradeId)
			}
		}
	}

	return upgradeIds
}
