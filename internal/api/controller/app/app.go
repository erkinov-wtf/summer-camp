package app

import (
	"camp-summer/internal/initializers"
	"camp-summer/internal/model/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (*appController) Create(context *gin.Context) {
	var body struct {
		Name         string  `form:"name" binding:"required" json:"name"`
		FirstNumber  float64 `form:"first_number" binding:"required" json:"first_number"`
		SecondNumber float64 `form:"second_number" binding:"required" json:"second_number"`
		Text         string  `form:"text" binding:"required" json:"text"`
	}

	if err := context.BindJSON(&body); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productType := app.App{
		Name:         body.Name,
		FirstNumber:  body.FirstNumber,
		SecondNumber: body.SecondNumber,
		Text:         body.Text,
	}

	result := initializers.DB.Create(&productType)

	if result.Error != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	context.IndentedJSON(http.StatusCreated, gin.H{"data": &productType})
}

func (*appController) GetById(context *gin.Context) {
	id := context.Param("id")

	var productType app.App
	if err := initializers.DB.Where("id = ?", id).First(&productType).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"data": &productType})
}

func (*appController) Patch(context *gin.Context) {
	// Get ID parameter from URL
	id := context.Param("id")

	// Fetch the existing data from the database based on the ID
	var existingProduct app.App
	if err := initializers.DB.Where("id = ?", id).First(&existingProduct).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	var updates struct {
		Name           *string  `json:"name_for_admin_patch"`
		FirstNumber    *float64 `json:"first_number_for_admin_patch"`
		SecondNumber   *float64 `json:"second_number_for_admin_patch"`
		ExpectedNumber *float64 `json:"expected_number"`
		Text           *string  `json:"text_for_admin_patch"`
	}

	if err := context.BindJSON(&updates); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updates.Name != nil {
		existingProduct.Name = *updates.Name
	}
	if updates.FirstNumber != nil {
		existingProduct.FirstNumber = *updates.FirstNumber
	}
	if updates.SecondNumber != nil {
		existingProduct.SecondNumber = *updates.SecondNumber
	}
	if updates.Text != nil {
		existingProduct.Text = *updates.Text
	}
	if updates.ExpectedNumber != nil {
		existingProduct.ExpectedNumber = updates.ExpectedNumber
	}

	if err := initializers.DB.Save(&existingProduct).Error; err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"data": existingProduct})
}
