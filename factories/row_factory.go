package factories

import (
  "github.com/shekodn/exploring-streaming-data-analysis/models"
)

// Assamble EVENT based on the type
func GetRow(event models.Event_iface) models.Row_iface {
	switch event.(type)  {

    case *models.TruckArrives:
      return models.SomeRow{}

    case *models.TruckDeparts:
      return models.SomeRow{}

    case *models.MechanicChangesOil:
      return models.NoneRow{}

    case *models.DriverDeliversPackage:
      return models.NoneRow{}

    case *models.DriverMissesCustomer:
      return models.NoneRow{}

    default:
  		return nil
  }
}
