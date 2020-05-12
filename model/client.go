package model

type ClientTotalInfo struct {
	ClientDetail
	ProductionInfo
	ProductionQueryInfos

}

type ClientDetail struct {
	ClientId         int `json:"client_id"`
	Name             string `json:"name"`
	Phone            string `json:"phone"`
	ChallengesDone   int    `json:"challengesDone"`
	Address          string `json:"address"`
	Score            int    `json:"score"`
	Gender string
	City string
	FamilyMembers    int
	Panels           int
	HeavyEquipment   int
	SunLevel         int
	BirthDate        string
	AccountCreation  string
	Email            string
	Trophees         int
	Ratio            string
	ProductionTotal  string
	LastBill         int
	RegistredDevices int
}

type ProductionInfo struct {
	FromGenToConsumer float64
	FromGenToGrid float64
	FromGridToConsumer float64
}

type ProductionStory struct {
	FromGenToConsumer []int
	FromGenToGrid []int
	FromGridToConsumer []int
	Score int
	AutoConsumption int
	TotalProduction int
}

type ProductionQueryInfos struct {
	ClientID string
	ClientAddress string
	StartDate string
	EndDate string
}