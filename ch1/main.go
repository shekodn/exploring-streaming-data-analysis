package main

import(
  "fmt"
  "encoding/csv"
  "encoding/json"
  "os"
  "strconv"
  "sync"
  "time"

  // "github.com/aws/aws-sdk-go/aws"
  // "github.com/aws/aws-lambda-go"
  // "github.com/aws/aws-sdk-go/service/kinesis"
  // "github.com/aws/aws-lambda-go/lambda"
  "github.com/google/uuid"
)

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

/////
// Creates Aggregator interface in order to be able to create an aggregation
// according to the type: TruckArrives (TA), TruckDeparts (TD), MechanicChangesOil (MCO)
type Aggregator_iface interface {}

// object Aggregator {
//
//   def map(event: Event): Option[Row] = event match {
//     case TA(ts, v, loc) => Some(Row(v.vin, v.mileage, None, Some(loc, ts)))
//     case TD(ts, v, loc) => Some(Row(v.vin, v.mileage, None, Some(loc, ts)))
//     case MCO(ts, _, v)  => Some(Row(v.vin, v.mileage, Some(v.mileage), None))
//     case _              => None
//   }

type TruckArrivesAggregator struct {
  // fmt.Println("TruckArrivesAggregator")
  // ts, v, loc
}

type TruckDepartsAggregator struct {
  // fmt.Println("TruckDepartsAggregator")
  // ts, v, loc
}

type MechanicChangesOilAggregator struct {
  // fmt.Println("MechanicChangesOilAggregator")
  // ts, _, v
}

func NewTruckArrivesAggregator() {
  fmt.Println("TruckArrivesAggregator")
  // ts, v, loc
}

func NewTruckDepartsAggregator() {
  fmt.Println("TruckDepartsAggregator")
  // ts, v, loc
}

func NewMechanicChangesOilAggregator() {
  fmt.Println("MechanicChangesOilAggregator")
  // ts, _, v
}

func GetRelevantRow(event string) Aggregator_iface {
	switch event {

    case "TRUCK_ARRIVES":
      return nil
    case "TRUCK_DEPARTS":
      return nil
    default:
  		fmt.Println("type undefined")
  		return nil
  }
}

//

type Row struct {
  EventType string
  Vin string
  Mileage int
  MileageAtOilChange int //optional - int
  Location //option - location, DateTime
  Timestamp time.Time //aux
  
}


func main() {

  // req := []byte(`{"event":"TRUCK_ARRIVES", "location": {"elevation":7,
  // "latitude":51.522834, "longitude":-0.081813},
  // "timestamp": "2018-01-12T12:42:00Z", "vehicle": {"mileage":33207,
  // "vin":"1HGCM82633A004352"}}`)
  //
  // // Declare a new EventSniffer struct
  // var eventSniffer EventSniffer
  // json.Unmarshal([]byte(req), &eventSniffer)
  //
  // event := GetTypeOfEvent(eventSniffer.Event, req)
  //
  // fmt.Println(event)

  fileName := "sample.csv"

  f, err := os.Open(fileName)

  if err != nil {
    panic(err)
  }

  defer f.Close()

  // Read Files into a Variable
  lines, err := csv.NewReader(f).ReadAll()

  if err != nil {
    panic(err)
  }

  // Create the first mapper list
  // Reference: https://medium.com/@jayhuang75/a-simple-mapreduce-in-go-42c929b000c5
  lists := make(chan []Row)

  // Ensure the final value after Reducer is obtained.
  finalValue := make(chan []Row)

  // Ensure all send operations are done.
  var wg sync.WaitGroup

  // Mapping
  wg. Add(len(lines))

  for _, line := range lines {
    go func(event []string) {
      defer wg.Done()
      lists <- Map(event)
    }(line)
  }

  go Reducer(lists, finalValue)

  wg.Wait()
  close(lists)
  ch := <-finalValue

  grouppedMap := Groupper(ch)

  for _, value := range grouppedMap {
    fmt.Println(value)
  }
}

// Mapper Implementation - Separates irrelevant rows from non-relevant.
// Relevant rows are: TruckArrives (TA), TruckDeparts (TD), MechanicChangesOil (MCO)
func Map(event []string) []Row {
  list := []Row{}

  // "TRUCK_ARRIVES","6","51.522834","-0.081813","2018-01-12T12:42:00Z","33207","1HGCM82633A004352"
  eventType := event[0]
  elevation := event[1]
  latitude  := event[2]
  longitude := event[3]
  timestamp := event[4]
  mileage   := event[5]
  truckVin  := event[6]

  parsedElevation, err := strconv.Atoi(elevation)

  if err != nil {
    panic(err)
  }

  parsedMileage, err := strconv.Atoi(mileage)

  if err != nil {
    panic(err)
  }


  parsedTs, err := time.Parse(time.RFC3339, timestamp)

  if err != nil {
    panic(err)
  }

  parsedLat, err := strconv.ParseFloat(latitude, 64)

  if err != nil {
    panic(err)
  }


  parsedLong, err := strconv.ParseFloat(longitude, 64)

  if err != nil {
    panic(err)
  }

  list = append(list, Row {
    EventType: eventType,
    Location: Location {
      Elevation: parsedElevation,
      Latitude: parsedLat,
      Longitude: parsedLong,
    },
    Timestamp: parsedTs,
    Mileage: parsedMileage,
    Vin: truckVin,
  })

  return list
}


func Reducer(mapList chan []Row, sendFinalValue chan []Row) {

  final := []Row{}

  for list := range mapList {
    for _, value := range list {
      if (value.EventType == "TRUCK_ARRIVES") {
        final = append(final, value)
      }
    }
  }
  sendFinalValue <- final

}

func Groupper(reducedList []Row)  map[string]Row {

  final := make(map[string]Row)

  for _, list := range reducedList {

    // Stores truck identifier
    truckVin := list.Vin

    if _, ok := final[truckVin]; ok {
      // Checks which time is more recent.
      if final[truckVin].Timestamp.Before(list.Timestamp) {
        final[truckVin] = list
      }
    } else {
      final[truckVin] = list
    }
  }

  return final
}
