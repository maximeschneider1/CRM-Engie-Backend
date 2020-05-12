package dao

import (
	"context"
	"data-back-real/model"
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

// QueryClientTagFromClientID returns personal tag for given client id
func QueryClientTagFromClientID(db *sql.DB, clientID string) ([]model.TagClient, error) {
	var allTag []model.TagClient

	id, err := strconv.Atoi(clientID)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT tag_id, name FROM tags WHERE client_id= $1;", id); if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	var tag model.TagClient
	for rows.Next() {
		err := rows.Scan(&tag.TagID, &tag.Name); if err != nil {
			return nil, err
		}
		allTag = append(allTag, tag)
	}

	// If there are no tags for the client in DB, return default tags
	if len(allTag)== 0 {
		for _, t := range defaultTags {
			tag.Name = t
			allTag = append(allTag, tag)
		}
	}

	return allTag, nil
}

// AddClientTagWithClientID inserts in the database a specific tag
func AddClientTagWithClientID(db *sql.DB, clientID string,  name string) error {
	//Get last tag id
	var lastID int
	err := db.QueryRow("SELECT tag_id FROM tags ORDER BY tag_id DESC LIMIT 1;").Scan(&lastID)
	if err != nil {
		fmt.Println("Error querying last tag id, error :", err.Error())
		return err
	}
	// Post tag in the DB
	var ctx = context.Background()
	tx, err := db.BeginTx(ctx, nil); if err != nil {
		log.Println("Error begining transaction :", err.Error())
		return err
	}
	_, err = tx.ExecContext(ctx, "INSERT INTO tags (tag_id, client_id, name) VALUES ($1, $2, $3)", lastID + 1, clientID[1], name[1]); if err != nil {
		// In case we find any error in the query execution, rollback the transaction
		log.Println("Error executing transaction :", err.Error())
		err = tx.Rollback(); if err != nil {
			log.Println("Error during rollback on transaction :", err.Error())
			return err
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Println("Error committing the transaction :", err.Error())
	}

	return nil
}
