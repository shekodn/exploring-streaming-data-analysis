package bizsrvc

import (

  "errors"

  "github.com/shekodn/exploring-streaming-data-analysis/models"
)

func AssemblingEvent(event models.Event_iface, dataBytes []byte) error {
	if event != nil {
		event.Assemble(dataBytes)
    return nil
	}

  return errors.New("Couldn't Assemble the Event")
}
