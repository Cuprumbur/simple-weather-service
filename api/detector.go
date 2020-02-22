package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

)

func (a *API) GetAllDetectors(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ds, err := a.storage.FindAllDetectors()
	if err != nil {
		fmt.Printf("storage.FindAllDetectors %s \n %s\n", r.URL.Path, err.Error())
		write(w, 500, nil)
		return
	}

	if (ds == nil)	{
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

	if (d == nil)	{
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