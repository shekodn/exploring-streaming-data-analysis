package handlers

import(

  "context"
  "encoding/json"
  "fmt"
  "os"

  "github.com/aws/aws-lambda-go/events"

  "github.com/shekodn/exploring-streaming-data-analysis/models"

)

var dbTableName string

func init() {

  dbTableName = os.Getenv("DYNAMO_DB_NAME")
  
  if dbTableName == "" {
    panic("Table Name is empty")
  }
}

func LambdaHandler(ctx context.Context, kinesisEvent events.KinesisEvent) error {

    for _, record := range kinesisEvent.Records {
        kinesisRecord := record.Kinesis
        dataBytes := kinesisRecord.Data
        // dataText := string(dataBytes)

        // Declare a new EventSniffer struct
        var eventSniffer models.EventSniffer
        json.Unmarshal([]byte(dataBytes), &eventSniffer)

        event := models.GetTypeOfEvent(eventSniffer.Event, dataBytes)
        fmt.Println(event)

        row := models.Row{Vin: "31"}

        models.Writer(row, dbTableName)
    }

    return nil
}
