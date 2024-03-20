package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"greatcomcatengineering.com/backend/database"
	"greatcomcatengineering.com/backend/middleware"
	"greatcomcatengineering.com/backend/models"
	"greatcomcatengineering.com/backend/utils"
)

// HandleCreateUser handles the creation of a new user. It performs the following steps:
//  1. Decodes the JSON request payload into a User struct.
//  2. Validates the user data using the ValidateUser function. If validation fails,
//     it responds with an HTTP 400 status code and an error message.
//  3. Hashes the user's password using bcrypt for secure storage.
//  4. Generates a unique UUID for the new user.
//  5. Attempts to add the new user to the database within a context with a 5-second timeout.
//     If this operation fails, it responds with an HTTP 500 status code and an error message.
//  6. Removes the password from the user object before sending a success response to ensure
//     sensitive information is not exposed.
//
// The function responds with the newly created user object, excluding the password, and an HTTP 201 status code.
//
// Params:
//
//	w http.ResponseWriter: The writer to send HTTP responses.
//	r *http.Request: The HTTP request object containing the user data.
func HandleCreateUser(c *gin.Context) {
	var req models.RegisterRequest

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if errMsg, valid := ValidateUserRegister(req); !valid {
		utils.RespondWithError(c, http.StatusBadRequest, errMsg)
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Error hashing password")
		return
	}

	user := models.User{
		ID:          uuid.New().String(),
		Email:       req.Email,
		Password:    string(hashed),
		AccountType: "default", // TODO: Default account type creation only for v0
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := database.AddUser(ctx, user); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utils.RespondWithJSON(c, http.StatusCreated, "User created successfully", user)
}

func HandleGetUserByEmail(c *gin.Context) {
	// Extract the ID from the URL path

	pathParts := strings.Split(c.Request.URL.Path, "/")
	var id string
	if len(pathParts) > 0 {
		id = pathParts[len(pathParts)-1] // Get the ID part
	}

	if id == "" {
		utils.RespondWithError(c, http.StatusBadRequest, "User ID is required")
		return
	}

	ctx := context.TODO()

	// Attempt to retrieve the user from the database
	user, err := database.GetUserByEmail(ctx, id)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to retrieve user: "+err.Error())
		return
	}

	// Successfully retrieved the user, respond with the user object
	utils.RespondWithJSON(c, http.StatusOK, "User retrieved successfully", user)
}

func HandleGetAllUsers(c *gin.Context) {

	ctx := context.TODO()

	// Attempt to retrieve all users from the database
	users, err := database.GetAllUsers(ctx)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to retrieve users: "+err.Error())
		return
	}

	// Successfully retrieved the users, respond with the users array
	utils.RespondWithJSON(c, http.StatusOK, "Users retrieved successfully", users)
}

func HandleLogin(c *gin.Context) {
	var loginRequest models.LoginRequest

	// Decode the request body into user struct
	if err := json.NewDecoder(c.Request.Body).Decode(&loginRequest); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	ctx := context.TODO()

	// Attempt to retrieve the user from the database
	dbUser, err := database.GetUserByEmail(ctx, loginRequest.Email)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to retrieve user: "+err.Error())
		return
	}

	fmt.Println("Encrypted: " + dbUser.Password)
	fmt.Println("Plain: " + loginRequest.Password)

	// Compare the stored hashed password with the password provided in the request
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginRequest.Password)); err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	token, err := middleware.GenerateJWTToken(dbUser.ID, dbUser.Email, dbUser.AccountType)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to generate token: "+err.Error())
		return
	}

	// Respond with the token
	utils.RespondWithJSON(c, http.StatusOK, "Login successful", map[string]string{"token": token})
}

func HandleGetCurrentUser(c *gin.Context) {
	// Extract the user ID from the request context
	val := c.Value("userID")
	userID, ok := val.(string)
	if !ok {
		utils.RespondWithError(c, http.StatusUnauthorized, "Unauthorized or invalid user ID")
		return
	}

	ctx := context.TODO()

	// Attempt to retrieve the user from the database
	user, err := database.GetUserByEmail(ctx, userID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to retrieve user: "+err.Error())
		return
	}

	// Successfully retrieved the user, respond with the user object
	utils.RespondWithJSON(c, http.StatusOK, "User retrieved successfully", user)
}
