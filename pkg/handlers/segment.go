package handlers

import (
	"avitoTask/pkg/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type getAllSegmentsResponse struct {
	Data []models.Segment `json:"data"`
}

// @Summary GetAllSegments
// @Description Возвращает все активные сегменты пользователя из Базы данных
// @ID get-all-segments
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/segments/all [get]
func (h *Handler) getAllSegments(c *gin.Context) {

	segments, err := h.services.GetAllSegments()

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":        "Can't get all segments",
			"errorMessage": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, getAllSegmentsResponse{
		Data: segments,
	})
}

// @Summary GetSegmentById
// @Description Возвращает сегмент по идентификатору
// @ID get-one-segment
// @Accept  json
// @Produce  json
// @Param id path int true "Segment ID"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/segments/get/:id [get]
func (h *Handler) getSegmentById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":        "Id not found",
			"errorMessage": err.Error(),
		})
		return
	}

	segment, err := h.services.GetSegmentById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":        "Segment not found",
			"errorMessage": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, segment)
}

// @Summary DeleteSegmentByName
// @Description Удаляет выбранный сегмент по уникальному имени
// @ID delete-segment
// @Accept  json
// @Produce  json
// @Param segmentName path string true "Segment Name"
// @Success 200 {boolean} bool
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/segments/delete/:segmentName [delete]
func (h *Handler) deleteSegmentByName(c *gin.Context) {

	segmentName := c.Param("segmentName")

	isDelete, err := h.services.Delete(segmentName)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"isDelete":     isDelete,
			"errorMessage": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"isDelete": isDelete,
	})
}

// @Summary CreateSegment
// @Description Создает сегмент в базе данных, при указании процента пользователей, которые должны находиться в сегменте, автоматически добавляет случайных пользователей в этот сегмент
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body models.Segment true "Информация о сегменте"
// @Success 200 {boolean} bool
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/segments/add [post]
func (h *Handler) createSegment(c *gin.Context) {

	var input models.Segment

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":        "Can't parse JSON",
			"errorMessage": err.Error(),
		})
		return
	}

	isInsert, err := h.services.Create(input)

	if input.Entirety.Int16 > 0 {
		users := h.services.GetRandomUsers(input.Entirety.Int16)
		for _, element := range users {
			h.services.ChangeSegments(element, []string{input.Name}, []string{})
			h.services.AddHistoryRecord(element, []string{input.Name}, []string{})
		}
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"isInsert":     isInsert,
			"errorMessage": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"isInsert": isInsert,
	})
}
