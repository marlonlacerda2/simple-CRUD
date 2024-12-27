package service

import (
	"log"

	"github.com/marlonlacerda2/simple-CRUD/internal/db"
	"github.com/marlonlacerda2/simple-CRUD/internal/model"
)

func GetUsers() ([]model.User, error) {
	conn, err := db.ConnectToDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT id,first_name, last_name, email  FROM users")
	if err != nil {
		log.Fatalf("Error querying database: %v", err)
	}
	defer rows.Close()
	var users []model.User

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error scanning row: %v", err)
		return nil, err
	}
	return users, nil
}
