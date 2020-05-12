package dao

import (
	"data-back-real/model"
	"database/sql"
	"log"
	"strconv"
)

// QueryClientsAlertsFromConseillerID returns alerts for a given employee id
func QueryClientsAlertsFromConseillerID(db *sql.DB, idString string) ([]model.AlertsDetails, error) {
	var ad []model.AlertsDetails

	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT alert_id, client_id, conseiller_id, date, motif FROM alerts WHERE conseiller_id = $1", id); if err != nil {
		return nil, err
	}
	defer rows.Close()

	var i = 0
	for rows.Next() {
		var a model.AlertsDetails
		err := rows.Scan(&a.AlertId, &a.ClientId, &a.ConseillerId, &a.Date,  &a.Motif); if err != nil {
			return nil, err
		}
		ad = append(ad, a)
		i++
		if i == 10 {
			return ad, nil
		}
	}
	return ad, nil
}

// QueryClientsAlertsFromConseillerID returns alerts for a given client id
func QueryClientsAlertsFromClientID(db *sql.DB, idString string) ([]model.AlertsDetails, error) {
	var ad []model.AlertsDetails

	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT alert_id, client_id, conseiller_id, date, motif FROM alerts WHERE client_id = $1", id); if err != nil {
		return nil, err
	}
	defer rows.Close()

	var i = 0
	for rows.Next() {
		var alert model.AlertsDetails
		err := rows.Scan(&alert.AlertId, &alert.ClientId, &alert.ConseillerId, &alert.Date,  &alert.Motif); if err != nil {
			log.Fatal(err)
		}
		ad = append(ad, alert)
		i++
		if i == 10 {
			return ad, nil
		}
	}
	return ad, nil
}