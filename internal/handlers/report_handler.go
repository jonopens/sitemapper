package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"jonopens/sitemapper/internal/services"
)

// ReportHandler handles report-related HTTP requests
type ReportHandler struct {
	reportService *services.ReportService
}

// NewReportHandler creates a new report handler
func NewReportHandler(reportService *services.ReportService) *ReportHandler {
	return &ReportHandler{reportService: reportService}
}

// Get retrieves a specific report
// GET /api/v1/reports/:id
func (h *ReportHandler) Get(c *gin.Context) {
	id := c.Param("id")
	report, err := h.reportService.GetReport(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Report not found",
		})
		return
	}
	
	c.JSON(http.StatusOK, report)
}

// List returns all reports
// GET /api/v1/reports
func (h *ReportHandler) List(c *gin.Context) {
	// TODO: Get user ID from auth context
	userID := c.Query("user_id")
	
	reports, err := h.reportService.ListReportsByUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch reports",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"reports": reports,
	})
}

// Generate triggers report generation
// POST /api/v1/reports/:id/generate
func (h *ReportHandler) Generate(c *gin.Context) {
	// TODO: Implement report generation trigger
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Report generation started",
	})
}

