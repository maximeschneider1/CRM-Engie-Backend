package handler

import (
	"data-back-real/model"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (s *server) HandleLeadInfos() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, jwtToken, nextPageToken, prevPageToken")


		//result := dao.QueryProductionFromClientID(s.db, ps.ByName("client_id"))

		leadId,err  := strconv.Atoi(ps.ByName("lead_id")); if err != nil {
			fmt.Println(err.Error())
		}

		result := model.Lead{
			LeadID:             leadId,
			Name:             "Jean Maxime",
			Phone:            "0987896543",
			Address:          "22 rue de la Moulaga",
			Score:              112,
			FirstContact:       "28/09/18",
			City:               "Clermont-Ferrand",
			ContentDownloaded:  5,
			TimeSpent:          32,
			OpenedEmails:       4,
			Profitability:      543,
			WeeksSinceInactive: 1,
		}

		jsonBody, err := json.Marshal(&result); if err != nil {
			fmt.Println(err.Error())
		}

		w.Write(jsonBody)
	}
}



func (s *server) HandleLeadHistory() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, jwtToken, nextPageToken, prevPageToken")


		//result := dao.QueryProductionFromClientID(s.db, ps.ByName("client_id"))

		leadId,err  := strconv.Atoi(ps.ByName("lead_id")); if err != nil {
			fmt.Println(err.Error())
		}


		var result = []model.LeadHistory{}

		for i := 1;  i<=10; i++ {
			contact := model.LeadHistory{
				LeadID:  leadId,
				Type:    "Email",
				Icon:    "mdi-phone",
				Date:    "29/10/19",
				Comment: "Premier contact",
			}
			result = append(result, contact)

			i++
		}


		jsonBody, err := json.Marshal(&result); if err != nil {
			fmt.Println(err.Error())
		}

		w.Write(jsonBody)
	}
}

func (s *server) HandleLeadTags() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, jwtToken, nextPageToken, prevPageToken")


		//result := dao.QueryProductionFromClientID(s.db, ps.ByName("client_id"))

		leadId,err  := strconv.Atoi(ps.ByName("lead_id")); if err != nil {
			fmt.Println(err.Error())
		}


		var result = []model.LeadTags{}

		a := model.LeadTags{
			LeadID:     leadId,
			TagID:      1,
			TagContent: "Piscine",
			TagIcon:    "mdi-phone",
		}
		b := model.LeadTags{
			LeadID:     leadId,
			TagID:      2,
			TagContent: "Famille Nombreuse",
			TagIcon:    "mdi-phone",
		}
		//
		c  := model.LeadTags{
			LeadID:     leadId,
			TagID:      3,
			TagContent: "Chauffage electrique",
			TagIcon:    "mdi-phone",
		}

		result = append(result, a, b, c)


		jsonBody, err := json.Marshal(&result); if err != nil {
			fmt.Println(err.Error())
		}

		w.Write(jsonBody)
	}
}