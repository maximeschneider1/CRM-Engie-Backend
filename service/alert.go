package service

import (
	"database/sql"
	"fmt"
	"time"
)
type lastProduction struct {
	alertID int
	clientID int
	conseilleriD int
	address string
	latestProd []float64
	motif string
}

// Limit variable is the loss of production % where the system thinks there is a problem with the panel production.
var limit = 90.0

func LowConso(db *sql.DB)  {

	ticker := time.NewTicker(5 * time.Hour)
	var clientAddress string
	var allAddress []string
	var allRecentProdAdress []lastProduction
	var problematicProduction []lastProduction

		for {
			select {
			case t := <-ticker.C:
				fmt.Println("Tick at", t)

				// chope toute les adress
				addresses, err := db.Query("SELECT street FROM hd_copy WHERE  date BETWEEN SYMMETRIC '2019-11-20' AND '2019-12-01' GROUP BY street;"); if err != nil {
					fmt.Println("Error querying addresses for alert check", err.Error())
				}
				err = addresses.Err(); if err != nil {
					fmt.Printf("Error querying the addresses %v", err)
				}
				defer addresses.Close()

				for addresses.Next() {
					err := addresses.Scan(&clientAddress); if err != nil {
						fmt.Println("Error Scanning selected row for addresses :", err.Error())
						return
					}
					allAddress = append(allAddress, clientAddress)
				}

				for _, ad := range allAddress {
					var recentProdAdress lastProduction

					lastProduction, err := db.Query("SELECT from_gen_to_consumer FROM hd_copy WHERE street = $1 AND  date BETWEEN SYMMETRIC '2019-11-20' AND '2019-12-01';", ad); if err != nil {
						fmt.Println("Error querying production info for selected address", err.Error())
					}
					err = lastProduction.Err(); if err != nil {
						fmt.Printf("Error querying the addresses %v", err)
					}
					defer lastProduction.Close()

					for lastProduction.Next() {
						var recentProd float64
						err := lastProduction.Scan(&recentProd); if err != nil {
							fmt.Println("Error Scanning selected row for address for production :", err.Error())
							return
						}
						recentProdAdress.address = ad
						recentProdAdress.latestProd = append(recentProdAdress.latestProd, recentProd)
					}
					allRecentProdAdress = append(allRecentProdAdress, recentProdAdress)
				}

				// Pour chaque address query la production des 10 derniers jours
				for _, a := range allRecentProdAdress {
					// Faire la moyenne de la production des 10 derniers jours
					sum := 0.0

					for _, valuex := range a.latestProd {
						sum = sum + valuex
					}
					lastValue := a.latestProd[len(a.latestProd)-1]

					average := sum / float64(len(a.latestProd))
					if average == 0 {
						fmt.Println("Not enough data to establish a diagnostic on address :", a.address)
						continue
					}
					diff := percentageChange(average, lastValue)

					// Comparer aujourd'hui à la moyenne des 10 derniers jours si la différence est de + de 90%
					if diff < - limit {
						problematicProduction = append(problematicProduction, a)
					}
				}

				// Trouver le client à partir de l'adresse, puis le conseiller
				// Faire un post dans la table "alerte conso"
				for _, o := range problematicProduction {

					var lastID int

					fmt.Println("Production issues detected  :", o.address,  o.latestProd)

					// Mettre toutes les streets à la suite dans 1 seule query pour optimiser ?
					err := db.QueryRow("SELECT client_id FROM foyer WHERE street = $1", o.address).Scan(&o.clientID); if err != nil {
						fmt.Printf("Error querying client ID for address : %v for alert check :", o.address, err.Error())
					}

					// Mettre toutes les streets à la suite dans 1 seule query pour optimiser ?
					err = db.QueryRow("SELECT conseiller_id FROM clients WHERE client_id = $1", o.clientID).Scan(&o.conseilleriD); if err != nil {
						fmt.Printf("Error querying conseiller ID for address : %v for alert check :", o.address, err.Error())
					}

					txn, err := db.Begin(); if err != nil {
						fmt.Printf("Error begining transaction : %v '%'", err.Error())
					}

					txn.QueryRow("SELECT COUNT(*) FROM alerts;").Scan(&lastID)
					o.alertID = lastID + 1
					o.motif = fmt.Sprintf("Baisse de consommation supérieure à %v pourcent", limit)

					_, err = txn.Exec("INSERT INTO alerts (alert_id, client_id, conseiller_id, date, motif) VALUES ($1, $2, $3, $4, $5);", o.alertID, o.clientID, o.conseilleriD, time.Now(), o.motif)
					if err != nil {
						fmt.Println("Error inserting rows:", err)
					}

					txn.Commit()

					fmt.Printf("Success writing alert to DB for : %#v", o)
				}
			}
		}
}

func percentageChange(old float64, new float64) (delta float64) {
	diff := float64(new - old)
	delta = (diff / float64(old)) * 100
	return
}