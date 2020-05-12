package handler

import (
	"data-back-real/dao"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// handleClientPersonalDetails returns a client's general infos
func (s *server) handleClientPersonalDetails() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		var resp response

		cd, err := dao.QueryClientInfoFromClientID(s.db, ps.ByName("client_id")); if err!= nil{
			log.Printf("Error querying client informations for client = %v, error : %v", ps.ByName("client_id"), err.Error())
			resp.Error = "Database error"
			resp.Message = "Internal Server Error"
			resp.StatusCode = http.StatusInternalServerError
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(resp); if err!= nil {
				log.Printf("Error encoding response : %v", err)
			}
			return
		}

		resp.Data = append(resp.Data, cd)
		resp.StatusCode = http.StatusOK
		resp.Message = "OK"
		resp.Error = "No error"
		resp.Meta.Query = fmt.Sprintf("Informations personnel du clients : %v", ps.ByName("client_id"))
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp); if err!= nil {
			log.Printf("Error encoding response : %v", err)
		}
	}
}

// handleClientProductionDetails returns the production history for a given client id
func (s *server) handleClientProductionDetails() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, jwtToken, nextPageToken, prevPageToken")

		var resp response

		p, err := dao.QueryProductionFromClientID(s.db, ps.ByName("client_id")); if err != nil {
			log.Printf("Error querrying production for client id : %, error : %v", ps.ByName("client_id"), err)
			resp.Error = "Database error"
			resp.Message = "Internal Server Error"
			resp.StatusCode = http.StatusInternalServerError
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(resp); if err!= nil {
				log.Printf("Error encoding response : %v", err)
			}
			return
		}

		resp.Data = append(resp.Data, p)
		resp.StatusCode = http.StatusOK
		resp.Message = "OK"
		resp.Error = "No error"
		resp.Meta.Query = fmt.Sprintf("Informations personnel du clients : %v", ps.ByName("client_id"))
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp); if err!= nil {
			log.Printf("Error encoding response : %v", err)
		}
	}
}

// handleClientPersonalTag returns client's personal tags used for customising the marketing pushes
func (s *server) handleClientPersonalTag() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, jwtToken, nextPageToken, prevPageToken")

		var resp response

		t, err := dao.QueryClientTagFromClientID(s.db, ps.ByName("client_id")); if err != nil {
			log.Printf("Error querying tag for client id : %v, error : %v", ps.ByName("client_id"), err)
			resp.Error = "Database error"
			resp.Message = "Internal Server Error"
			resp.StatusCode = http.StatusInternalServerError
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(resp); if err!= nil {
				log.Printf("Error encoding response : %v", err)
			}
			return
		}

		resp.Data = append(resp.Data, t)
		resp.StatusCode = http.StatusOK
		resp.Message = "OK"
		resp.Error = "No error"
		resp.Meta.Query = fmt.Sprintf("Informations personnel du clients : %v", ps.ByName("client_id"))
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp); if err!= nil {
			log.Printf("Error encoding response : %v", err)
		}
	}
}

// handleClientsPostTag handles the posting of a new client personal tag to the database
func (s *server) handleClientsPostTag() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		var resp response

		body, err := ioutil.ReadAll(r.Body); if err != nil {
			log.Println("Error reading request body, error :", err.Error())
			resp.Error = "Error reading request body"
			resp.Message = "Internal Server Error"
			resp.StatusCode = http.StatusInternalServerError
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(resp); if err!= nil {
				log.Printf("Error encoding response : %v", err)
			}
			return
		}
		// Take request header value and map to value, we could also pass by a real request body
		a := strings.Split(string(body), "&")
		name := strings.Split(a[0], "=")
		clientID := strings.Split(a[1], "=")

		err = dao.AddClientTagWithClientID(s.db, clientID[1], name[1])

		resp.StatusCode = http.StatusOK
		resp.Message = "OK"
		resp.Error = "No error"
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp); if err!= nil {
			log.Printf("Error encoding response : %v", err)
		}
	}
}

// handleGetClientAlert returns client's alert for the given client id
func (s *server) handleGetClientAlert() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, jwtToken, nextPageToken, prevPageToken")

		var resp response

		ad, err := dao.QueryClientsAlertsFromClientID(s.db, ps.ByName("client_id")); if err != nil {
			log.Printf("Error querying alerts from the database for given client id : %v, error : %v ", ps.ByName("client_id"), err.Error())
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
		resp.Meta.Query = fmt.Sprintf("Informations personnel du clients : %v", ps.ByName("client_id"))
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp); if err!= nil {
			log.Printf("Error encoding response : %v", err)
		}
	}
}