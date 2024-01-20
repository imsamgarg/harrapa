package server

import (
	"database/sql"
	"encoding/json"
	"harrapa/internal/database"
	"harrapa/internal/utils"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (s *Server) AddInscriptionHandler(w http.ResponseWriter, r *http.Request) {

	var params database.InscriptionModel

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		utils.SendBadRequestResponse(w)
		return
	}

	_, err := s.db.AddInscription(r.Context(), database.AddInscriptionParams{
		Area:             utils.SqlMaybeNullString(params.Area),
		ArtifactType:     utils.SqlMaybeNullString(params.ArtifactType),
		CisiID:           utils.SqlMaybeNullString(params.CisiID),
		ExcavationNumber: sql.NullInt32{Int32: params.ExcavationNumber, Valid: true},
		FieldSymbol:      utils.SqlMaybeNullString(params.FieldSymbol),
		MaterialType:     utils.SqlMaybeNullString(params.MaterialType),
		SequenceImages:   params.SequenceImages,
		SequenceNumbers:  params.SequenceNumbers,
		Site:             utils.SqlMaybeNullString(params.Site),
		WellsID:          utils.SqlMaybeNullString(params.WellsID),
	})

	if err != nil {
		log.Println(err)
		//TODO(sam): handle sql errors
		utils.SendBadRequestResponse(w)
		return
	}

	utils.SendResponse(w, 200, utils.NewSuccessResponse("Inscription added successfully!"))
}

func (s *Server) GetInscritionById(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		utils.SendResponse(w, 400, utils.NewErrorResponse("Invalid id"))
		return
	}

	inscription, err := s.db.GetInsciptionById(r.Context(), uuid)

	if err != nil {
		if err == sql.ErrNoRows {
			utils.SendResponse(w, 404, utils.NewErrorResponse("Not found"))
		} else {
			utils.SendServerSideErrorResponse(w)
		}

		return
	}

	inscriptionModel := database.NewInscriptionModel(inscription)

	utils.SendResponse(w, 200, inscriptionModel)
}
func (s *Server) DeleteInscrition(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		utils.SendResponse(w, 400, utils.NewErrorResponse("Invalid id"))
		return
	}

	err = s.db.DeleteInsciption(r.Context(), uuid)

	if err != nil {
		if err == sql.ErrNoRows {
			utils.SendResponse(w, 404, utils.NewErrorResponse("Not Found"))
		} else {
			utils.SendServerSideErrorResponse(w)
		}

		return
	}

	utils.SendResponse(w, 200, utils.NewSuccessResponse("Inscription deleted successfully!"))
}

func (s *Server) GetAllInsciptions(w http.ResponseWriter, r *http.Request) {

	inscriptions, err := s.db.GetAllInsciptions(r.Context())

	if err != nil {
		if err == sql.ErrNoRows {
			utils.SendResponse(w, 200, []database.InscriptionModel{})
		} else {
			utils.SendServerSideErrorResponse(w)
		}

		return
	}

	inscriptionModel := database.NewInscriptionModelList(inscriptions)

	utils.SendResponse(w, 200, inscriptionModel)
}
