package config

import (
	"encoding/json"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type dbConfig struct {
	Name     string `json:"dbname"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Database string `json:"database"`
	Port     int `json:"port"`
	Password string `json:"password"`
}

// ReturnDB reads json config file and returns an DB connection
func ReturnDB(configPath string) (*sql.DB, error) {

	dbc, err := readConfig(configPath); if err != nil {
		return nil, err
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbc.Host, dbc.Port, dbc.User, dbc.Password, dbc.Name)

	db, err := sql.Open("postgres", psqlInfo); if err != nil {
		return nil, err
	}

	err = db.Ping(); if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to DB")
	return db, nil
}

// readConfig reads json client secret to return db config
func readConfig(configPath string) (dbConfig, error) {
	var DBconfig = dbConfig{}

	jsonFile, err := os.Open(configPath); if err != nil {
		return dbConfig{}, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile); if err != nil {
		return dbConfig{}, err
	}
	json.Unmarshal(byteValue, &DBconfig)

	return DBconfig, nil
}