package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	filePath string = "./historyfiles/"
)

// @Summary Download
// @Description Позволяет скачать сгенерированный CSV-файл с сервера
// @ID download-file
// @Accept  json
// @Produce  json
// @Param id path string true "File ID"
// @Success 200
// @Router /service/download/:id [get]
func (h *Handler) downloadFile(c *gin.Context) {
	fileName := c.Param("id")

	c.FileAttachment(fmt.Sprintf("%s/%s.csv", filePath, fileName), "history.csv")
}
