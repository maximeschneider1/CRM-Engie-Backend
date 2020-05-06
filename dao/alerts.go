package dao

import (
	"data-back-real/model"
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

func QueryClientsAlertsFromConseillerID(db *sql.DB, idString string) []model.AlertsDetails {

	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println(err.Error())
	}
	rows, err := db.Query("SELECT alert_id, client_id, conseiller_id, date, motif FROM alerts WHERE conseiller_id = $1", id); if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	var allResult []model.AlertsDetails

	var i = 0
	for rows.Next() {
		var alert model.AlertsDetails
		err := rows.Scan(&alert.AlertId, &alert.ClientId, &alert.ConseillerId, &alert.Date,  &alert.Motif); if err != nil {
			log.Fatal(err)
		}
		allResult = append(allResult, alert)

		i++
		if i == 10 {
			return allResult
		}
	}

	return allResult
}

func QueryClientsAlertsFromClientID(db *sql.DB, idString string) []model.AlertsDetails {

	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println(err.Error())
	}
	rows, err := db.Query("SELECT alert_id, client_id, conseiller_id, date, motif FROM alerts WHERE client_id = $1", id); if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	var allResult []model.AlertsDetails

	var i = 0
	for rows.Next() {
		var alert model.AlertsDetails
		err := rows.Scan(&alert.AlertId, &alert.ClientId, &alert.ConseillerId, &alert.Date,  &alert.Motif); if err != nil {
			log.Fatal(err)
		}
		allResult = append(allResult, alert)

		i++
		if i == 10 {
			return allResult
		}
	}

	return allResult
}