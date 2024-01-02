package handler

import (
	"github.com/MrRytis/key-store/internal/storage"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Handler struct {
	Storage *storage.Storage
}

func NewHandler(storage *storage.Storage) *Handler {
	return &Handler{
		Storage: storage,
	}
}

type Error struct {
	Message string `json:"message"`
}

// GetAllValues godoc
// @Summary Get all values
// @Description Get all values from the store
// @Tags store
// @Accept json
// @Produce json
// @Success 200 {array} storage.Value
// @Router /api/v1/store [get]
func (h *Handler) GetAllValues(c echo.Context) error {
	return c.JSON(http.StatusOK, h.Storage.GetAllValues())
}

// GetValueByKey godoc
// @Summary Get value by key
// @Description Get value from the store by key
// @Tags store
// @Accept json
// @Produce json
// @Param key path string true "Key"
// @Success 200 {object} storage.Value
// @Failure 404 {object} Error
// @Router /api/v1/store/{key} [get]
func (h *Handler) GetValueByKey(c echo.Context) error {
	k := c.Param("key")

	v, err := h.Storage.GetValueByKey(k)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Value not found")
	}

	return c.JSON(http.StatusOK, v)
}

// StoreValue godoc
// @Summary Store value
// @Description Store value in the store
// @Tags store
// @Accept json
// @Produce json
// @Param value body storage.Value true "Value"
// @Success 201 {object} storage.Value
// @Failure 400 {object} Error
// @Router /api/v1/store [post]
func (h *Handler) StoreValue(c echo.Context) error {
	v := new(storage.Value)
	if err := c.Bind(v); err != nil {
		return err
	}

	if v.ExpireAt == 0 {
		v.ExpireAt = time.Now().Unix() + 300 // 5 minutes
	}

	if v.ExpireAt < time.Now().Unix() {
		return echo.NewHTTPError(http.StatusBadRequest, "ExpireAt must be in the future")
	}

	currentValue, err := h.Storage.GetValueByKey(v.Key)
	if err != nil {
		v = h.Storage.StoreValue(v.Key, v.Value, v.ExpireAt)

		return c.JSON(http.StatusCreated, v)
	}

	currentValue = h.Storage.UpdateValue(v.Key, v.Value, v.ExpireAt)

	return c.JSON(http.StatusOK, currentValue)
}

// DeleteValue godoc
// @Summary Delete value
// @Description Delete value from the store
// @Tags store
// @Accept json
// @Produce json
// @Param key path string true "Key"
// @Success 204
// @Router /api/v1/store/{key} [delete]
func (h *Handler) DeleteValue(c echo.Context) error {
	h.Storage.DeleteValue(c.Param("key"))

	return c.JSON(http.StatusNoContent, nil)
}
