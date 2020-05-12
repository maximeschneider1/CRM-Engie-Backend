package handler

import (
	"data-back-real/config"
	"data-back-real/service"
	"database/sql"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

// Please configure the config path to the emplacement of the DB credential on your machine
var configPath = "/Users/max/go/src/data-back/config/config.json"

// Type server is the base structure of the API
type server struct {
	db     *sql.DB
	router *httprouter.Router
}

// response contains all response infos
type response struct {
	StatusCode int  `json:"status_code"`
	Error      string `json:"error"`
	Message    string `json:"message"`
	Meta       struct {
		Query       interface{} `json:"query,omitempty"`
		ResultCount int         `json:"result_count,omitempty"`} `json:"meta"`
	Data []interface{} `json:"data"`
}

// StartWebServer is the function responsible for launching the API
func StartWebServer() {
	db, err := config.ReturnDB(configPath); if err != nil {
		log.Printf("Error returning a connection to the datase : %v", err.Error())
	}
	s := server{
		db: db,
		router: httprouter.New(),
	}
	s.router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", r.Header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
			header.Set("Access-Control-Allow-Headers", "*")
		}
		w.WriteHeader(http.StatusNoContent)
	})
	s.router.PanicHandler = handlePanic
	s.routes()

	// Launch predictive scripts. This could be putted in a separated scheduler
	go func() {
		service.LowConsoDetection(s.db)
	}()

	log.Fatal(http.ListenAndServe(":8085", s.router))
}

// routes function launches all application's routes
func (s *server) routes() {
	// Employee related routes
	s.router.GET("/clients_list/:conseiller_id", s.handleGetClients())
	s.router.GET("/leads_list/:conseiller_id", s.handleGetLeads())
	s.router.GET("/home/:conseiller_id", s.handleHomeInfos())
	s.router.GET("/todo/:conseiller_id", s.handleTodoInfos())
	s.router.GET("/alerts/:conseiller_id", s.handleGetEmployeeAlert())

	// Client related routes
	s.router.GET("/details/production/:client_id", s.handleClientProductionDetails())
	s.router.GET("/details/info/:client_id", s.handleClientPersonalDetails())
	s.router.GET("/details/tag/:client_id", s.handleClientPersonalTag())
	s.router.GET("/clients/alerts/:client_id", s.handleGetClientAlert())
	s.router.POST("/client/new_tag/:conseiller_id", s.handleClientsPostTag())

	// Lead related routes
	s.router.GET("/leads/history/:lead_id", s.handleLeadHistory())
	s.router.GET("/leads/tags/:lead_id", s.handleLeadTags())
	s.router.GET("/leads/info/:lead_id", s.handleLeadPersonalDetails())
}

// Gracefully handle panic without crashing the server
func handlePanic(w http.ResponseWriter, r *http.Request, err interface{}) {
	log.Println(r.URL.Path, err)
	w.WriteHeader(http.StatusInternalServerError)
}