package dao

import (
	"data-back-real/model"
	"database/sql"
	"fmt"
	"strconv"
)

// QueryTodoFromEmployee returns a to do list for a given employee id and category, which are leads or clients
func QueryTodoFromEmployee(db *sql.DB, conseillerID string, category string) ([]model.Todo, error) {
	var allTodo []model.Todo

	id, err := strconv.Atoi(conseillerID); if err != nil {
		return nil, err
	}

	// Query todos infos from advisor ID
	clientInfo, err := db.Query("SELECT todo_id, client_id, category, motif FROM todo WHERE conseiller_id= $1;", id); if err != nil {
		return nil, err
	}
	defer clientInfo.Close()

	var todo model.Todo
	for clientInfo.Next() {
		err := clientInfo.Scan(&todo.Id, &todo.ClientID, &todo.Category, &todo.Motif); if err != nil {
			return nil, err
		}
		// This if statement is meant to get the category Client or Lead or if specified "home" get all clients and leads
		if todo.Category == category {
			// Query client's info from client iD
			prepareQuery := fmt.Sprintf("SELECT name, phone FROM %v WHERE client_id = %v", category, todo.ClientID)
			clienOtherInfo, err := db.Query(prepareQuery); if err != nil {
				return nil, err
			}
			defer clienOtherInfo.Close()
			for clienOtherInfo.Next() {
				err := clienOtherInfo.Scan(&todo.Name, &todo.Phone); if err != nil {
					return nil, err
				}
			}
			allTodo = append(allTodo, todo)
		}
		if len(allTodo) > 10  {
			break
		}
	}

	// If there are no tags for the client in DB, return default todos
	if len(allTodo) == 0 {
		for _, t := range defaultTodo {
			allTodo = append(allTodo, t)
			return allTodo, nil
		}
	}
	return allTodo, nil
}

// QueryHomeTodoFromEmployee returns a to do list for a given employee id
func QueryHomeTodoFromEmployee(db *sql.DB, conseillerID string, category string) ([]model.Todo, error) {
	//var todo model.Todo
	//var allTodo []model.Todo
	//
	//id, err := strconv.Atoi(conseillerID)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//
	//// Query todos infos from advisor ID
	//clientInfo, err := db.Query("SELECT todo_id, client_id, category, motif FROM todo WHERE conseiller_id= $1;", id); if err != nil {
	//	fmt.Println(err.Error())
	//}
	//defer clientInfo.Close()
	//
	//for clientInfo.Next() {
	//	err := clientInfo.Scan(&todo.Id, &todo.ClientID, &todo.Category, &todo.Motif); if err != nil {
	//		fmt.Println("Error scanning results for todo", err.Error())
	//	}
	//	// This if statement is meant to get the category Client or Lead or if specified "home" get all clients and leads
	//	if todo.Category == category || todo.Category == "Home" {
	//		// Query client's info from client iD
	//		clienOtherInfo, err := db.Query("SELECT name, phone FROM $1 WHERE client_id= $2", category, todo.ClientID); if err != nil {
	//			fmt.Println(err.Error())
	//		}
	//		defer clienOtherInfo.Close()
	//		for clienOtherInfo.Next() {
	//			err := clienOtherInfo.Scan(&todo.Name, &todo.Phone); if err != nil {
	//				log.Fatal(err)
	//			}
	//			fmt.Println("client info", todo.Name)
	//
	//		}
	//		allTodo = append(allTodo, todo)
	//	}
	//	if len(allTodo) > 10  {
	//		break
	//	}
	//}
	//
	//// If there are no tags for the client in DB, return default todos
	//if len(allTodo) == 0 {
	//	for _, t := range defaultTodo {
	//		allTodo = append(allTodo, t)
	//		return allTodo
	//	}
	//}

	return defaultTodo, nil
}

// QueryHomeKPI returns important KPI for the employee's homepage
func QueryHomeKPI(db *sql.DB, conseillerID string) (model.HomeInfo, error) {
	var hi model.HomeInfo
	hi = defaultHomeInfo
	return hi, nil
}