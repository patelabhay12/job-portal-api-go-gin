package handlers

import (
	"database/sql"
	"job_portal/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user Id"})
			return
		}

		user, err := services.GetUserById(db, id)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, user)

	}
}
