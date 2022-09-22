package models

import (
	"github.com/jackc/pgtype"
)

type Film struct {
	Id         []uint8     `json:"-" db:"id"`
	Name       string      `json:"name" binding:"required"`
	Genre      string      `json:"genre" binding:"required"`
	DirectorId string      `json:"director_id" binding:"required"`
	Rate       float32     `json:"rate" binding:"required"`
	Year       pgtype.Date `json:"year" binding:"required"`
	Minutes    byte        `json:"minutes" binding:"required,gte=1"`
}
