package main

import (
  "database/sql"
  "fmt"
  "log"

  _ "github.com/go-sql-driver/mysql"
)

const (
  db_user = "USER"
  db_passwd = "PASSWORD"
  db_addr = "IP/ENDPOINT"
  db_db = "DATABASE"
)

type Person struct {
  Id int
  Name string
  Age int
  Location string
}

func main() {
  db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_passwd, db_addr, db_db))
  HandleError(err)
  defer db.Close()

  // Insert data
  people := GetData()
  err = insertData(db, people)
  HandleError(err)

  // Get all data
  people, err = getAllData(db)
  HandleError(err)
  fmt.Println(people)
  fmt.Printf("We have %v people in our db!\n", len(people))

  // Get all people above 30
  people, err = getAllAboveAge(db, 30)
  HandleError(err)
  fmt.Println(people)
  fmt.Printf("We have %v people in our db that are older than 30!\n", len(people))

  // Delete all people above 30
  err = deleteAllAboveAge(db, 30)
  HandleError(err)

  // Update Marianne's age to 21
  err = updatePersonAge(db, "Marianne C Reagan", 21)
  HandleError(err)
}

// Function to clean up the main function
func HandleError(err error) {
  if err != nil {
    fmt.Println(err)
    log.Fatal(err)
  }
}

// Insert []Person in the database
func insertData(db *sql.DB, people []Person) error {
  for _, person := range people {
    q := "INSERT INTO `person` (name, age, location) VALUES (?, ?, ?)"
    insert, err := db.Prepare(q)
    defer insert.Close()

    if err != nil {
      return err
    }

    _, err = insert.Exec(person.Name, person.Age, person.Location)
    if err != nil {
      return err
    }
  }

  return nil
}

// Select all people/data
func getAllData(db *sql.DB) (people []Person, err error) {
  resp, err := db.Query("SELECT * FROM `person`")
  defer resp.Close()

  if err != nil {
    return people, err
  }

  for resp.Next() {
    var pPerson Person
    err = resp.Scan(&pPerson.Id, &pPerson.Name, &pPerson.Age, &pPerson.Location)
    if err != nil {
      return people, err
    }

    people = append(people, pPerson)
  }

  return people, nil
}

// Select all people above `age`
func getAllAboveAge(db *sql.DB, age int) (people []Person, err error) {
  q := "SELECT * FROM `person` WHERE `age` > ? "
  resp, err := db.Query(q, age)
  defer resp.Close()

  if err != nil {
    return people, err
  }

  for resp.Next() {
    var pPerson Person
    err = resp.Scan(&pPerson.Id, &pPerson.Name, &pPerson.Age, &pPerson.Location)
    if err != nil {
      return people, err
    }

    people = append(people, pPerson)
  }

  return people, nil
}

// Delete all people above `age`
func deleteAllAboveAge(db *sql.DB, age int) error {
  q := "DELETE FROM `person` WHERE `age` > ?"
  drop, err := db.Prepare(q)
  defer drop.Close()

  if err != nil {
    return err
  }

  _, err = drop.Exec(age)
  if err != nil {
    return err
  }

  return nil
}

// Update a persons `age` based on `name`
func updatePersonAge(db *sql.DB, name string, age int) error {
  q := "UPDATE `person` SET `age` = ? WHERE `name` like ?"
  update, err := db.Prepare(q)
  defer update.Close()

  if err != nil {
    return err
  }

  _, err = update.Exec(age, name)
  if err != nil {
    return err
  }

  return nil

}
