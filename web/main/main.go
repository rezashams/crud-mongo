/**
 * Copyright (c) 2019 TRIALBLAZE PTY. LTD. All rights reserved.
 *
 * Created by Reza Shams (reza.shams@trialblaze.com).
 * User: reza
 * Date: 6/30/19
 * Time: 4:27 PM
 *
 * Description: main
 *
 */
package main

import (
	"../../db"
	"../../model"
	"html/template"
	"net/http"
)

type CarsPageData struct {
	Cars     []model.Car
}

func main() {
	tmpl := template.Must(template.ParseFiles("./web/forms.html"))
	tmplCars := template.Must(template.ParseFiles("./web/carsShow.html"))
	http.HandleFunc("/cars", func (w http.ResponseWriter, r *http.Request) {
		var cars CarsPageData
		cars.Cars = db.GetCars()
		data :=cars
		tmplCars.Execute(w, data)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		details := model.Car{
			Color:   r.FormValue("color"),
			Model: r.FormValue("model"),
		}

		// do something with details
		db.InsertCar(details)

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":8081", nil)
}