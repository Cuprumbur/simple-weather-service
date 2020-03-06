package http

import (
	"strconv"

	"github.com/Cuprumbur/weather-service/detector"
	"github.com/Cuprumbur/weather-service/model"
	"github.com/labstack/echo/v4"
)

type detectorHandler struct {
	UseCase detector.UseCase
}

func SetupDetectorHandler(e *echo.Echo, u detector.UseCase) {
	handler := &detectorHandler{u}
	g := e.Group("/api/v1")
	g.POST("/detectors", handler.CreateDetector)
	g.GET("/detectors", handler.GetAllDetectors)
	g.GET("/detectors/:id", handler.GetDetector)
	g.POST("/detectors/:id", handler.UpdateDetector)
	g.DELETE("/detectors/:id", handler.DeleteDetector)

}

// GetAllDetectors godoc
// @Summary List detectors
// @Description get all detectors
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Detector
// @Router /detectors [get]
func (h *detectorHandler) GetAllDetectors(c echo.Context) error {

	detectors, err := h.UseCase.FindAllDetectors()

	if err != nil {
		return c.NoContent(500)
	}

	if detectors == nil {
		return c.NoContent(404)
	}

	return c.JSON(200, detectors)
}

// GetDetector godoc
// @Summary get detector by id
// @Description get detector by id
// @Accept  json
// @Produce  json
// @Param   id     path    int     true "Id of detector"
// @Success 200 {object} model.Detector
// @Router /detectors/{id} [get]
func (h *detectorHandler) GetDetector(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(400)
	}

	d, err := h.UseCase.FindDetector(id)
	if err != nil {
		return c.NoContent(500)
	}

	if d == nil {
		return c.NoContent(404)
	}

	return c.JSON(200, d)
}

// CreateDetector godoc
// @Summary create detector
// @Description  create detector
// @Accept  json
// @Produce  json
// @Param   detector     body    model.Detector     true "Detector"
// @Success 200 {object} model.Detector
// @Error 400 {object} httputil.HTTPError
// @Router /detectors [post]
func (h *detectorHandler) CreateDetector(c echo.Context) error {

	var d model.Detector
	err := c.Bind(&d)
	if err != nil {
		return c.NoContent(400)
	}

	if _, err := h.UseCase.Store(&d); err != nil {
		return c.NoContent(500)
	}

	return c.JSON(200, d)
}

// UpdateDetector godoc
// @Summary update detector
// @Description  update detector
// @Accept  json
// @Produce  json
// @Param   id     path    int     true "Id of detector"
// @Param   detector     body    model.Detector     true "Detector"
// @Router /detectors/{id} [post]
func (h *detectorHandler) UpdateDetector(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(400)
	}

	var d model.Detector
	err = c.Bind(&d)
	if err != nil {
		return c.NoContent(500)
	}
	d.ID = id
	err = h.UseCase.Update(&d)
	if err != nil {
		return c.NoContent(500)
	}

	return c.NoContent(200)
}

// DeleteDetector godoc
// @Summary delete detector
// @Description  delete detector
// @Accept  json
// @Produce  json
// @Param   id     path    int     true "Id of detector"
// @Router /detectors/{id} [delete]
func (h *detectorHandler) DeleteDetector(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(400)
	}

	err = h.UseCase.Delete(id)
	if err != nil {
		return c.NoContent(500)
	}

	return c.NoContent(200)
}
