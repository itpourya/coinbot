package buy

type Upgrade struct {
	Id                 string
	Level              float64
	Name               string
	Price              int
	ProfitPerHourDelta int
	IsExpired          bool
	IsAvailable        bool
}

type UpgradeId struct {
	Id    string
	Level float64
}
