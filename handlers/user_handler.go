package handlers

import (
	"database/sql"
	"net/http"

	"learning/models"
	"learning/utils"
)

type UserHandler struct {
	DB *sql.DB
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query("SELECT id, name, phone, address FROM users")
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, map[string]string{"error": "query error"})
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Phone, &u.Address); err != nil {
			utils.JSON(w, http.StatusInternalServerError, map[string]string{"error": "scan error"})
			return
		}
		users = append(users, u)
	}

	utils.JSON(w, http.StatusOK, utils.Response{
		Status:  true,
		Message: "fetch all users",
		Data:    users,
	})
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.JSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	var u models.User
	if err := utils.ParseJSON(r, &u); err != nil {
		utils.JSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	result, err := h.DB.Exec(
		"INSERT INTO users (name, phone, address) VALUES (?, ?, ?)",
		u.Name, u.Phone, u.Address,
	)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, map[string]string{"error": "insert error"})
		return
	}

	id, _ := result.LastInsertId()
	u.ID = int(id)

	utils.JSON(w, http.StatusCreated, utils.Response{
		Status:  true,
		Message: "user created successfully",
		Data:    u,
	})
}
