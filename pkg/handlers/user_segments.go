package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type changeSegment struct {
	SegmentsToAdd    []string `json:"add" binding: "required"`
	SegmentsToDelete []string `json:"delete" binding: "required"`
}

type expiredSegment struct {
	Segment    string `json:"segment" binding: "required"`
	TimeToLive int    `json:"time" binding: "required"`
}

// @Summary GetActiveSegments
// @Description Получает информацию об активных сегментах, в которых находится пользователь
// @ID get-active-segments
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param input body dateInput true "Период за который необходимо получить информацию"
// @Success 200 {array}  []string
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/users/active-segments/:id [get]
func (h *Handler) getActiveUserSegments(c *gin.Context) {

	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":        "Can't find User Id",
			"errorMessage": err.Error(),
		})
		return
	}

	expiredSegments, err := h.services.DeleteExpiredSegments(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":        "Can't delete expired segment records",
			"errorMessage": err.Error(),
		})
		return
	}

	h.services.AddHistoryRecord(userId, []string{}, expiredSegments)
	activeSegments, err := h.services.GetActiveUserSegments(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":        "Can't get segment records",
			"errorMessage": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"activeSegments": activeSegments,
	})

}

// @Summary AddExpiredSegment
// @Description Добавляет пользователя в сегмент на время
// @ID add-expired-segment
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param input body expiredSegment true "Принимает имя сегмента и время, через которое необходимо удалить пользователя из сегмента (в часах)"
// @Success 200 {boolean}  bool
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/users/expired-segments/:id [post]
func (h *Handler) addExpiredSegment(c *gin.Context) {

	var input expiredSegment

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":        "Can't parse JSON",
			"errorMessage": err.Error(),
		})
		return
	}

	fmt.Println(input)

	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":        "Can't find User Id",
			"errorMessage": err.Error(),
		})
		return
	}

	isSet, err := h.services.SetExpiredSegment(userId, input.TimeToLive, input.Segment)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":        "Error while changing segments",
			"errorMessage": err.Error(),
		})
		return
	}

	h.services.AddHistoryRecord(userId, []string{input.Segment}, []string{})

	c.JSON(http.StatusOK, map[string]interface{}{
		"isSegmentListChanged": isSet,
	})
}

// @Summary ChangeUserSegments
// @Description Позволяет добавить/удалить пользователя в сегмент (или несколько сегментов)
// @ID change-user-segments
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param input body changeSegment true "Принимает массивы сегментов для удаления и добавления"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/users/change-segments/:id [post]
func (h *Handler) changeUserSegments(c *gin.Context) {

	var input changeSegment

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":        "Can't parse JSON",
			"errorMessage": err.Error(),
		})
		return
	}

	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":        "Can't find User Id",
			"errorMessage": err.Error(),
		})
		return
	}

	addingSegments := getValidSegments(h, input.SegmentsToAdd)
	deletingSegments := getValidSegments(h, input.SegmentsToDelete)

	isChanged, err := h.services.ChangeSegments(userId, addingSegments, deletingSegments)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":        "Error while changing segments",
			"errorMessage": err.Error(),
		})
		return
	}

	h.services.AddHistoryRecord(userId, addingSegments, deletingSegments)

	c.JSON(http.StatusOK, map[string]interface{}{
		"isSegmentListChanged": isChanged,
		"added":                len(addingSegments),
		"removed":              len(deletingSegments),
	})

}

func getValidSegments(h *Handler, segmentNames []string) []string {
	var validSegments []string

	for _, segment := range segmentNames {
		valid, _ := h.services.CheckSegment(segment)
		if valid {
			validSegments = append(validSegments, segment)
		}
	}

	return validSegments
}
