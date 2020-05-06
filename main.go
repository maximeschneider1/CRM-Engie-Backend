package main

import (
	"data-back-real/model"
	"data-back-real/handler"

)

func main() {
	handler.StartWebServer()
}









var _ = model.ClientTotalInfo{
	ClientDetail:         model.ClientDetail{
		Name:           "",
		Phone:          "",
		ChallengesDone: 0,
		Address:        "",
		Score:          0,
		FamilyMembers:  3,
		Panels:         2,
		HeavyEquipment: 2,
		SunLevel:       2,
	},
	ProductionInfo:     model.ProductionInfo{
		FromGenToConsumer:  1500,
		FromGenToGrid:      300,
		FromGridToConsumer: 1000,
	},
	ProductionQueryInfos: model.ProductionQueryInfos{
		ClientID:      "",
		ClientAddress: "",
		StartDate:     "",
		EndDate:       "",
	},
}