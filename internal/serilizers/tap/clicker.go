package serilizers

type ClickerUser struct {
	ID             string                 `json:"id"`
	TotalCoins     float64                `json:"totalCoins"`
	BalanceCoins   float64                `json:"balanceCoins"`
	Level          int                    `json:"level"`
	AvailableTaps  int                    `json:"availableTaps"`
	LastSyncUpdate int                    `json:"lastSyncUpdate"`
	ExchangeID     string                 `json:"exchangeId"`
	Boosts         map[string]interface{} `json:"boosts"`
	Upgrades       map[string]struct {
		ID                     string `json:"id"`
		Level                  int    `json:"level"`
		LastUpgradeAt          int    `json:"lastUpgradeAt"`
		SnapshotReferralsCount int    `json:"snapshotReferralsCount"`
	} `json:"upgrades"`
	Tasks map[string]struct {
		ID          string `json:"id"`
		CompletedAt string `json:"completedAt"`
		Days        int    `json:"days"`
	} `json:"tasks"`
	AirdropTasks       map[string]interface{} `json:"airdropTasks"`
	ReferralsCount     int                    `json:"referralsCount"`
	MaxTaps            int                    `json:"maxTaps"`
	EarnPerTap         int                    `json:"earnPerTap"`
	EarnPassivePerSec  float64                `json:"earnPassivePerSec"`
	EarnPassivePerHour float64                `json:"earnPassivePerHour"`
	LastPassiveEarn    float64                `json:"lastPassiveEarn"`
	TapsRecoverPerSec  int                    `json:"tapsRecoverPerSec"`
	Referral           struct {
		Friend struct {
			ID                int64  `json:"id"`
			IsBot             bool   `json:"isBot"`
			FirstName         string `json:"firstName"`
			LastName          string `json:"lastName"`
			LanguageCode      string `json:"languageCode"`
			WelcomeBonusCoins int    `json:"welcomeBonusCoins"`
		} `json:"friend"`
	} `json:"referral"`
}
