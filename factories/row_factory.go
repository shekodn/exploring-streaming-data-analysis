package factories

import (
  "github.com/shekodn/exploring-streaming-data-analysis/models"
)

func GetSomeRow(event models.Event_iface) models.Row_iface {
	switch event.(type)  {

    case *models.TruckArrives:
      return &models.SomeRow{}

    case *models.TruckDeparts:
      return &models.SomeRow{}

    case *models.MechanicChangesOil:
      return &models.SomeRow{}

    default:
  		return nil
  }
}
