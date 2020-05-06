package service

import (
	"data-back-real/model"
)

// Receive infos:
	// id and prod
	// Niveau d’ensoleillement de la région (1, 2, 3 ou 4)
	// Chaque panneau solaire du foyer
	// Chaque équipement lourd
	// Chaque habitant du foyer
// convert all info the weighted criteras
// calculate auto conso ratio
// calcul score


type clientScoringCriteras struct {
	SunLevelWeight	int
	PanelsWeight	int
	HeavyEquipmentWeight	int
	FamilyMembersWeight	int
}

// ClientScoreCalculator takes all infos we have for a client and calculates a home energy score
func ClientScoreCalculator(cti model.ClientTotalInfo) (float64, float64) {

	scoringCriteras := fromDBInfosToWeightedCriteria(cti)

	autoConsoRatio := ratioAutoConsoCalculator(cti)

	score := 100 - scoringCriteras.SunLevelWeight - scoringCriteras.PanelsWeight  + scoringCriteras.HeavyEquipmentWeight + scoringCriteras.FamilyMembersWeight

	totalScore := float64(score) * autoConsoRatio

	return totalScore, autoConsoRatio
}

// ratioAutoConsoCalculator calculates the auto consumption ratio
func ratioAutoConsoCalculator(details model.ClientTotalInfo) float64 {

	totalHomeComsuption := details.FromGenToConsumer + details.FromGridToConsumer
	totalPannelProduction := details.FromGenToConsumer + details.FromGenToGrid

	ratio := totalPannelProduction / totalHomeComsuption * 100

	return ratio
}

// fromDBInfosToWeightedCriteria convert data from the database to weighted score criteria to be calculated in the score
func fromDBInfosToWeightedCriteria(cti model.ClientTotalInfo) clientScoringCriteras {

	var criteras = clientScoringCriteras{
		SunLevelWeight:       clientWeightCalculator(cti.SunLevel),
		PanelsWeight:         clientWeightCalculator(cti.Panels),
		HeavyEquipmentWeight: clientWeightCalculator(cti.HeavyEquipment),
		FamilyMembersWeight:  clientWeightCalculator(cti.FamilyMembers),
	}

	return criteras
}

// clientWeightCalculator is an algorithm that gives a weight of 10 per units
func clientWeightCalculator(number int) int {
	var weight int
	if number == 0 {
		return 0
	}
	for i := 1;  i<=number; i++ {
		weight = weight + 10
	}
	return weight
}




















func SunLevelWeight(number int) int {
	var weight int
	if number == 0 {
		return 0
	}
	for i := 1;  i<=number; i++ {
		weight = weight + 10
	}
	return weight
}

func PanelsWeight(number int) int {
	var weight int
	if number == 0 {
		return 0
	}
	for i := 1;  i<=number; i++ {
		weight = weight + 10
	}
	return weight
}

func HeavyEquipmentWeight(number int) int {
	var weight int
	if number == 0 {
		return 0
	}
	for i := 1;  i<=number; i++ {
		weight = weight + 10
	}
	return weight
}

func FamilyMembersWeight(number int) int {
	var weight int
	if number == 0 {
		return 0
	}
	for i := 1;  i<=number; i++ {
		weight = weight + 10
	}
	return weight
}