package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (db CustomerRepositoryDB) FindAll() ([]Customer, error) {

	findAllSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"
	rows, err := db.client.Query(findAllSQL)
	if err != nil {
		log.Println("Error while querying customer table", err.Error())
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning customers", err.Error())
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	client, err := sql.Open("mysql", "root:codecamp@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDB{client}
}
