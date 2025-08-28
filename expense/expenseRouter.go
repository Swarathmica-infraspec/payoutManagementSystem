package payoutmanagementsystem

import (
	"context"
	"database/sql"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func PayeePostAPI() {

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@db:5432/postgres?sslmode=disable"
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	store := NewPostgresExpenseDB(db)

	router := gin.Default()

	router.POST("/expense", func(c *gin.Context) {
		var req struct {
			Title    string `json:"title"`
			Amount   float64 `json:"amount"`
			DateIncurred string    `json:"dateIncurred"`
			Category     string `json:"category"`
			Notes     string `json:"notes"`
			PayeeID    int `json:"payeeID"`
			ReceiptURI   string    `json:"receiptURI"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
			return
		}

		p, err := NewExpense(req.Title, req.Amount, req.DateIncurred, req.Category, req.Notes, req.PayeeID, req.ReceiptURI)
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
