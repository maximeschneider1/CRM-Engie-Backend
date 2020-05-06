package dao

import (
	"data-back-real/model"
	"data-back-real/service"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
	"strconv"
)

// Default variables are used if there are not enough info in DB to return coherent data to the user
var defaultScore = 150
var defaultAutoConsumption = 40
var defaultTags = []string{"Chauffage Gaz", "Famille Nombreuse", "Maison de vacance"}


func QueryProductionFromClientID(db *sql.DB, clientID string) model.ProductionStory {
	var pi model.ProductionQueryInfos
	var pd model.ProductionInfo
	//var pdList []model.ProductionInfo
	var totalPD model.ProductionInfo
	var productionHistory model.ProductionStory

	// Concert the id from route info from string to int
	id, err := strconv.Atoi(clientID)
	if err != nil {
		fmt.Println("Error reading client id", err.Error())
		return model.ProductionStory{}
	}

	// Query address for client id
	err = db.QueryRow("SELECT street FROM foyer WHERE client_id= $1;", id).Scan(&pi.ClientAddress)
	if err == sql.ErrNoRows {
		fmt.Printf("Error querying address for id %, error : %v", id, err.Error())
	}
	if err != nil {
		fmt.Printf("Error querying address for id %, error : %v", id, err.Error())
	}

	// Query production for address. Please note that this query take a reference data of the 22nd
	// of december, as it is the date with the most overall clients data. In a normal situation,
	// we would take the last result and limit the query to 8 returned rows
	prodInfos, err := db.Query("SELECT from_gen_to_consumer, from_gen_to_grid, from_grid_to_consumer FROM hd_copy WHERE date BETWEEN SYMMETRIC '2019-12-15' AND '2019-12-22'  AND street = $1;", pi.ClientAddress); if err != nil {
		fmt.Println("Error querying production info for selected address", err.Error())
	}
	err = prodInfos.Err()
	if err != nil {
		fmt.Printf("Error querying the production history for client %v, %v", clientID, err)
	}
	defer prodInfos.Close()

	// Scan all row and map data to pd struct
	for prodInfos.Next() {
		err := prodInfos.Scan(&pd.FromGenToConsumer, &pd.FromGenToGrid, &pd.FromGridToConsumer); if err != nil {
			fmt.Println("Error Scanning selected row for production info :", err.Error())
			return model.ProductionStory{}
		}
		// Client's production data are added to a slice so
		productionHistory.FromGridToConsumer = append(productionHistory.FromGridToConsumer, int(pd.FromGridToConsumer))
		productionHistory.FromGenToConsumer = append(productionHistory.FromGenToConsumer, int(pd.FromGenToConsumer))
		productionHistory.FromGenToGrid = append(productionHistory.FromGenToGrid, int(pd.FromGenToGrid))

		// Informations are also added to a totalPD to check if there are not enough data in database
		totalPD.FromGridToConsumer = totalPD.FromGridToConsumer + pd.FromGridToConsumer
		totalPD.FromGenToConsumer = totalPD.FromGenToConsumer + pd.FromGenToConsumer
		totalPD.FromGenToGrid = totalPD.FromGenToGrid + pd.FromGenToGrid
		//pdList = append(pdList, pd)
	}

	// If there are not enough data in DB for parameters, function returns mock data
	if totalPD.FromGridToConsumer == 0 {
		fmt.Printf("Not enough production data in database for client ID : %v, returning mock data", clientID)
		//pd = model.ProductionInfo{
		//	FromGenToConsumer:  2000,
		//	FromGenToGrid:      500,
		//	FromGridToConsumer: 1000,
		//}
		//pdList = append(pdList, pd)
		return model.ProductionStory{
			FromGenToConsumer:  []int{1950, 2050, 2400, 1900, 2200, 2000, 2050, 2090},
			FromGenToGrid:      []int{850, 750, 500, 1000, 400, 600, 950, 590},
			FromGridToConsumer: []int{950, 1050, 1400, 900, 1000, 1000, 1050, 1090},
			Score:              defaultScore,
			AutoConsumption:    defaultAutoConsumption,
		}
	}

	// Create a new client object to pass to ClientScoreCalculator
	newClient := model.ClientTotalInfo{
		ClientDetail:         model.ClientDetail{
			FamilyMembers:  3,
			Panels:         2,
			HeavyEquipment: 2,
			SunLevel:       2,
		},
		ProductionInfo: totalPD,
	}

	// Calculate the client auto consumption ratio and score
	score, autoConsumption := service.ClientScoreCalculator(newClient)
	productionHistory.Score = int(score)
	productionHistory.AutoConsumption = int(autoConsumption)
	productionHistory.TotalProduction = int(totalPD.FromGenToConsumer)

	return productionHistory
}


func QueryClientsFromConseillerID(db *sql.DB, idString string) []model.ClientDetail {

	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println(err.Error())
	}
	rows, err := db.Query("SELECT client_id, name, phone, city FROM clients_2 WHERE conseiller_id= $1", id); if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	var allResult []model.ClientDetail

	for rows.Next() {
		var id model.ClientDetail
		err := rows.Scan(&id.ClientId, &id.Name, &id.Phone, &id.Address); if err != nil {
			log.Fatal(err)
		}
		id.Score = rand.Intn(1000 - 100) + 100
		allResult = append(allResult, id)
	}
	return allResult
}



func QueryClientInfoFromClientID(db *sql.DB, clientID string) model.ClientDetail {
	var client model.ClientDetail

	id, err := strconv.Atoi(clientID)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Query address for client id
	err = db.QueryRow("SELECT street FROM foyer WHERE client_id= $1;", id).Scan(&client.Address)
	if err == sql.ErrNoRows {
		fmt.Printf("Error querying address for id %, error : %v", id, err.Error())
	}
	if err != nil {
		fmt.Printf("Error querying address for id %, error : %v", id, err.Error())
	}

	rows, err := db.Query("SELECT client_id, name, email, challenges_done, phone, account_creation, birthdate, registred_appliance, city, gender FROM clients_2 WHERE client_id= $1", id); if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&client.ClientId, &client.Name, &client.Email, &client.ChallengesDone, &client.Phone, &client.AccountCreation, &client.BirthDate, &client.RegistredDevices, &client.City, &client.Gender); if err != nil {
			log.Fatal(err)
		}
	}
	// Client's bill is calculated randomly
	client.LastBill = rand.Intn(1000 - 500) + 500
	return client
}

func QueryClientTagFromClientID(db *sql.DB, clientID string) []model.TagClient {
	var tag model.TagClient
	var allTag []model.TagClient

	id, err := strconv.Atoi(clientID)
	if err != nil {
		fmt.Println(err.Error())
	}

	rows, err := db.Query("SELECT tag_id, name FROM tags WHERE client_id= $1;", id); if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&tag.TagID, &tag.Name); if err != nil {
			log.Fatal(err)
		}
		allTag = append(allTag, tag)
	}

	// If there are no tags for the client in DB, return default tags
	if len(allTag)== 0 {
		for _, t := range defaultTags {
			tag.Name = t
			allTag = append(allTag, tag)
		}
	}

	return allTag
}

