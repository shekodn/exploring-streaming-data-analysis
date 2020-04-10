package models

import (
  "encoding/json"
  "fmt"
  "time"
)


//// Type of events
// Creates event interface in order to be able to create an event according
// to the type
type Event_iface interface {
    Assemble(dataBytes []byte)
    String() string
}

type Event struct {
  Timestamp time.Time
}

type EventSniffer struct {
  Event string
}

type TruckArrives struct {
  Event
  Vehicle Vehicle
  Location Location
}

type TruckDeparts struct {
  Event
  Vehicle Vehicle
  Location Location
}

type MechanicChangesOil struct {
  Event
  Employee Employee
  Vehicle Vehicle
}

type DriverDeliversPackage struct {
  Event
  Employee Employee
  // `package`: Package, customer: Customer, location: Location)
}

type DriverMissesCustomer struct {
  Event
  Employee Employee
  // `package`: Package, customer: Customer, location: Location)
}

// Interfcae Methods

func (event *TruckArrives) Assemble(dataBytes []byte) {
	fmt.Println("Assembling components for TruckArrives")
  json.Unmarshal(dataBytes, &event)
  var _ Event_iface = event // Enforce interface compliance
}

func (event *TruckDeparts) Assemble(dataBytes []byte) {
	fmt.Println("Assembling components for TruckDeparts")
  json.Unmarshal(dataBytes, &event)
  var _ Event_iface = event // Enforce interface compliance
}

func (event *MechanicChangesOil) Assemble(dataBytes []byte) {
	fmt.Println("Assembling components for MechanicChangesOil")
  json.Unmarshal(dataBytes, &event)
  var _ Event_iface = event // Enforce interface compliance
}

func (event *DriverDeliversPackage) Assemble(dataBytes []byte) {
	fmt.Println("Assembling components for DriverDeliversPackage")
  json.Unmarshal(dataBytes, &event)
  var _ Event_iface = event // Enforce interface compliance
}

func (event *DriverMissesCustomer) Assemble(dataBytes []byte) {
	fmt.Println("Assembling components for DriverMissesCustomer")
  json.Unmarshal(dataBytes, &event)
  var _ Event_iface = event // Enforce interface compliance
}


func (e TruckArrives) String() string {
    return fmt.Sprintf("TruckArrives: %s - %s - %s - %s - %s - %s",
      e.Timestamp,
      e.Vehicle.Vin, e.Vehicle.Mileage,
      e.Location.Latitude, e.Location.Longitude, e.Location.Elevation,
    )
}

func (e TruckDeparts) String() string {
    return fmt.Sprintf("TruckDeparts: %s - %s - %s - %s - %s - %s",
      e.Timestamp,
      e.Vehicle.Vin, e.Vehicle.Mileage,
      e.Location.Latitude, e.Location.Longitude, e.Location.Elevation,
    )
}

func (e MechanicChangesOil) String() string {
    return fmt.Sprintf("MechanicChangesOil: %s - %s - %s - %s - %s",
      e.Timestamp,
      e.Employee.Id, e.Employee.JobRole,
      e.Vehicle.Vin, e.Vehicle.Mileage,
    )
}

func (e DriverDeliversPackage) String() string {
    return fmt.Sprintf("DriverDeliversPackage: %s - %s - %s",
      e.Timestamp,
      e.Employee.Id, e.Employee.JobRole,
    )
}

func (e DriverMissesCustomer) String() string {
    return fmt.Sprintf("DriverMissesCustomer: %s - %s - %s",
      e.Timestamp,
      e.Employee.Id, e.Employee.JobRole,
    )
}
