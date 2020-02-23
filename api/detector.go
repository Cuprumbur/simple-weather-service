package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

)

// GetAllDetectors godoc
// @Summary List detectors
// @Description get all detectors
// @Accept  json
// @Produce  json
// @Success 200 {array} storage.Detector
// @Router /detectors [get]
func (a *API) GetAllDetectors(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ds, err := a.storage.FindAllDetectors()
	if err != nil {
		fmt.Printf("storage.FindAllDetectors %s \n %s\n", r.URL.Path, err.Error())
		write(w, 500, nil)
		return
	}

	if ds == nil {
		fmt.Printf("storage.FindAllDetectors %s \n no found", r.URL.Path)
		write(w, 404, nil)
		return
	}

	b, err := json.Marshal(ds)
	if err != nil {
		fmt.Printf("json.Marshal %s \n %s\n", r.URL.Path, err.Error())
		write(w, 500, nil)
		return
	}

	write(w, 200, b)
}

// GetDetector godoc
// @Summary get detector by id
// @Description get detector by id
// @Accept  json
// @Produce  json
// @Param id query string false "name search by id" Format(int)
// @Success 200 {object} storage.Detector
// @Router /detector/{id} [get]
func (a *API) GetDetector(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		fmt.Printf("strconv.Atoi %s \n %s\n", r.URL.Path, err.Error())
		write(w, 500, nil)
		return
	}

	d, err := a.storage.FindDetector(id)
	if err != nil {
		fmt.Printf("storage.FindDetector %s \n %s\n", r.URL.Path, err.Error())
		write(w, 500, nil)
		return
	}

	if d == nil {
		fmt.Printf("storage.FindDetector %s \n no found", r.URL.Path)
		write(w, 404, nil)
		return
	}

	b, err := json.Marshal(d)
	if err != nil {
		fmt.Printf("json.Marshal %s \n %s\n", r.URL.Path, err.Error())
		write(w, 500, nil)
		return
	}

	write(w, 200, b)
}