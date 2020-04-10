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
    GetVin() string
    GetTimestamp() time.Time
    GetLatitude() float64
    GetLongitude() float64
    GetMileage() int
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

func (e Event) GetTimestamp() time.Time {
    return e.Timestamp
}

func (e TruckArrives) GetVin() string {
    return e.Vehicle.Vin
}

func (e TruckDeparts) GetVin() string {
    return e.Vehicle.Vin
}

func (e MechanicChangesOil) GetVin() string {
    return e.Vehicle.Vin
}

func (e DriverDeliversPackage) GetVin() string {
    return ""
}

func (e DriverMissesCustomer) GetVin() string {
    return ""
}
func (e TruckArrives) GetLatitude() float64 {
    return e.Location.Latitude
}

func (e TruckDeparts) GetLatitude() float64 {
    return e.Location.Latitude
}

func (e MechanicChangesOil) GetLatitude() float64 {
    return 0
}

func (e DriverDeliversPackage) GetLatitude() float64 {
    return 0
}

func (e DriverMissesCustomer) GetLatitude() float64 {
    return 0
}

func (e TruckArrives) GetLongitude() float64 {
    return e.Location.Longitude
}

func (e TruckDeparts) GetLongitude() float64 {
    return e.Location.Longitude
}

func (e MechanicChangesOil) GetLongitude() float64 {
    return 0
}

func (e DriverDeliversPackage) GetLongitude() float64 {
    return 0
}

func (e DriverMissesCustomer) GetLongitude() float64 {
    return 0
}

func (e TruckArrives) GetMileage() int {
    return e.Vehicle.Mileage
}

func (e TruckDeparts) GetMileage() int {
    return e.Vehicle.Mileage
}

func (e MechanicChangesOil) GetMileage() int {
    return e.Vehicle.Mileage
}

func (e DriverDeliversPackage) GetMileage() int {
    return 0
}

func (e DriverMissesCustomer) GetMileage() int {
    return 0
}
