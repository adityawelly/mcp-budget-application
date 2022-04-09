package models

import (
	"database/sql"
	"time"
)

type Models struct {
	DB DBModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

type Income struct {
	ID          int       `json:"id"`
	TanggalIn   time.Time `json:"tanggal_in"`
	JumlahIn    int       `json:"jumlah_in"`
	CatatanIn   string    `json:"catatan_in"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}

type Expenses struct {
	ID          int       `json:"id"`
	TanggalEx   time.Time `json:"tanggal_ex"`
	JumlahEx    int       `json:"jumlah_ex"`
	CatatanEx   string    `json:"catatan_ex"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}
