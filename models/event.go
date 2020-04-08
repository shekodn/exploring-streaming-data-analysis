package models

import (
  "encoding/json"
  "fmt"
  "time"
)


//// Type of events
// Creates event interface in order to be able to create an event according
// to the type
type Event_iface interface {}

type EventSniffer struct {
  Event string
}

type TruckArrives struct {
  Timestamp time.Time
  Vehicle Vehicle
  Location Location
}

type TruckDeparts struct {
  Timestamp time.Time
  Vehicle Vehicle
  Location Location
}

type MechanicChangesOil struct {
  Timestamp time.Time
  Employee Employee
  Vehicle Vehicle
}

type DriverDeliversPackage struct {
  Timestamp time.Time
  Employee Employee
  // `package`: Package, customer: Customer, location: Location)
}

type DriverMissesCustomer struct {
  Timestamp time.Time
  Employee Employee
  // `package`: Package, customer: Customer, location: Location)
}

// Factory - Assamble
func GetTypeOfEvent(event string, req []byte) Event_iface {
	switch event {

    case "TRUCK_ARRIVES":
      return NewTruckArrives(req)

    case "TRUCK_DEPARTS":
      return NewTruckDeparts(req)

    case "MECHANIC_CHANGES_OIL":
      return NewMechanicChangesOil(req)

    case "DRIVER_DELIVERS_PACKAGE":
      return NewDriverDeliversPackage(req)

    case "DRIVER_MISSES_CUSTOMER":
      return NewDriverMissesCustomer(req)

    default:
  		fmt.Println("type undefined")
  		return nil
  }
}


// Factory Methods
func NewTruckArrives(req []byte) *TruckArrives {
	event := &TruckArrives{}
  json.Unmarshal(req, &event)
	var _ Event_iface = event // Enforce interface compliance
	return event
}

func NewTruckDeparts(req []byte) *TruckDeparts {
  event := &TruckDeparts{}
  json.Unmarshal(req, &event)
  var _ Event_iface = event // Enforce interface compliance
  return event
}

func NewMechanicChangesOil(req []byte) *MechanicChangesOil {
  event := &MechanicChangesOil{}
  json.Unmarshal(req, &event)
  var _ Event_iface = event // Enforce interface compliance
  return event
}

func NewDriverDeliversPackage(req []byte) *DriverDeliversPackage {
  event := &DriverDeliversPackage{}
  json.Unmarshal(req, &event)
  var _ Event_iface = event // Enforce interface compliance
  return event
}

func NewDriverMissesCustomer(req []byte) *DriverMissesCustomer {
  event := &DriverMissesCustomer{}
  json.Unmarshal(req, &event)
  var _ Event_iface = event // Enforce interface compliance
  return event
}
