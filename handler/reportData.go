package handler

import (
	"go-web-server/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListReportData return response
func (r *Registry) ListReportData(c *gin.Context) {
	// start := c.DefaultQuery("start", "0")
	// end := c.DefaultQuery("end", "0")
	// scale := c.DefaultQuery("scale", "0")
	result := r.DAO.ListReportData()
	c.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": "Query Success",
	})
}

// CreateReportData return response
func (r *Registry) CreateReportData(c *gin.Context) {
	var data dao.ReportData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Data structure mismatched: " + err.Error(),
		})
		return
	}
	if err := r.DAO.CreateReportData(data.Timestamp, data.Type, data.Value, data.UUID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Report Success",
	})
}
