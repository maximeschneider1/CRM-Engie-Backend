package dao

import "data-back-real/model"

// Default variables are used if there are not enough info in DB to return coherent data to the user
var defaultScore = 150
var defaultAutoConsumption = 40
var defaultTags = []string{"Chauffage Gaz", "Famille Nombreuse", "Maison de vacance"}
var defaultTodo = []model.Todo{
	{
		Id:       1,
		Name: "Henri",
		Phone:    "0627654390",
		Category: "Client",
		Motif:    "Visite",
	}, {
		Id:       1,
		Name: "Henri",
		Phone:    "0627654390",
		Category: "Client",
		Motif:    "Visite",
	},{
		Id:       1,
		Name: "Henri",
		Phone:    "0627654390",
		Category: "Client",
		Motif:    "Visite",
	},
	{
		Name:       "Maxime",
		Id:         1,
		Phone:  "0987097654",
		Category:   "Lead",
		Motif:      "Production ",
	},
	{
		Name:       "Etienne",
		Id:         4,
		Phone:  "7812397654",
		Category:   "Client",
		Motif:      "Production ",
	},
	{
		Name:       "Camille",
		Id:         6,
		Phone:  "7812397654",
		Category:   "Client",
		Motif:      "Production ",
	},
}
var defaultProd = model.ProductionStory{
	FromGenToConsumer:  []int{1950, 2050, 2400, 1900, 2200, 2000, 2050, 2090},
	FromGenToGrid:      []int{850, 750, 500, 1000, 400, 600, 950, 590},
	FromGridToConsumer: []int{950, 1050, 1400, 900, 1000, 1000, 1050, 1090},
	Score:              defaultScore,
	AutoConsumption:    defaultAutoConsumption,
}
var defaultHomeInfo = model.HomeInfo{
TotalClients:   34,
Todo:           12,
NewLeads:       70,
NewDocuments:   3,
PotentialValue: 113879,
}
var defaultLeadDetail = model.Lead{
	Name:             "Antoine Dugomier",
	Phone:            "0987896543",
	Address:          "22 rue de la Moulaga",
	Score:              112,
	FirstContact:       "28/09/18",
	City:               "Clermont-Ferrand",
	ContentDownloaded:  5,
	TimeSpent:          32,
	OpenedEmails:       4,
	Profitability:      543,
	WeeksSinceInactive: 1,
}
var defaultLeadHistory = model.LeadHistory{
Type:    "Email",
Icon:    "mdi-phone",
Date:    "29/10/19",
Comment: "Premier contact",
}
var defaultLeadTags= []model.LeadTags{
	{
		TagID:      1,
		TagContent: "Piscine",
		TagIcon:    "mdi-phone",
	},
	{
		TagID:      2,
		TagContent: "Famille Nombreuse",
		TagIcon:    "mdi-phone",
	},
	{
		TagID:      3,
		TagContent: "Chauffage electrique",
		TagIcon:    "mdi-phone",
	},
}