package transactioncontroller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taufiqtab/dulrestful/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var transactions []models.Transaction

	models.DB.Find(&transactions)
	c.JSON(http.StatusOK, gin.H{"transaction": transactions})
}

func Show(c *gin.Context) {
	var transaction models.Transaction
	id := c.Param("id")

	if err := models.DB.First(&transaction, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"transaction": transaction})
}

func Create(c *gin.Context) {
	var transaction models.Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&transaction)
	c.JSON(http.StatusOK, gin.H{"transaction": transaction})
}

func Update(c *gin.Context) {
	var transaction models.Transaction
	id := c.Param("id")

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&transaction).Where("id = ?", id).Updates(&transaction).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "tidak dapat mengupdate data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diupdate"})
}

func Delete(c *gin.Context) {
	var transaction models.Transaction

	//input := map[string]string{"id": "0"}
	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	//id, _ := strconv.ParseInt(input["id"], 10, 64)
	id, _ := input.Id.Int64()
	if models.DB.Delete(&transaction, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat menghapus data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
