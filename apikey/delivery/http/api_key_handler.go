package http

import (
	"strconv"
	"strings"

	"github.com/Cuprumbur/weather-service/apikey"
	"github.com/labstack/echo/v4"

)

type apikeyHandler struct {
	UseCase apikey.UseCase
}

func SetupApiKeyHandler(e *echo.Echo, u apikey.UseCase) {
	handler := &apikeyHandler{u}
	g := e.Group("/api/v1")
	g.POST("/apikeys", handler.CreateApiKey)
	g.GET("/apikeys", handler.GetAllApiKeys)
	g.GET("/apikeys/:detector_id", handler.GetAllApiKeysByDetector)
	g.POST("/apikeys/:id,scopes", handler.UpdateScopes)
	g.DELETE("/apikeys/:id", handler.DeleteApiKey)

}

// GetAllApiKeys godoc
// @Summary List api keys
// @Description get all api keys
// @Accept  json
// @Produce  json
// @Success 200 {array} model.ApiKey
// @Router /apikeys [get]
func (h *apikeyHandler) GetAllApiKeys(c echo.Context) error {

	apiKeys, err := h.UseCase.FindAllApiKeys()

	if err != nil {
		return c.NoContent(500)
	}

	if apiKeys == nil {
		return c.NoContent(404)
	}

	return c.JSON(200, apiKeys)
}

func (h *apikeyHandler) GetAllApiKeysByDetector(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("detector_id"))
	if err != nil {
		return c.NoContent(400)
	}

	apiKeys, err := h.UseCase.FindApiKeys(id)
	if err != nil {
		return c.NoContent(500)
	}

	if apiKeys == nil {
		return c.NoContent(404)
	}

	return c.JSON(200, apiKeys)
}

// CreateApiKey godoc
// @Summary Create api key
// @Description create api key
// @Accept  json
// @Produce  json
// @Param   detector_id     query    int     true "Id of detector"
// @Param   scopes     query    string     true "scopes"
// @Success 200 {object} string
// @Router /apikeys [post]
func (h *apikeyHandler) CreateApiKey(c echo.Context) error {

	detectorID, err := strconv.Atoi(c.QueryParam("detector_id"))
	if err != nil {
		return c.NoContent(400)
	}
	scopes := c.QueryParam("scopes")
	if len(scopes) == 0 {
		return c.NoContent(400)
	}

	apiKey, err := h.UseCase.CreateApiKey(detectorID, strings.Split(scopes, ","))
	if err != nil {
		return c.NoContent(500)
	}

	return c.String(200, apiKey)
}

func (h *apikeyHandler) UpdateScopes(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(400)
	}

	scopes := c.Param("scopes")
	if len(scopes) == 0 {
		return c.NoContent(400)
	}

	err = h.UseCase.UpdateScopes(id, strings.Split(scopes, ","))

	if err != nil {
		return err
	}

	return c.NoContent(200)
}

func (h *apikeyHandler) DeleteApiKey(c echo.Context) error {

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
