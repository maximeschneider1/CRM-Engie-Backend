package handler

import (
	"data-back-real/dao"
	"data-back-real/model"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// handleHomeInfos returns important KPI for the homepage
func (s *server) handleHomeInfos() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, jwtToken, nextPageToken, prevPageToken")
		var resp response

		hi, err := dao.QueryHomeKPI(s.db, ps.ByName("conseiller_id")); if err != nil {
			log.Printf("Error querying home informations for employee = %v, error : %v", ps.ByName("conseiller_id"), err.Error())
			resp.Error = "Database error"
			resp.Message = "Internal Server Error"
			resp.StatusCode = http.StatusInternalServerError
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(resp); if err!= nil {
				log.Printf("Error encoding response : %v", err)
			}
			return
		}

		resp.Data = append(resp.Data, hi)
		resp.StatusCode = http.StatusOK
		resp.Message = "OK"
		resp.Error = "No error"
		resp.Meta.Query = fmt.Sprintf("Informations de l'accueil de l'employé : %v", ps.ByName("client_id"))
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp); if err!= nil {
			log.Printf("Error encoding response : %v", err)
		}
	}
}

// handleTodoInfos returns a to do list. Weather it is only for the clients, leads or
// a global one for the homepage  is set in a post header "clients", "leads" or "home
func (s *server) handleTodoInfos() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		var resp response

		origin := r.Header.Get("Page-Origin")
		var result []model.Todo
		var err error

		switch origin {
		case "clients":
			result, err = dao.QueryTodoFromEmployee(s.db, ps.ByName("conseiller_id"), "clients"); if err != nil {
				log.Printf("Error querrying clients to do list for employee id : %v, error : %v", ps.ByName("conseiller_id"), err)
				resp.Error = "Database error"
				resp.Message = "Internal Server Error"
				resp.StatusCode = http.StatusInternalServerError
				w.WriteHeader(http.StatusInternalServerError)
				err = json.NewEncoder(w).Encode(resp); if err!= nil {
					log.Printf("Error encoding response : %v", err)
				}
			return
			}
		case "leads":
			result, err = dao.QueryTodoFromEmployee(s.db, ps.ByName("conseiller_id"), "leads"); if err != nil {
			log.Printf("Error querrying leads to do list for employee id : %v, error : %v", ps.ByName("conseiller_id"), err)
			resp.Error = "Database error"
			resp.Message = "Internal Server Error"
			resp.StatusCode = http.StatusInternalServerError
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(resp); if err!= nil {
				log.Printf("Error encoding response : %v", err)
			}
			return
			}
		case "home":
			result, err = dao.QueryHomeTodoFromEmployee(s.db, ps.ByName("conseiller_id"), "Home"); if err != nil {
			log.Printf("Error querrying home to do list for employee id : %v, error : %v", ps.ByName("conseiller_id"), err)
			resp.Error = "Database error"
			resp.Message = "Internal Server Error"
			resp.StatusCode = http.StatusInternalServerError
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(resp); if err!= nil {
				log.Printf("Error encoding response : %v", err)
			}
			return
			}
		}
		resp.Data = append(resp.Data, result)
		resp.StatusCode = http.StatusOK
		resp.Message = "OK"
		resp.Error = "No error"
		resp.Meta.Query = fmt.Sprintf("Tâches du conseiller : %v", ps.ByName("conseiller_id"))
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp); if err!= nil {
			log.Printf("Error encoding response : %v", err)
		}
	}
}

// handleGetLeads returns a client list for a given employee id
func (s *server) handleGetClients() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		var resp response

		ac, err := dao.QueryClientsFromConseillerID(s.db, ps.ByName("conseiller_id")); if err != nil {
			log.Printf("Error querrying clients for given employee ID :%v, error : %v",ps.ByName("conseiller_id"), err)
			resp.Error = "Database error"
			resp.Message = "Internal Server Error"
			resp.StatusCode = http.StatusInternalServerError
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(resp); if err!= nil {
				log.Printf("Error encoding response : %v", err)
			}
			return
		}
		resp.Data = append(resp.Data, ac)
		resp.StatusCode = http.StatusOK
		resp.Message = "OK"
		resp.Error = "No error"
		resp.Meta.Query = fmt.Sprintf("Listes de clients de l'employé : %v", ps.ByName("client_id"))
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp); if err!= nil {
			log.Printf("Error encoding response : %v", err)
		}
	}
}

// handleGetLeads returns a client list for a given employee id
func (s *server) handleGetLeads() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		var resp response

		ac, err := dao.QueryLeadsFromConseillerID(s.db, ps.ByName("conseiller_id")); if err != nil {
			log.Printf("Error querrying clients for given employee ID :%v, error : %v",ps.ByName("conseiller_id"), err)
			resp.Error = "Database error"
			resp.Message = "Internal Server Error"
			resp.StatusCode = http.StatusInternalServerError
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(resp); if err!= nil {
				log.Printf("Error encoding response : %v", err)
			}
			return
		}
		resp.Data = append(resp.Data, ac)
		resp.StatusCode = http.StatusOK
		resp.Message = "OK"
		resp.Error = "No error"
		resp.Meta.Query = fmt.Sprintf("Listes de clients de l'employé : %v", ps.ByName("client_id"))
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp); if err!= nil {
			log.Printf("Error encoding response : %v", err)
		}
	}
}


// handleGetEmployeeAlert returns adviser alert for the given employee id
func (s *server) handleGetEmployeeAlert() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		var resp response

		ad, err := dao.QueryClientsAlertsFromConseillerID(s.db, ps.ByName("conseiller_id")); if err != nil {
			log.Printf("Error querying alerts from the database for given employee id : %v, error : %v ", ps.ByName("conseiller_id"), err.Error())
			resp.Error = "Database error"
			resp.Message = "Internal Server Error"
			resp.StatusCode = http.StatusInternalServerError
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(resp); if err!= nil {
				log.Printf("Error encoding response : %v", err)
			}
			return
		}

		resp.Data = append(resp.Data, ad)
		resp.StatusCode = http.StatusOK
		resp.Message = "OK"
		resp.Error = "No error"
		resp.Meta.Query = fmt.Sprintf("Alertes clients du conseiller : %v", ps.ByName("conseiller_id"))
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp); if err!= nil {
			log.Printf("Error encoding response : %v", err)
		}
	}
}
