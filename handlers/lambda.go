package handlers

import(

  "context"
  "encoding/json"
  "errors"
  "fmt"
  "os"

  "github.com/aws/aws-lambda-go/events"
  "github.com/sirupsen/logrus"

  "github.com/shekodn/exploring-streaming-data-analysis/bizsrvc"
  "github.com/shekodn/exploring-streaming-data-analysis/factories"
  "github.com/shekodn/exploring-streaming-data-analysis/models"

)

var dbTableName string
var log = logrus.New()

func init() {

  dbTableName = os.Getenv("DYNAMO_DB_NAME")

  if dbTableName == "" {
    panic("Table Name is empty")
  }
}

func LambdaHandler(ctx context.Context, kinesisEvent events.KinesisEvent) error {

    var events []models.Event_iface

    for _, record := range kinesisEvent.Records {
        kinesisRecord := record.Kinesis
        dataBytes := kinesisRecord.Data

        // Declare a new EventSniffer struct
        var eventSniffer models.EventSniffer
        json.Unmarshal(dataBytes, &eventSniffer)

        if (eventSniffer.Event == "") {
          return errors.New("eventSniffer is empty")
        }

        event := factories.GetEvent(eventSniffer.Event)

        if err := bizsrvc.AssemblingEvent(event, dataBytes); err != nil {
          log.Error(err)
          return err
        }

        events = append(events, event)
        log.Info("New Event:", event.String())

        // Row or not
        // row := factories.GetRow(event)
        // fmt.Printf("%T\n", row)

        // row := models.Row{Vin: "31"}
    }

    mappedList, err := Map(events);

    if err != nil {
        log.Error(err)
        return err
    }

    finalMap := Groupper(mappedList)

    log.Println("finalMap:", finalMap)

    log.Println("for finalMap:", finalMap)

    for _, row := range finalMap {

        fmt.Printf("%T\n", row)
        fmt.Printf("%T\n", row)
        models.Writer(row, dbTableName)

    }

    return nil
}

func Map(events []models.Event_iface) ([]models.Row_iface, error) {

  mappedList := []models.Row_iface{}

  for _, event := range events {

    row := factories.GetSomeRow(event)

    bizsrvc.AssemblingRow(row, event)

    if row != nil {
        log.Info("Added:", event.String())
        mappedList = append(mappedList, row)
    } else {
      log.Info("Skipped:", event.String())
    }
  }

  return mappedList, nil
}

func Groupper(reducedList []models.Row_iface)  map[string]models.Row_iface {

  // Creates a final map data structure with only one value pero Vin (ID)
  final := make(map[string]models.Row_iface)

  for _, list := range reducedList {

    // Stores truck identifier
    truckVin := list.GetVin()

    if _, ok := final[truckVin]; ok {
      // Checks which time is more recent. Current time or the one that it is
      // being processed
      if final[truckVin].GetLocationTs().Before(list.GetLocationTs()) {
        // If current mapped time happened before processed time, updates time
        // to keep the most recent one
        final[truckVin] = list
      }
    } else {
      final[truckVin] = list
    }
  }

  return final
}
