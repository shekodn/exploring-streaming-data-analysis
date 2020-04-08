package models

import (
  "fmt"

  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/guregu/dynamo"
)

var db *dynamo.DB

func init () {
  // session := db.New(session.New())

  db = dynamo.New(session.New())

}

func GetDb() *dynamo.DB {
    return db
}

func Writer(theRow Row, dbTableName string) {

  table := db.Table(dbTableName)

  fmt.Println("Attempting to write in", dbTableName)

  // Put a new item, only if it doesn't already exist.
  err1 := table.Put(theRow).If("attribute_not_exists(Vin)").Run()

  if err1 != nil {
    // If it already exists, tries to update element
    fmt.Println("Checking if the update is going to be made or not ...")
    err2 := table.Update("Vin", theRow.Vin).
        // Set("Mileage", theRow.Mileage).
        // Set("Timestamp", theRow.Timestamp).
        // TODO: Update the whole object
        // If("'Timestamp' < ?", theRow.Timestamp).
        Run()

    if err2 != nil {
      // Cannot update, because condition failed
      fmt.Println(err2)
    } else {
      // Condition met - Updated!
      fmt.Println("Success: Conditional Write Made")
    }

  } else {
    // New Row with new Vin is made
    fmt.Println("New Entry")
  }

}
