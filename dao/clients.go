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
// I would have declared these function as methods for *sql.DB type but it is declared in another package

// QueryProductionFromClientID return production info for a given client id
func QueryProductionFromClientID(db *sql.DB, clientID string) (model.ProductionStory, error) {
	var ps model.ProductionStory

	// Concert the id from route info from string to int
	id, err := strconv.Atoi(clientID)
	if err != nil {
		return model.ProductionStory{}, err
	}

	// Query address for client id
	var pi model.ProductionQueryInfos
	err = db.QueryRow("SELECT street FROM foyer WHERE client_id= $1;", id).Scan(&pi.ClientAddress)
	if err == sql.ErrNoRows {
		fmt.Printf("Error querying address for id %, error : %v", id, err.Error())
	}
	if err != nil {
		return model.ProductionStory{}, err
	}

	// Query production for address. Please note that this query take a reference data of the 22nd
	// of december, as it is the date with the most overall clients data. In a normal situation,
	// we would take the last result and limit the query to 8 returned rows
	prodInfos, err := db.Query("SELECT from_gen_to_consumer, from_gen_to_grid, from_grid_to_consumer FROM hd_copy WHERE date BETWEEN SYMMETRIC '2019-12-15' AND '2019-12-22'  AND street = $1;", pi.ClientAddress); if err != nil {
		fmt.Println("Error querying production info for selected address", err.Error())
	}
	if err != nil {
		return model.ProductionStory{}, err
	}
	defer prodInfos.Close()

	// Scan all row and map data to pd struct
	var pd model.ProductionInfo
	var totalPd model.ProductionInfo
	for prodInfos.Next() {
		err := prodInfos.Scan(&pd.FromGenToConsumer, &pd.FromGenToGrid, &pd.FromGridToConsumer); if err != nil {
			return model.ProductionStory{}, err
		}
		// Client's production data are added to a slice so
		ps.FromGridToConsumer = append(ps.FromGridToConsumer, int(pd.FromGridToConsumer))
		ps.FromGenToConsumer = append(ps.FromGenToConsumer, int(pd.FromGenToConsumer))
		ps.FromGenToGrid = append(ps.FromGenToGrid, int(pd.FromGenToGrid))

		// Informations are also added to a totalPd to check if there are not enough data in database
		totalPd.FromGridToConsumer = totalPd.FromGridToConsumer + pd.FromGridToConsumer
		totalPd.FromGenToConsumer = totalPd.FromGenToConsumer + pd.FromGenToConsumer
		totalPd.FromGenToGrid = totalPd.FromGenToGrid + pd.FromGenToGrid
	}

	// If there are not enough data in DB for parameters, function returns mock data
	if totalPd.FromGridToConsumer == 0 {
		log.Printf("Not enough production data in database for client ID : %v, returning mock data", clientID)
		return defaultProd, nil
	}

	// Create a new client object to pass to ClientScoreCalculator
	newClient := model.ClientTotalInfo{
		ClientDetail:         model.ClientDetail{
			FamilyMembers:  3,
			Panels:         2,
			HeavyEquipment: 2,
			SunLevel:       2,
		},
		ProductionInfo: totalPd,
	}

	// Calculate the client auto consumption ratio and score
	score, autoConsumption := service.ClientScoreCalculator(newClient)
	ps.Score = int(score)
	ps.AutoConsumption = int(autoConsumption)
	ps.TotalProduction = int(totalPd.FromGenToConsumer)

	return ps, nil
}

// QueryClientsFromConseillerID returns a list of clients for a given employee id
func QueryClientsFromConseillerID(db *sql.DB, idString string) ([]model.ClientDetail, error) {
	var ac []model.ClientDetail

	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT client_id, name, phone, city FROM clients WHERE conseiller_id= $1", id); if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id model.ClientDetail
		err := rows.Scan(&id.ClientId, &id.Name, &id.Phone, &id.Address); if err != nil {
			return nil, err
		}
		id.Score = rand.Intn(1000 - 100) + 100
		ac = append(ac, id)
	}
	return ac, nil
}

// QueryClientInfoFromClientID Returns detailed info of a client for a given client id
func QueryClientInfoFromClientID(db *sql.DB, clientID string) (model.ClientDetail, error) {
	var client model.ClientDetail

	id, err := strconv.Atoi(clientID)
	if err != nil {
		return model.ClientDetail{}, nil
	}
	// Query address for client id
	err = db.QueryRow("SELECT street FROM foyer WHERE client_id= $1;", id).Scan(&client.Address)
	if err == sql.ErrNoRows {
		return model.ClientDetail{}, err
	}; if err != nil {
		return model.ClientDetail{}, err
	}

	rows, err := db.Query("SELECT client_id, name, email, challenges_done, phone, account_creation, birthdate, registred_appliance, city, gender FROM clients WHERE client_id= $1", id); if err != nil {
		return model.ClientDetail{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&client.ClientId, &client.Name, &client.Email, &client.ChallengesDone, &client.Phone, &client.AccountCreation, &client.BirthDate, &client.RegistredDevices, &client.City, &client.Gender); if err != nil {
			log.Fatal(err)
		}
	}
	// Client's bill is calculated randomly
	client.LastBill = rand.Intn(1000 - 500) + 500

	return client, nil
}
