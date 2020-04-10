package bizsrvc

import (

  _ "errors"

  "github.com/shekodn/exploring-streaming-data-analysis/models"
)

func AssemblingRow(row models.Row_iface, event models.Event_iface) {
  	if row != nil {
  		row.Assemble(event)
  	}
}
