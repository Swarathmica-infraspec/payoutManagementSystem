package payoutmanagementsystem

import (
	"context"
	"database/sql"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var store *PayeePostgresDB

func initStore() *PayeePostgresDB {
	if store != nil {
		return store
	}
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@db:5432/postgres?sslmode=disable"
	}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	store = PostgresPayeeDB(db)
	return store
}

func PayeePostAPI() {

	store := initStore()

	router := gin.Default()

	router.POST("/payees", func(c *gin.Context) {
		var req struct {
			Name     string `json:"name"`
			Code     string `json:"code"`
			AccNo    int    `json:"account_number"`
			IFSC     string `json:"ifsc"`
			Bank     string `json:"bank"`
			Email    string `json:"email"`
			Mobile   int    `json:"mobile"`
			Category string `json:"category"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
			return
		}

		p, err := NewPayee(req.Name, req.Code, req.AccNo, req.IFSC, req.Bank, req.Email, req.Mobile, req.Category)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed", "details": err.Error()})
			return
		}

		id, err := store.Insert(context.Background(), p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "DB insert failed", "details": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": id})
	})

	router.Run(":8080")
}
