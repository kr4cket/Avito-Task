package handlers

import (
	"avitoTask/pkg/helpers"
	"avitoTask/pkg/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary AddUser
// @Description Создает пользователя в базе данных
// @ID add-user
// @Accept  json
// @Produce  json
// @Param input body models.User true "Информация о пользователе"
// @Success 200 {boolean} bool
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/users/add [post]
func (h *Handler) addUser(c *gin.Context) {

	var input models.User

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":        "Can't parse JSON",
			"errorMessage": err.Error(),
		})
		return
	}

	isAdded, err := h.services.AddUser(input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"userAdded":    isAdded,
			"errorMessage": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"userAdded": isAdded,
	})
}

type dateInput struct {
	Month int `json:"month" binding:"required"`
	Year  int `json:"year" binding:"required"`
}

// @Summary GetHistory
// @Description Создает ссылку на CSV-файл с историей добавлений/удалений пользователя в сегменты
// @ID get-history
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param input body dateInput true "Период за который необходимо получить информацию"
// @Success 200 {string} string "Link"
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/users/history/:id [post]
func (h *Handler) getHistory(c *gin.Context) {

	var uHelper helpers.UrlHelper
	var fHelper helpers.FileHelper

	var operations []models.Operation

	var input dateInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":        "Can't parse JSON",
			"errorMessage": err.Error(),
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":        "Id not found",
			"errorMessage": err.Error(),
		})
		return
	}

	operations, err = h.services.GetHistory(id, input.Month, input.Year)

	for index := range operations {

		if operations[index].Action == "false" {
			operations[index].Action = "added"
		} else {
			operations[index].Action = "deleted"
		}
	}

	name := uHelper.GenerateId()
	link := uHelper.GenerateUrl(c, name)

	c.JSON(http.StatusOK, map[string]interface{}{
		"downloadLink": link,
	})

	fHelper.CreateCSVFile(operations, name)

}
