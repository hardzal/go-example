package main

import (
	"database/sql"
	"fmt"
)

import _ "github.com/go-sql-driver/mysql"

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func connect() (*sql.DB, error) {
	// user => root
	// password =>
	// host => 127.0.0.1 atau localhost
	// port => 3306
	// dbname => golang_first
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_first")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var age = 27
	rows, err := db.Query("SELECT id, name, grade FROM tb_student WHERE age = ?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []student

	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.name)
	}

}

func sqlQueryRow() {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var result = student{}
	var id = "E001"
	err = db.QueryRow("SELECT name, grade FROM tb_student WHERE id = ?", id).Scan(&result.name, &result.grade)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Name\t: %s\nGrade\t: %d\n", result.name, result.grade)
}

func sqlPrepare() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT name, grade FROM tb_student WHERE id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result1 = student{}
	stmt.QueryRow("E001").Scan(&result1.name, &result1.grade)
	fmt.Printf("Name: %s\tGrade: %d\n", result1.name, result1.grade)

	var result2 = student{}
	stmt.QueryRow("W001").Scan(&result2.name, &result2.grade)
	fmt.Printf("Name: %s \tGrade: %d\n", result2.name, result2.grade)

	var result3 = student{}
	stmt.QueryRow("B001").Scan(&result3.name, &result3.grade)
	fmt.Printf("Name: %s\tGrade: %d\n", result3.name, result3.grade)
}

func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO tb_student VALUES (?, ?, ?, ?)", "G001", "Galahad", 29, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Insert success!")

	_, err = db.Exec("UPDATE tb_student SET age = ? WHERE id = ?", 20, "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Update success!")

	_, err = db.Exec("DELETE FROM tb_student WHERE id = ?", "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Delete success!")
}

func main() {
	// sqlQuery()
	// sqlQueryRow()
	// sqlPrepare()
	sqlExec()
}
