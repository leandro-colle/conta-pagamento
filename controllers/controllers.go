package controllers

import (
	"conta-pagamento/database"
	"conta-pagamento/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

func GetAll(c *gin.Context) {
	var accounts []models.Account
	database.DB.Find(&accounts)
	c.JSON(http.StatusOK, accounts)
}

// @BasePath /api/v1

// GetById godoc
// @Summary Get account
// @Description Gets an account by ID
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path int true "Account ID"
// @Success 200 {object} models.Account
// @Failure 400 {object} httputil.HTTPError
// @Router /account/{id} [get]
func GetById(c *gin.Context) {
	var account models.Account

	database.DB.First(&account, c.Params.ByName("id"))

	if account == (models.Account{}) {
		c.JSON(http.StatusNotFound, gin.H{"error": "account is not found"})
		return
	}

	c.JSON(http.StatusOK, account)
}

func Create(c *gin.Context) {
	var account models.Account

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.Validate(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&account)

	c.JSON(http.StatusOK, account)
}

func Delete(c *gin.Context) {
	var accountId = c.Params.ByName("id")
	var account models.Account

	database.DB.First(&account, accountId)

	if account == (models.Account{}) {
		c.JSON(http.StatusNotFound, gin.H{"error": "account is not found"})
		return
	}

	database.DB.Delete(&account, accountId)

	c.JSON(http.StatusOK, account)
}

func Update(c *gin.Context) {
	var accountId = c.Params.ByName("id")
	var account models.Account

	database.DB.First(&account, accountId)

	if account == (models.Account{}) {
		c.JSON(http.StatusNotFound, gin.H{"error": "account is not found"})
		return
	}

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.Validate(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&account).UpdateColumns(account)

	c.JSON(http.StatusOK, account)
}

func GetByAgencyNumber(c *gin.Context) {
	var agencyNumber = c.Param("agencynumber")
	var accounts []models.Account

	database.DB.Where(&models.Account{AgencyNumber: agencyNumber}).Find(&accounts)

	if len(accounts) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "account is not found"})
		return
	}

	c.JSON(http.StatusOK, accounts)
}
