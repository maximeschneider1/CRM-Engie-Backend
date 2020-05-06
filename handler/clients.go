package handler

import (
	"data-back-real/dao"
	"data-back-real/model"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (s *server) HandleGetClients() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, jwtToken, nextPageToken, prevPageToken")
		
		result := dao.QueryClientsFromConseillerID(s.db, ps.ByName("conseiller_id"))

		jsonBody, err := json.Marshal(&result); if err != nil {
			fmt.Println(err.Error())
		}

		w.Write(jsonBody)
	}
}

func (s *server) HandleHomeInfos() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, jwtToken, nextPageToken, prevPageToken")

		//result := dao.QueryClientsFromConseillerID(s.db, ps.ByName("conseiller_id"))

		result := model.HomeInfo{
			TotalClients:   34,
			Todo:           12,
			NewLeads:       70,
			NewDocuments:   3,
			PotentialValue: 113879,
		}


		jsonBody, err := json.Marshal(&result); if err != nil {
			fmt.Println(err.Error())
		}

		w.Write(jsonBody)
	}
}

func (s *server) HandleTodoInfos() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, jwtToken, nextPageToken, prevPageToken")

		//result := dao.QueryClientsFromConseillerID(s.db, ps.ByName("conseiller_id"))

		a := model.Todo{
			Name:       "Maxime",
			Id:         1,
			Telephone:  "0987097654",
			Category:   "Lead",
			Motif:      "Chute de production ",
		}
		b := model.Todo{
			Name:       "Michel",
			Id:         2,
			Telephone:  "7812397654",
			Category:   "Client",
			Motif:      "Chute de production",
		}
		c := model.Todo{
			Name:       "Antoine",
			Id:         3,
			Telephone:  "7812397654",
			Category:   "Client",
			Motif:      "Chute de production ",
		}
		d := model.Todo{
			Name:       "Etienne",
			Id:         4,
			Telephone:  "7812397654",
			Category:   "Client",
			Motif:      "Forte consommation ",
		}

		e := model.Todo{
			Name:       "Pierre",
			Id:         5,
			Telephone:  "7812397654",
			Category:   "Client",
			Motif:      "Chute de production ",
		}
		f := model.Todo{
			Name:       "Camille",
			Id:         6,
			Telephone:  "7812397654",
			Category:   "Client",
			Motif:      "Chute de production ",
		}
		result := []model.Todo{a, b, c, d, e, f}


		jsonBody, err := json.Marshal(&result); if err != nil {
			fmt.Println(err.Error())
		}

		w.Write(jsonBody)
	}
}

// HandleClientProductionDetails returns the production history for a given client id
func (s *server) HandleClientProductionDetails() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, jwtToken, nextPageToken, prevPageToken")

		result := dao.QueryProductionFromClientID(s.db, ps.ByName("client_id"))

		jsonBody, err := json.Marshal(&result); if err != nil {
			fmt.Println(err.Error())
		}

		w.Write(jsonBody)
	}
}

func (s *server) HandleClientPersonalDetails() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, jwtToken, nextPageToken, prevPageToken")

		result := dao.QueryClientInfoFromClientID(s.db, ps.ByName("client_id"))

		jsonBody, err := json.Marshal(&result); if err != nil {
			fmt.Println(err.Error())
		}

		w.Write(jsonBody)
	}
}

func (s *server) HandleClientPersonalTag() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, jwtToken, nextPageToken, prevPageToken")

		result := dao.QueryClientTagFromClientID(s.db, ps.ByName("client_id"))

		jsonBody, err := json.Marshal(&result); if err != nil {
			fmt.Println("Error marshalling databse tags to json response", err.Error())
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(200)
		w.Write(jsonBody)
	}
}