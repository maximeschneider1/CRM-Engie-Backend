package dao

import (
	"data-back-real/model"
	"data-back-real/service"
	"database/sql"
	"log"
	"math/rand"
	"strconv"
)

// QueryClientsFromConseillerID returns a list of clients for a given employee id
func QueryLeadsFromConseillerID(db *sql.DB, idString string) ([]model.Lead, error) {
	var ac []model.Lead

	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT client_id, name, phone, city, downloads, emails_opening, potential_gains, step FROM leads WHERE conseiller_id= $1", id); if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id model.Lead
		err := rows.Scan(&id.LeadID, &id.Name, &id.Phone, &id.City, &id.ContentDownloaded, &id.OpenedEmails, &id.Profitability, &id.Step); if err != nil {
			return nil, err
		}
		id.TimeSpent = rand.Intn(100 - 10) + 10
		id.WeeksSinceInactive = rand.Intn(30 - 1) + 1
		id = service.FromDBToWeightedCriteras(id)
		id.Score = service.ScoreCalculator(id)
		id.StepConverted = stepConverter(id.Step)

		ac = append(ac, id)
	}
	return ac, nil
}

// QueryLeadInfoFromLeadID returns detailed info of a lead for a given lead id
func QueryLeadInfoFromLeadID(db *sql.DB, leadID string) (model.Lead, error) {
	var cd model.Lead

	id, err := strconv.Atoi(leadID)
	if err != nil {
		return model.Lead{}, nil
	}

	rows, err := db.Query("SELECT client_id, name, email, phone, account_creation, city, downloads, emails_opening, potential_gains FROM leads WHERE client_id= $1", id); if err != nil {
		return model.Lead{}, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&cd.LeadID, &cd.Name, &cd.Email, &cd.Phone, &cd.FirstContact, &cd.City, &cd.ContentDownloaded, &cd.OpenedEmails, &cd.Profitability); if err != nil {
			log.Fatal(err)
		}
		cd.TimeSpent = rand.Intn(100 - 10) + 10
		cd.WeeksSinceInactive = rand.Intn(30 - 1) + 1
		cd = service.FromDBToWeightedCriteras(cd)
		cd.Score = service.ScoreCalculator(cd)
		cd.Address = "20 rue Rousseau"
		cd.StepConverted = stepConverter(cd.Step)
	}


	return cd, nil
}

// QueryLeadInfoFromLeadID returns contact history of a lead for a given lead id
func QueryLeadHistoryFromLeadID(db *sql.DB, leadID int) ([]model.LeadHistory, error) {
	var lh = []model.LeadHistory{}

	for i := 1;  i<=10; i++ {
		defaultLeadHistory.LeadID = leadID
		defaultLeadHistory.Comment = historyConverter(rand.Intn(4 - 1) + 1)
		defaultLeadHistory.Icon = iconConverter(rand.Intn(4 - 1) + 1)
		lh = append(lh, defaultLeadHistory)
		i++
	}

	return lh, nil
}

// QueryLeadTagsFromLeadID returns tags of a lead for a given lead id
func QueryLeadTagsFromLeadID(db *sql.DB, leadID int) ([]model.LeadTags, error) {
	var lt = []model.LeadTags{}

	for _, l := range defaultLeadTags {
		l.LeadID = leadID
		lt = append(lt, l)
	}

	return lt, nil
}

// stepConverter is used to convert the step int from the database to a string text
func stepConverter(s int) string {
	switch s {
	case 1 :
		return "Découverte"
	case 2 :
		return "Compatibilité"
	case 3 :
		return "Signature"
	case 4 :
		return "Installation"
	}
	return "Installation"
}

// historyConverter is used to simulate the lead contact history
func historyConverter(s int) string {
	switch s {
	case 1 :
		return "Mail"
	case 2 :
		return "Appel"
	case 3 :
		return "Renseignement"
	case 4 :
		return "Rendez-vous"
	}
	return "Rendez-vous"
}

// iconConverter is used to simulate the lead contact history icon
func iconConverter(s int) string {
	switch s {
	case 1 :
		return "mdi-phone"
	case 2 :
		return "mdi-email"
	case 3 :
		return "mdi-calendar-check"
	case 4 :
		return "mdi-solar-power"
	}
	return "mdi-email"
}
