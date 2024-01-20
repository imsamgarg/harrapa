package server

import (
	"encoding/json"
	"harrapa/internal/database"
	"harrapa/internal/utils"
	"log"
	"net/http"
)

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {

	var params struct {
		Email     string `json:"email" validator:"required,email"`
		Passsword string `json:"password" validator:"required"`
	}

	if err := s.JsonDecodeAndValidate(r.Body, &params); err != nil {
		utils.SendBadRequestResponse(w)
		return
	}

	user, err := s.db.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		utils.SendResponse(w, 400, utils.NewErrorResponse("Incorrect email or password"))
		return
	}

	isSamePassword, err := utils.CompareHashAndPassword(user.Password, params.Passsword)
	if err != nil || isSamePassword {
		utils.SendResponse(w, 400, utils.NewErrorResponse("Incorrect email or password"))
		return
	}

	jwtString, err := utils.GenerateJWT(user.ID.String(), database.UserRoleUser)
	if err != nil {
		utils.SendServerSideErrorResponse(w)
		return
	}

	utils.SendResponse(w, 200, map[string]string{
		"token": jwtString,
		"msg":   "Login successful!!",
	})

}

func (s *Server) RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var params struct {
		Email          string `json:"email" validator:"required,email"`
		Passsword      string `json:"password" validator:"required"`
		Name           string `json:"name" validator:"required"`
		ProfilePicture string `json:"profilePicture" validator:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		utils.SendBadRequestResponse(w)
		return
	}

	if err := validate.Struct(params); err != nil {
		utils.SendBadRequestResponse(w)
		return
	}

	userExists, err := s.db.DoesUserExists(r.Context(), params.Email)

	if err != nil {
		utils.SendServerSideErrorResponse(w)
		return
	}

	if userExists {
		utils.SendResponse(w, http.StatusBadRequest, utils.NewErrorResponse("User already exists"))
		return
	}

	hashedPass, err := utils.GenerateHashedPassword(params.Passsword)

	if err != nil {
		log.Println(err)
		utils.SendServerSideErrorResponse(w)
		return
	}

	id, err := s.db.CreateUser(r.Context(), database.CreateUserParams{
		Email:          params.Email,
		Name:           params.Name,
		Password:       hashedPass,
		ProfilePicture: utils.SqlMaybeNullString(params.ProfilePicture),
		Role:           database.UserRoleUser,
	})

	if err != nil {
		log.Println(err)
		utils.SendBadRequestResponse(w)
		return
	}

	jwtString, err := utils.GenerateJWT(id.String(), database.UserRoleUser)
	if err != nil {
		utils.SendServerSideErrorResponse(w)
		return
	}

	utils.SendResponse(w, 201, map[string]string{
		"token": jwtString,
		"msg":   "User registered successfully!",
	})

}
