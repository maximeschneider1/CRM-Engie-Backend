package service

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)
type lastProduction struct {
	alertID    int
	clientID   int
	employeeID int
	address    string
	latestProd []float64
	motif      string
}

// Limit variable is the loss of production % where the system thinks there is a problem with the panel production.
var limit = 90.0
var scheduling = 5 * time.Hour

// LowConsoDetection alerts loop on the database every scheduled time, searches for drastic loss of a
// client's solar panel production and logs an alert in the database if a loss is found
func LowConsoDetection(db *sql.DB)  {

	var ticker = time.NewTicker(scheduling)
	var allRecentProductions []lastProduction
	var problematicProduction []lastProduction

		for {
			select {
			case t := <-ticker.C:
				log.Println("Tick at", t)

				// Select all addresses
				addresses, err := db.Query("SELECT street FROM hd_copy WHERE  date BETWEEN SYMMETRIC '2019-11-20' AND '2019-12-01' GROUP BY street;"); if err != nil {
					log.Println("Error querying addresses for alert check, error :", err.Error())
				}
				err = addresses.Err(); if err != nil {
					log.Printf("Error querying the addresses, error : %v", err)
				}
				defer addresses.Close()
				var ad string
				var allAddresses []string
				for addresses.Next() {
					err := addresses.Scan(&ad); if err != nil {
						log.Println("Error Scanning selected row for addresses :", err.Error())
						return
					}
					allAddresses = append(allAddresses, ad)
				}

				// For all addresses, select the last production, we take a reference date of the 1st december
				// of 2019 as it is the time with the most clients values at the same time in the database
				for _, ad := range allAddresses {
					var recentProdAdress lastProduction
					lastProduction, err := db.Query("SELECT from_gen_to_consumer FROM hd_copy WHERE street = $1 AND  date BETWEEN SYMMETRIC '2019-11-20' AND '2019-12-01';", ad); if err != nil {
						log.Println("Error querying production info for selected address", err.Error())
					}
					err = lastProduction.Err(); if err != nil {
						log.Printf("Error querying the production for address : %v, error : %v", ad,  err)
					}
					defer lastProduction.Close()
					for lastProduction.Next() {
						var recentProd float64
						err := lastProduction.Scan(&recentProd); if err != nil {
							log.Println("Error Scanning selected row for address for production, error :", err.Error())
							return
						}
						recentProdAdress.address = ad
						recentProdAdress.latestProd = append(recentProdAdress.latestProd, recentProd)
					}
					allRecentProductions = append(allRecentProductions, recentProdAdress)
				}

				// For every production of every clients, check for drastic loss of solar panel production
				for _, a := range allRecentProductions {
					// Find the average production during on the last 10 days
					sum := 0.0

					for _, valuex := range a.latestProd {
						sum = sum + valuex
					}
					lastValue := a.latestProd[len(a.latestProd)-1]

					average := sum / float64(len(a.latestProd))
					if average == 0 {
						log.Println("Not enough data to establish a diagnostic on address :", a.address)
						continue
					}
					diff := percentageChange(average, lastValue)

					// If the difference of the last production is superior to the limit, system considers there is a problem
					if diff < limit {
						problematicProduction = append(problematicProduction, a)
					}
				}

				// Trouver le client à partir de l'adresse, puis le conseiller
				// Faire un post dans la table "alerte conso"
				for _, o := range problematicProduction {
					var lastID int
					log.Println("Production issues detected  :", o.address,  o.latestProd)

					// Mettre toutes les streets à la suite dans 1 seule query pour optimiser ?
					err := db.QueryRow("SELECT client_id FROM foyer WHERE street = $1", o.address).Scan(&o.clientID); if err != nil {
						log.Printf("Error querying client ID for address : %v for alert check :", o.address, err.Error())
					}

					// Mettre toutes les streets à la suite dans 1 seule query pour optimiser ?
					err = db.QueryRow("SELECT conseiller_id FROM clients WHERE client_id = $1", o.clientID).Scan(&o.employeeID); if err != nil {
						log.Printf("Error querying conseiller ID for address : %v for alert check :", o.address, err.Error())
					}

					txn, err := db.Begin(); if err != nil {
						log.Printf("Error begining transaction : %v '%'", err.Error())
					}

					txn.QueryRow("SELECT COUNT(*) FROM alerts;").Scan(&lastID)
					o.alertID = lastID + 1
					o.motif = fmt.Sprintf("Baisse de consommation supérieure à %v pourcent", limit)

					_, err = txn.Exec("INSERT INTO alerts (alert_id, client_id, conseiller_id, date, motif) VALUES ($1, $2, $3, $4, $5);", o.alertID, o.clientID, o.employeeID, time.Now(), o.motif)
					if err != nil {
						log.Println("Error inserting rows:", err)
					}

					txn.Commit()

					log.Printf("Success writing alert to DB for : %#v", o)
				}
			}
		}
}

func percentageChange(old float64, new float64) (delta float64) {
	diff := float64(new - old)
	delta = (diff / float64(old)) * 100
	return
}