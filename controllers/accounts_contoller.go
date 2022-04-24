package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/bee7ch7/go-gin-postgres-sqlx/db"
	"github.com/gin-gonic/gin"
)

type Account struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type NewAccount struct {
	Owner    string `json:"owner"`
	Currency string `json:"currency"`
}

func CreateAccount(c *gin.Context) {

	var req NewAccount
	var account Account

	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "bad body request",
		})
		return
	}

	var balance int64 = 0 // fresh account only with 0 balance
	// "INSERT INTO accounts(owner, currency, balance) VALUES($1, $2, $3) returning id, owner, balance, currency, created_at;"
	err := db.DBConn.QueryRow("INSERT INTO accounts(owner, currency, balance) VALUES($1, $2, $3) returning id, owner, currency, balance, created_at;",
		req.Owner, req.Currency, balance).Scan(&account.ID, &account.Owner, &account.Currency, &account.Balance, &account.CreatedAt)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "cannot get data from the server",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"error": false,
		"id":    account,
	})
}

func GetAccounts(c *gin.Context) {
	var accounts []Account

	err := db.DBConn.Select(&accounts, "Select id, owner, balance, currency, created_at from accounts;")

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "cannot get data from the server",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, accounts)

}

func GetAccount(c *gin.Context) {
	req_account_id := c.Param("id")
	id, err := strconv.Atoi(req_account_id)
	var singleAccount Account

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "the id must be integer",
		})
		return
	}

	err = db.DBConn.Get(&singleAccount, "Select id, owner, balance, currency, created_at from accounts where id=$1;", id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "cannot fetch the value, account doesn't exist",
		})
		return

	}

	c.IndentedJSON(http.StatusOK, singleAccount)

}
