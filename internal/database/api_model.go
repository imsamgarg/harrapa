package database

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

// InscriptionModel represents the struct for client-facing API
type InscriptionModel struct {
	ID               string   `json:"id"`
	Site             string   `json:"site,omitempty"`
	CisiID           string   `json:"cisiID,omitempty"`
	SequenceImages   []string `json:"sequenceImages,omitempty"`
	SequenceNumbers  []int32  `json:"sequenceNumbers,omitempty"`
	WellsID          string   `json:"wellsID,omitempty"`
	ArtifactType     string   `json:"artifactType,omitempty"`
	MaterialType     string   `json:"materialType,omitempty"`
	FieldSymbol      string   `json:"fieldSymbol,omitempty"`
	ExcavationNumber int32    `json:"excavationNumber,omitempty"`
	Area             string   `json:"area,omitempty"`
}

// NewInscriptionModel creates an InscriptionModel from an Inscription
func NewInscriptionModel(i Inscription) InscriptionModel {
	return InscriptionModel{
		ID:               i.ID.String(),
		Site:             i.Site.String,
		CisiID:           i.CisiID.String,
		SequenceImages:   i.SequenceImages,
		SequenceNumbers:  i.SequenceNumbers,
		WellsID:          i.WellsID.String,
		ArtifactType:     i.ArtifactType.String,
		MaterialType:     i.MaterialType.String,
		FieldSymbol:      i.FieldSymbol.String,
		ExcavationNumber: i.ExcavationNumber.Int32,
		Area:             i.Area.String,
	}
}

// ToInscription converts InscriptionModel back to Inscription
func ToInscription(im InscriptionModel) (Inscription, error) {
	id, err := uuid.Parse(im.ID)
	if err != nil {
		return Inscription{}, fmt.Errorf("failed to parse ID: %w", err)
	}

	return Inscription{
		ID:               id,
		Site:             sql.NullString{String: im.Site, Valid: im.Site != ""},
		CisiID:           sql.NullString{String: im.CisiID, Valid: im.CisiID != ""},
		SequenceImages:   im.SequenceImages,
		SequenceNumbers:  im.SequenceNumbers,
		WellsID:          sql.NullString{String: im.WellsID, Valid: im.WellsID != ""},
		ArtifactType:     sql.NullString{String: im.ArtifactType, Valid: im.ArtifactType != ""},
		MaterialType:     sql.NullString{String: im.MaterialType, Valid: im.MaterialType != ""},
		FieldSymbol:      sql.NullString{String: im.FieldSymbol, Valid: im.FieldSymbol != ""},
		ExcavationNumber: sql.NullInt32{Int32: im.ExcavationNumber, Valid: im.ExcavationNumber != 0},
		Area:             sql.NullString{String: im.Area, Valid: im.Area != ""},
	}, nil
}

// NewInscriptionModelList converts a slice of Inscription to a slice of InscriptionModel
func NewInscriptionModelList(inscriptions []Inscription) []InscriptionModel {
	arr := make([]InscriptionModel, len(inscriptions))

	for _, i := range inscriptions {
		arr = append(arr, NewInscriptionModel(i))
	}

	return arr
}
