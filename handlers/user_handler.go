package handlers

import (
	"database/sql"
	"net/http"

	"learning/models"
	"learning/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	DB *sql.DB
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT 
			users.id, users.name, users.email,
			biodatas.phone, biodatas.address
		FROM users
		LEFT JOIN biodatas ON biodatas.user_id = users.id
	`
	rows, err := h.DB.Query(query)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, map[string]string{"error": "query error"})
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Phone, &u.Address); err != nil {
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, map[string]string{"error": "password hash error"})
		return
	}
	u.Password = string(hashedPassword)

	tx, err := h.DB.Begin()
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, map[string]string{"error": "transaction error"})
		return
	}

	res, err := h.DB.Exec(
		"INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
		u.Name, u.Email, u.Password,
	)
	if err != nil {
		tx.Rollback()
		utils.JSON(w, http.StatusInternalServerError, map[string]string{"error": "insert user error"})
		return
	}

	id, _ := res.LastInsertId()
	u.ID = int(id)

	_, err = tx.Exec(
		"INSERT INTO biodatas (user_id, phone, address) VALUES (?, ?, ?)",
		id, u.Phone, u.Address,
	)
	if err != nil {
		tx.Rollback()
		utils.JSON(w, http.StatusInternalServerError, map[string]string{"error": "insert biodata error"})
		return
	}

	if err := tx.Commit(); err != nil {
		utils.JSON(w, http.StatusInternalServerError, map[string]string{"error": "commit error"})
		return
	}

	u.Password = ""

	utils.JSON(w, http.StatusCreated, utils.Response{
		Status:  true,
		Message: "user created successfully",
		Data:    u,
	})
}
