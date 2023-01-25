package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var users = []User{
	{ID: 1, Username: "alexymumo", Email: "alexymumo2000@gmail.com", Phone: "079345", Password: "*****", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All Users")
	json.NewEncoder(w).Encode(users)
}
