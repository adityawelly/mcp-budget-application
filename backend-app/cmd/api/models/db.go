package models

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

// ===================================== INCOME =====================================

func (m *DBModel) GetIn(id int) (*Income, error) {
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

func (m *DBModel) AllIn() ([]*Income, error) {
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

func (m *DBModel) InsertIncome(income Income) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into tblm_income (tanggal_in, jumlah_in, catatan_in, created_date, updated_date) values ($1, $2, $3, $4, $5)`

	_, err := m.DB.ExecContext(ctx, stmt,
		income.TanggalIn,
		income.JumlahIn,
		income.CatatanIn,
		income.CreatedDate,
		income.UpdatedDate,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// ===================================== INCOME =====================================

// ==================================== EXPENSES ====================================

func (m *DBModel) GetEx(id int) (*Expenses, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, tanggal_ex, jumlah_ex, catatan_ex, created_date, updated_date from tblm_expenses where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var expenses Expenses

	err := row.Scan(
		&expenses.ID,
		&expenses.TanggalEx,
		&expenses.JumlahEx,
		&expenses.CatatanEx,
		&expenses.CreatedDate,
		&expenses.UpdatedDate,
	)
	if err != nil {
		return nil, err
	}

	return &expenses, nil
}

func (m *DBModel) AllEx() ([]*Expenses, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, tanggal_ex, jumlah_ex, catatan_ex, created_date, updated_date from tblm_expenses order by tanggal_ex desc`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []*Expenses

	for rows.Next() {
		var expenses Expenses
		err := rows.Scan(
			&expenses.ID,
			&expenses.TanggalEx,
			&expenses.JumlahEx,
			&expenses.CatatanEx,
			&expenses.CreatedDate,
			&expenses.UpdatedDate,
		)
		if err != nil {
			return nil, err
		}
	}
	return expenses, nil
}

func (m *DBModel) InsertExpenses(expenses Expenses) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into tblm_expenses (tanggal_ex, jumlah_ex, catatan_ex, created_date, updated_date) values ($1, $2, $3, $4, $5)`

	_, err := m.DB.ExecContext(ctx, stmt,
		expenses.TanggalEx,
		expenses.JumlahEx,
		expenses.CatatanEx,
		expenses.CreatedDate,
		expenses.UpdatedDate,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// ==================================== EXPENSES ====================================
