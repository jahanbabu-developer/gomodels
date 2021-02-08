package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var a int
	fmt.Println("Welcome to Employee crud Operation: 1.Insert 2.Delete 3.Update 4.View")
	fmt.Println("enter your Choice:")
	fmt.Scan(&a)
	switch a {
	case 1:
		insert()
		break
	case 2:
		Deleteemp()
		break
	case 3:
		Updateemp()
		break
	case 4:
		Viewemp()

	}

}
func insert() {
	db, err := sql.Open("mysql", "root:Decoders@123@tcp(127.0.0.1:3306)/sys")
	if err != nil {
		panic(err.Error())

	}
	defer db.Close()

	var empid, sal, cno int
	var name, desg, doj string
	fmt.Println("Enter the emplid:\t")
	fmt.Scan(&empid)
	fmt.Println("Enter the name: ")
	fmt.Scan(&name)
	fmt.Println("Enter the designation: ")
	fmt.Scan(&desg)
	fmt.Println("Enter the salary: ")
	fmt.Scan(&sal)
	fmt.Println("Enter the contact: ")
	fmt.Scan(&cno)
	fmt.Println("Enter the doj:")
	fmt.Scan(&doj)
	insert, err := db.Query("INSERT INTO employee VALUES(?,?,?,?,?,?)", name, desg, sal, doj, cno, empid)
	if err != nil {
		panic(err.Error())

	}
	insert.Close()
	fmt.Println("Successfully Inserted")
}

func Deleteemp() {
	db, err := sql.Open("mysql", "root:Decoders@123@tcp(127.0.0.1:3306)/sys")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var d int
	fmt.Println("Enter the id to delete:")
    fmt.Scan(&d)
	rows, err := db.Exec("DELETE FROM employee WHERE empid = ?", d)

	rowCount, _ := rows.RowsAffected()
	fmt.Printf("Deleted %d row(s) of data.\n", rowCount)
	fmt.Println("Done.")

}
func Updateemp() {
	db, err := sql.Open("mysql", "root:Decoders@123@tcp(127.0.0.1:3306)/sys")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var sal int
	var name string

	fmt.Println("Enter the name: ")
	fmt.Scan(&name)

	fmt.Println("Enter the sal: ")
	fmt.Scan(&sal)
	rows, err := db.Exec("UPDATE employee SET sal= ? WHERE name = ?", sal, name)

	rowCount, err := rows.RowsAffected()
	fmt.Printf("Updated %d row(s) of data.\n", rowCount)
	fmt.Println("Done.")
}

type employee struct {
	empid int
	name string
	desg string
	sal int
	doj string
	cno int
}

func Viewemp() {
	db, err := sql.Open("mysql", "root:Decoders@123@tcp(127.0.0.1:3306)/sys")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM  employee")
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()
	for results.Next() {

		var emp employee
		err := results.Scan(&emp.empid, &emp.name, &emp.desg, &emp.sal, &emp.doj, &emp.cno)

		if err != nil {
			panic(err)
		}

		fmt.Printf("%v\n", emp)
	}
}