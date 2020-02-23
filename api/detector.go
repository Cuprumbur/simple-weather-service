package api

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Cuprumbur/weather-service/storage"
	"github.com/gin-gonic/gin"

)

// GetAllDetectors godoc
// @Summary List detectors
// @Description get all detectors
// @Accept  json
// @Produce  json
// @Success 200 {array} storage.Detector
// @Router /detectors [get]
func (a *API) GetAllDetectors(c *gin.Context) {
	ds, err := a.storage.FindAllDetectors()
	if err != nil {
		fmt.Printf("storage.FindAllDetectors %s \n %s\n", c.Request.URL.Path, err.Error())

		write(c.Writer, 500, nil)
		return
	}

	if ds == nil {
		fmt.Printf("storage.FindAllDetectors %s \n no found", c.Request.URL.Path)
		write(c.Writer, 404, nil)
		return
	}

	b, err := json.Marshal(ds)
	if err != nil {
		fmt.Printf("json.Marshal %s \n %s\n", c.Request.URL.Path, err.Error())
		write(c.Writer, 500, nil)
		return
	}

	write(c.Writer, 200, b)
}

// GetDetector godoc
// @Summary get detector by id
// @Description get detector by id
// @Accept  json
// @Produce  json
// @Param   id     path    int     true "Id of detector"
// @Success 200 {object} storage.Detector
// @Router /detectors/{id} [get]
func (a *API) GetDetector(c *gin.Context) {

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		fmt.Printf("strconv.Atoi %s \n %s\n", c.Request.URL.Path, err.Error())
		write(c.Writer, 500, nil)
		return
	}

	d, err := a.storage.FindDetector(id)
	if err != nil {
		fmt.Printf("storage.FindDetector %s \n %s\n", c.Request.URL.Path, err.Error())
		write(c.Writer, 500, nil)
		return
	}

	if d == nil {
		fmt.Printf("storage.FindDetector %s \n no found", c.Request.URL.Path)
		write(c.Writer, 404, nil)
		return
	}

	b, err := json.Marshal(d)
	if err != nil {
		fmt.Printf("json.Marshal %s \n %s\n", c.Request.URL.Path, err.Error())
		write(c.Writer, 500, nil)
		return
	}

	write(c.Writer, 200, b)
}

// CreateDetector godoc
// @Summary create detector
// @Description  create detector
// @Accept  json
// @Produce  json
// @Param   detector     body    storage.Detector     true "Detector"
// @Router /detectors [post]
func (a *API) CreateDetector(c *gin.Context) {

	var d storage.Detector
	err := c.BindJSON(&d)
	if err != nil {
		fmt.Printf("BindJSON %s \n %s\n", c.Request.URL.Path, err.Error())
		write(c.Writer, 500, nil)
		return
	}
	id, err := a.storage.Store(&d)
	if err != nil {
		fmt.Printf("storage.Store %s \n %s\n", c.Request.URL.Path, err.Error())
		write(c.Writer, 500, nil)
		return
	}

	if id != 0 {
		fmt.Printf("CreateDetector success %d", id)
	}
}

// UpdateDetector godoc
// @Summary update detector
// @Description  update detector
// @Accept  json
// @Produce  json
// @Param   id     path    int     true "Id of detector"
// @Param   detector     body    storage.Detector     true "Detector"
// @Router /detectors/{id} [post]
func (a *API) UpdateDetector(c *gin.Context) {

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		fmt.Printf("strconv.Atoi %s \n %s\n", c.Request.URL.Path, err.Error())
		write(c.Writer, 500, nil)
		return
	}

	var d storage.Detector
	err = c.BindJSON(&d)
	if err != nil {
		fmt.Printf("BindJSON %s \n %s\n", c.Request.URL.Path, err.Error())
		write(c.Writer, 500, nil)
		return
	}
	d.ID = id
	err = a.storage.Update(&d)
	if err != nil {
		fmt.Printf("storage.FindDetector %s \n %s\n", c.Request.URL.Path, err.Error())
		write(c.Writer, 500, nil)
		return
	}
}

// DeleteDetector godoc
// @Summary delete detector
// @Description  delete detector
// @Accept  json
// @Produce  json
// @Param   id     path    int     true "Id of detector"
// @Router /detectors/{id} [delete]
func (a *API) DeleteDetector(c *gin.Context) {

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		fmt.Printf("strconv.Atoi %s \n %s\n", c.Request.URL.Path, err.Error())
		write(c.Writer, 500, nil)
		return
	}

	err = a.storage.Delete(id)
	if err != nil {
		fmt.Printf("storage.Delete %s \n %s\n", c.Request.URL.Path, err.Error())
		write(c.Writer, 500, nil)
		return
	}
}
