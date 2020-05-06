package handler

import (
	"data-back-real/dao"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// HandleGetClientAlert returns adviser alert for the given adviser id
func (s *server) HandleGetAdviserAlert() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, jwtToken, nextPageToken, prevPageToken")

		result := dao.QueryClientsAlertsFromConseillerID(s.db, ps.ByName("conseiller_id"))

		jsonBody, err := json.Marshal(&result); if err != nil {
			fmt.Println(err.Error())
		}

		w.Write(jsonBody)
	}
}

// HandleGetClientAlert returns client's alert for the given client id
func (s *server) HandleGetClientAlert() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, jwtToken, nextPageToken, prevPageToken")

		result := dao.QueryClientsAlertsFromClientID(s.db, ps.ByName("client_id"))

		jsonBody, err := json.Marshal(&result); if err != nil {
			fmt.Println(err.Error())
		}

		w.Write(jsonBody)
	}
}