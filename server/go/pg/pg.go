package pg

import (
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Users struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Email string `json:"email"`
}

func (Users) TableName() string {
    return "Users"
}

func OpenPostgres (mux *http.ServeMux) {
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			return
		}

		var user Users
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			return
		}

		result := db.Create(&Users{Email: user.Email})
		if result.Error != nil {
			return
		}
		w.WriteHeader(http.StatusCreated)
	})

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			return
		}

		var users []Users
		result := db.Find(&users)
		if result.Error != nil {
			return
		}

		jsonData, err := json.Marshal(users)
		if err != nil {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	})
}