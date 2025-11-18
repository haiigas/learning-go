package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"learning/db"
	"learning/models"
	"learning/request"
	"learning/response"
	"learning/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	if err := h.DB.Preload("Biodata").Find(&users).Error; err != nil {
		utils.JSON(w, http.StatusInternalServerError, map[string]string{"error": "query error"})
		return
	}

	var responses []*response.UserResponse
	for _, u := range users {
		responses = append(responses, response.ToUserResponse(&u))
	}

	utils.JSON(w, http.StatusOK, utils.Response{
		Status:  true,
		Message: "fetch all users",
		Data:    responses,
	})
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.JSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	var req request.CreateUserRequest
	if err := utils.ParseJSON(r, &req); err != nil {
		utils.JSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, map[string]string{"error": "password hash error"})
		return
	}

	tx := db.DB.Begin()
	if tx.Error != nil {
		utils.JSON(w, http.StatusInternalServerError, map[string]string{"error": "transaction error"})
		return
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		utils.JSON(w, http.StatusInternalServerError, map[string]string{"error": "insert user error"})
		return
	}

	biodata := models.Biodata{
		UserID:  user.ID,
		Phone:   req.Phone,
		Address: req.Address,
	}

	if err := tx.Create(&biodata).Error; err != nil {
		tx.Rollback()
		utils.JSON(w, http.StatusInternalServerError, map[string]string{"error": "insert biodata error"})
		return
	}

	user.Biodata = &biodata

	if err := tx.Commit().Error; err != nil {
		utils.JSON(w, http.StatusInternalServerError, map[string]string{"error": "commit error"})
		return
	}

	utils.JSON(w, http.StatusCreated, utils.Response{
		Status:  true,
		Message: "user created successfully",
		Data:    response.ToCreateUserResponse(&user),
	})
}

func (h *UserHandler) GetUserDetail(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/users/"), "/")
	if len(parts) == 0 || parts[0] == "" {
		utils.JSON(w, http.StatusBadRequest, map[string]string{"error": "user id required"})
		return
	}

	id, err := strconv.Atoi(parts[0])
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, map[string]string{"error": "invalid user id"})
		return
	}

	var user models.User
	if err := h.DB.Preload("Biodata").First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JSON(w, http.StatusNotFound, map[string]string{"error": "user not found"})
			return
		}
		utils.JSON(w, http.StatusInternalServerError, map[string]string{"error": "query error"})
		return
	}

	utils.JSON(w, http.StatusOK, utils.Response{
		Status:  true,
		Message: "fetch user detail",
		Data:    response.ToUserResponse(&user),
	})
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		utils.JSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/users/"), "/")
	if len(parts) == 0 || parts[0] == "" {
		utils.JSON(w, http.StatusBadRequest, map[string]string{"error": "user id required"})
		return
	}

	id, err := strconv.Atoi(parts[0])
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, map[string]string{"error": "invalid user id"})
		return
	}

	var req request.UpdateUserRequest
	if err := utils.ParseJSON(r, &req); err != nil {
		utils.JSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	var user models.User
	if err := h.DB.Preload("Biodata").First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JSON(w, http.StatusNotFound, map[string]string{"error": "user not found"})
			return
		}
		utils.JSON(w, http.StatusInternalServerError, map[string]string{"error": "query error"})
		return
	}

	tx := db.DB.Begin()
	if tx.Error != nil {
		utils.JSON(w, http.StatusInternalServerError, map[string]string{"error": "transaction error"})
		return
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" && req.Email != user.Email {
		user.Email = req.Email
	}

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		utils.JSON(w, http.StatusInternalServerError, map[string]string{"error": "update user error"})
		return
	}

	if req.Phone != "" {
		user.Biodata.Phone = req.Phone
	}
	if req.Address != "" {
		user.Biodata.Address = req.Address
	}

	if err := tx.Save(&user.Biodata).Error; err != nil {
		tx.Rollback()
		utils.JSON(w, http.StatusInternalServerError, map[string]string{"error": "update biodata error"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		utils.JSON(w, http.StatusInternalServerError, map[string]string{"error": "commit error"})
		return
	}

	utils.JSON(w, http.StatusOK, utils.Response{
		Status:  true,
		Message: "user updated successfully",
		Data:    response.ToUserResponse(&user),
	})
}
