package service

import "data-back-real/model"

// get infos from db
// convert info to critera
// calculate the score
func GetScoreFromLeadID(id int)  {
	//leadInfo := sqlrequest(id)

	//lead := CriteraAdaptor(id)
	//
	//score := Lead.scoreCalculator()

}

// scoreCalculator gives back the score once lead info are available
func scoreCalculator(l model.Lead) int {
	var score int
	score = l.ContentDownloaded + l.TimeSpent + l.OpenedEmails + l.Profitability - l.WeeksSinceInactive
	return score
}

// fromDBToWeightedCriteras gives back a Lead struct filled with all necessary infos to calculate the score
func fromDBToWeightedCriteras(l model.Lead) model.Lead {

	l.ContentDownloaded = SiteContentWeightScore(l.ContentDownloaded)

	return l
}

// SiteContentWeightScore returns the weight of the amount of site contents user downloaded
func SiteContentWeightScore(number int) int  {
	var weight = 0

	if number == 0 {
		return 0
	}

	for i := 1;  i<=number; i++ {
		weight = weight + 5
	}

	return weight
}
