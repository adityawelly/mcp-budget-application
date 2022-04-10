package main

import (
	"backend/cmd/api/models"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

type jsonResp struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// ===================================== INCOME =====================================

func (app *application) getOneIncome(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("Invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("id is", id)

	income, err := app.models.DB.GetIn(id)

	if err != nil {
		app.logger.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, income, "income")
	if err != nil {
		app.logger.Println(err)
	}

	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllIncome(w http.ResponseWriter, r *http.Request) {
	income, err := app.models.DB.AllIn()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, income, "income")
	if err != nil {
		app.errorJSON(w, err)
		return
	}

}

func (app *application) deleteIncome(w http.ResponseWriter, r *http.Request) {

}

func (app *application) insertIncome(w http.ResponseWriter, r *http.Request) {

}

type IncomePayload struct {
	ID        string `json:"id"`
	TanggalIn string `json:"tanggal_in"`
	JumlahIn  string `json:"jumlah_in"`
	CatatanIn string `json:"catatan_in"`
}

func (app *application) editIncome(w http.ResponseWriter, r *http.Request) {
	var payload IncomePayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err)
		return
	}

	var income models.Income

	income.ID, _ = strconv.Atoi(payload.ID)
	income.TanggalIn, _ = time.Parse("2006-01-02", payload.TanggalIn)
	income.JumlahIn, _ = strconv.Atoi(payload.JumlahIn)
	income.CatatanIn = payload.CatatanIn
	income.CreatedDate = time.Now()
	income.UpdatedDate = time.Now()

	err = app.models.DB.InsertIncome(income)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	ok := jsonResp{
		OK: true,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) searchIncome(w http.ResponseWriter, r *http.Request) {

}

// ===================================== INCOME =====================================

// ==================================== EXPENSES ====================================

func (app *application) getOneExpenses(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("Invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("id is", id)

	expenses, err := app.models.DB.GetEx(id)

	if err != nil {
		app.logger.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, expenses, "expenses")
	if err != nil {
		app.logger.Println(err)
	}

	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllExpenses(w http.ResponseWriter, r *http.Request) {
	expenses, err := app.models.DB.AllEx()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, expenses, "expenses")
	if err != nil {
		app.errorJSON(w, err)
		return
	}

}

func (app *application) deleteExpenses(w http.ResponseWriter, r *http.Request) {

}

func (app *application) insertExpenses(w http.ResponseWriter, r *http.Request) {

}

type ExpensesPayload struct {
	ID        string `json:"id"`
	TanggalEx string `json:"tanggal_ex"`
	JumlahEx  string `json:"jumlah_ex"`
	CatatanEx string `json:"catatan_ex"`
}

func (app *application) editExpenses(w http.ResponseWriter, r *http.Request) {
	var payload ExpensesPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err)
		return
	}

	var expenses models.Expenses

	expenses.ID, _ = strconv.Atoi(payload.ID)
	expenses.TanggalEx, _ = time.Parse("2006-01-02", payload.TanggalEx)
	expenses.JumlahEx, _ = strconv.Atoi(payload.JumlahEx)
	expenses.CatatanEx = payload.CatatanEx
	expenses.CreatedDate = time.Now()
	expenses.UpdatedDate = time.Now()

	err = app.models.DB.InsertExpenses(expenses)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	ok := jsonResp{
		OK: true,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) searchExpenses(w http.ResponseWriter, r *http.Request) {

}

// ==================================== EXPENSES ====================================
