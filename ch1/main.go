package main

import(
  "fmt"
  "encoding/json"
  "time"

  // "github.com/aws/aws-sdk-go/aws"
  // "github.com/aws/aws-lambda-go"
  // "github.com/aws/aws-sdk-go/service/kinesis"
  // "github.com/aws/aws-lambda-go/lambda"
  "github.com/google/uuid"
)

// Creates event interface in order to be able to create an event according
// to the type
type Event_iface interface {}

type EventSniffer struct {
  Event string
}

type Employee struct {
  Id uuid.UUID
  JobRole string
}

type Vehicle struct {
  Vin string
  Mileage int
}

type Location struct {
  Latitude float64
  Longitude float64
  Elevation int
}

type Package struct {
  Id uuid.UUID
}

type Customer struct {
  Id uuid.UUID
  IsVip bool
}

//// Type of events
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


func main() {

  req := []byte(`{"event":"TRUCK_ARRIVES", "location": {"elevation":7,
  "latitude":51.522834, "longitude":-0.081813},
  "timestamp": "2018-01-12T12:42:00Z", "vehicle": {"mileage":33207,
  "vin":"1HGCM82633A004352"}}`)

  // Declare a new EventSniffer struct
  var eventSniffer EventSniffer
  json.Unmarshal([]byte(req), &eventSniffer)

  event := GetTypeOfEvent(eventSniffer.Event, req)

  fmt.Println(event)
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
