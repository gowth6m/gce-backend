package user

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"greatcomcatengineering.com/backend/database"
	"greatcomcatengineering.com/backend/models"
	"greatcomcatengineering.com/backend/utils"
	"net/http"
	"strings"
	"time"
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
func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	// Decode the request body into newUser struct
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate newUser data (implement ValidateUser or similar function)
	if errMsg, valid := ValidateUser(newUser); !valid {
		utils.RespondWithError(w, http.StatusBadRequest, errMsg)
		return
	}

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error encrypting password")
		return
	}
	newUser.Password = string(hashedPassword)

	// Using context with timeout for database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Generate a UUID for the new user
	newUser.ID = uuid.New().String()

	// Attempt to add the new user to the database
	if err := database.AddUser(ctx, newUser); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create user: "+err.Error())
		return
	}

	responseUser := newUser
	responseUser.Password = ""

	utils.RespondWithJSON(w, http.StatusCreated, "User created successfully", responseUser)
}

func HandleGetUserByEmail(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL path

	pathParts := strings.Split(r.URL.Path, "/")
	var id string
	if len(pathParts) > 0 {
		id = pathParts[len(pathParts)-1] // Get the ID part
	}

	if id == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	ctx := context.TODO()

	// Attempt to retrieve the user from the database
	user, err := database.GetUserByEmail(ctx, id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve user: "+err.Error())
		return
	}

	// Successfully retrieved the user, respond with the user object
	utils.RespondWithJSON(w, http.StatusOK, "User retrieved successfully", user)
}

func HandleGetAllUsers(w http.ResponseWriter, r *http.Request) {

	ctx := context.TODO()

	// Attempt to retrieve all users from the database
	users, err := database.GetAllUsers(ctx)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve users: "+err.Error())
		return
	}

	// Successfully retrieved the users, respond with the users array
	utils.RespondWithJSON(w, http.StatusOK, "Users retrieved successfully", users)
}
