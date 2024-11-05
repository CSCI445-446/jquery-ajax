package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"

	"jquery-demo/templates"
	"jquery-demo/types"
)

func connect() *sql.DB {
	config := mysql.Config{
		User: "jq",
		Passwd: "password",
		DBName: "jq_demo",
	}

	sql, _ := sql.Open("mysql", config.FormatDSN())

	return sql
}

func findCustomersByLastName(conn *sql.DB, lastNameFragment string) []types.Customer {
	stmt, _ := conn.Prepare("SELECT * FROM customer WHERE last_name LIKE ?")
	arg := fmt.Sprintf("%%%s%%", lastNameFragment)
	rows, _ := stmt.Query(arg)

	var customers []types.Customer
	var customer types.Customer

	for rows.Next() {
		_ = rows.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.Email)
		customers = append(customers, customer)
	}

	return customers
}

func main() {
	e := echo.New()

	e.GET("/", func(ctx echo.Context) error {
		return Render(ctx, http.StatusOK, templates.CustomerSearch())
	})

	e.GET("/customer_search", func(ctx echo.Context) error {
		conn := connect()
		lastNameQueryFragment := ctx.QueryParam("lastName")
		customers := findCustomersByLastName(conn, lastNameQueryFragment)
		return Render(ctx, http.StatusOK, templates.SearchResults(customers))
	})

	e.Logger.Fatal(e.Start(":8000"))
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
