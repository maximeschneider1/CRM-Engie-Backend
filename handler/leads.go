package handler

import (
	"data-back-real/dao"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

// handleLeadPersonalDetails returns a lead's general informations
func (s *server) handleLeadPersonalDetails() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, jwtToken, nextPageToken, prevPageToken")

		var resp response

		l, err := dao.QueryLeadInfoFromLeadID(s.db, ps.ByName("lead_id")); if err != nil {
			resp.Error = "Database error"
			resp.Message = "Internal Server Error"
			resp.StatusCode = http.StatusInternalServerError
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(resp); if err!= nil {
				log.Printf("Error encoding response : %v", err)
			}
			return
		}

		resp.Data = append(resp.Data, l)
		resp.StatusCode = http.StatusOK
		resp.Message = "OK"
		resp.Error = "No error"
		resp.Meta.Query = fmt.Sprintf("Informations du prospect : %v", ps.ByName("client_id"))
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp); if err!= nil {
			log.Printf("Error encoding response : %v", err)
		}
	}
}

// handleLeadHistory returns a history of the contacts between a lead and the company
func (s *server) handleLeadHistory() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, jwtToken, nextPageToken, prevPageToken")

		var resp response

		leadId,err  := strconv.Atoi(ps.ByName("lead_id")); if err != nil {
			log.Println("Error converting route parameter to int, error :", err.Error())
			resp.Error = "Database error"
			resp.Message = "Internal Server Error"
			resp.StatusCode = http.StatusInternalServerError
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(resp); if err!= nil {
				log.Printf("Error encoding response : %v", err)
			}
			return
		}

		lh, err := dao.QueryLeadHistoryFromLeadID(s.db, leadId); if err!= nil {
			resp.Error = "Database error"
			resp.Message = "Internal Server Error"
			resp.StatusCode = http.StatusInternalServerError
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(resp); if err!= nil {
				log.Printf("Error encoding response : %v", err)
			}
			return
		}

		resp.Data = append(resp.Data, lh)
		resp.StatusCode = http.StatusOK
		resp.Message = "OK"
		resp.Error = "No error"
		resp.Meta.Query = fmt.Sprintf("Historique de contact du prospect : %v", ps.ByName("client_id"))
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp); if err!= nil {
			log.Printf("Error encoding response : %v", err)
		}
	}
}

// handleLeadTags returns lead's personal tags used for customising the marketing pushes
func (s *server) handleLeadTags() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, jwtToken, nextPageToken, prevPageToken")

		var resp response

		leadId,err  := strconv.Atoi(ps.ByName("lead_id")); if err != nil {
			log.Println("Error converting route parameter to int, error :", err.Error())
			resp.Error = "Database error"
			resp.Message = "Internal Server Error"
			resp.StatusCode = http.StatusInternalServerError
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(resp); if err!= nil {
				log.Printf("Error encoding response : %v", err)
			}
			return
		}

		lh, err := dao.QueryLeadTagsFromLeadID(s.db, leadId); if err!= nil {
			resp.Error = "Database error"
			resp.Message = "Internal Server Error"
			resp.StatusCode = http.StatusInternalServerError
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(resp); if err!= nil {
				log.Printf("Error encoding response : %v", err)
			}
			return
		}

		resp.Data = append(resp.Data, lh)
		resp.StatusCode = http.StatusOK
		resp.Message = "OK"
		resp.Error = "No error"
		resp.Meta.Query = fmt.Sprintf("Tag du prospect : %v", ps.ByName("client_id"))
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp); if err!= nil {
			log.Printf("Error encoding response : %v", err)
		}
	}
}