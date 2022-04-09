package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneIncome(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("Invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("id is", id)

	income, err := app.models.DB.Get(id)

	if err != nil {
		app.logger.Println(err)
	}

	// income := models.Income {
	// 	ID:          id,
	// 	TanggalIn:   time.Date(2022, 04, 9, 9, 0, 0, 0, time.Local),
	// 	JumlahIn:    20000,
	// 	CatatanIn:   "Jajan",
	// 	CreatedDate: time.Now(),
	// 	UpdatedDate: time.Now(),
	// }

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
	income, err := app.models.DB.All()
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

func (app *application) updateIncome(w http.ResponseWriter, r *http.Request) {

}

func (app *application) searchIncome(w http.ResponseWriter, r *http.Request) {

}
