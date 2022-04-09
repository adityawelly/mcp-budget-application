package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

// Get returns one income and error, if any
func (m *DBModel) Get(id int) (*Income, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, tanggal_in, jumlah_in, catatan_in, created_date, updated_date from tblm_income where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var income Income

	err := row.Scan(
		&income.ID,
		&income.TanggalIn,
		&income.JumlahIn,
		&income.CatatanIn,
		&income.CreatedDate,
		&income.UpdatedDate,
	)
	if err != nil {
		return nil, err
	}

	return &income, nil
}

func (m *DBModel) All() ([]*Income, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, tanggal_in, jumlah_in, catatan_in, created_date, updated_date from tblm_income order by tanggal_in desc`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var income []*Income

	for rows.Next() {
		var income Income
		err := rows.Scan(
			&income.ID,
			&income.TanggalIn,
			&income.JumlahIn,
			&income.CatatanIn,
			&income.CreatedDate,
			&income.UpdatedDate,
		)
		if err != nil {
			return nil, err
		}
	}
	return income, nil
}
