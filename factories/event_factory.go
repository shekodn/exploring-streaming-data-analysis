package factories

import (
  "github.com/shekodn/exploring-streaming-data-analysis/models"
)

// Assamble EVENT based on the type
func GetEvent(event string) models.Event_iface {
	switch event {

    case "TRUCK_ARRIVES":
      return &models.TruckArrives{}

    case "TRUCK_DEPARTS":
      return &models.TruckDeparts{}

    case "MECHANIC_CHANGES_OIL":
      return &models.MechanicChangesOil{}

    case "DRIVER_DELIVERS_PACKAGE":
      return &models.DriverDeliversPackage{}

    case "DRIVER_MISSES_CUSTOMER":
      return &models.DriverMissesCustomer{}

    default:
  		return nil
  }
}
