package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/viza/banking/app/errs"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Dateofbirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	var rows *sql.Rows
	var err error

	if status == "" { // e.g. localhost:8000/customers
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		rows, err = d.client.Query(findAllSql)

	} else { // e.g. localhost:8000/customers?status=active
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		rows, err = d.client.Query(findAllSql, status)

	}

	if err != nil {
		log.Println("Error while quering customer table " + err.Error())
		return nil, errs.NewUnexpectedError("Error while quering customer table " + err.Error())
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Dateofbirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning customers " + err.Error())
			return nil, errs.NewUnexpectedError("Error while scanning customers " + err.Error())
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	client, err := sql.Open("mysql", "root:!vetalr00t!@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDB{client: client}
}
