package handler

import (
	"context"
	"data-back-real/config"
	"data-back-real/service"
	"database/sql"
	"io/ioutil"
	"strings"

	//"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	//"io/ioutil"
	"log"
	"net/http"
)

var configPath = "/Users/max/go/src/data-back/config/config.json"

type server struct {
	db     *sql.DB
	router *httprouter.Router
}

func StartWebServer() {
	db, err := config.ReturnDB(configPath); if err != nil {
		fmt.Println(err.Error())
	}

	s := server{
		db: db,
		router: httprouter.New(),
	}

	//Launch predictive scripts
	go func() {
		service.LowConso(s.db)
	}()


	s.routes()


	log.Fatal(http.ListenAndServe(":8085", s.router))
}

func (s *server) routes() {
	s.router.GET("/clients_list/:conseiller_id", s.HandleGetClients())
	s.router.GET("/home/:conseiller_id", s.HandleHomeInfos())
	s.router.GET("/todo/:conseiller_id", s.HandleTodoInfos())
	s.router.GET("/alerts/:conseiller_id", s.HandleGetAdviserAlert())

	s.router.GET("/details/production/:client_id", s.HandleClientProductionDetails())
	s.router.GET("/details/info/:client_id", s.HandleClientPersonalDetails())
	s.router.GET("/details/tag/:client_id", s.HandleClientPersonalTag())
	s.router.GET("/clients/alerts/:client_id", s.HandleGetClientAlert())

	s.router.GET("/leads/history/:lead_id", s.HandleLeadHistory())
	s.router.GET("/leads/tags/:lead_id", s.HandleLeadTags())
	s.router.GET("/leads/info/:lead_id", s.HandleLeadInfos())

	s.router.POST("/client/new_tag/:conseiller_id", s.HandlePostTag())
}

func (s *server) HandlePostTag() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading request body, error :", err.Error())
			w.WriteHeader(500)
			return
		}
		// **La manipulation suivante existe parce que je ne suis arrivé à passer que un payload de type string et pas JSON avec le post de axios**
		a := strings.Split(string(body), "&")
		name := strings.Split(a[0], "=")
		clientID := strings.Split(a[1], "=")

		//Get last tag id
		var lastID int
		err = s.db.QueryRow("SELECT tag_id FROM tags ORDER BY tag_id DESC LIMIT 1;").Scan(&lastID)
		if err != nil {
			fmt.Println("Error querying last tag id, error :", err.Error())
			w.WriteHeader(500)
			return
		}

		// Post the alert to the DB
		var ctx = context.Background()
		tx, err := s.db.BeginTx(ctx, nil); if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(500)
		}
		_, err = tx.ExecContext(ctx, "INSERT INTO tags (tag_id, client_id, name) VALUES ($1, $2, $3)", lastID + 1, clientID[1], name[1])
		if err != nil {
			// In case we find any error in the query execution, rollback the transaction
			tx.Rollback()
			w.WriteHeader(500)
			return
		}
		err = tx.Commit()
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(200)
		w.Write([]byte("OK"))
	}
}
