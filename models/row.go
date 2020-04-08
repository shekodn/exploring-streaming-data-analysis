package models

import(
  "fmt"
  "strconv"
  "time"
)

// DynamoDB Schema:
// Truck VIN | Latitude | Longitude | Location timestamp | Mileage | Mileage at oil change

type Row struct {
  EventType string
  Vin string
  Location //option - location, DateTime
  Timestamp time.Time //aux
  Mileage int
  MileageAtOilChange int //optional - int
}


func Reducer(mapList chan []Row, sendFinalValue chan []Row) {

  final := []Row{}

  for list := range mapList {
    for _, value := range list {
      if (value.EventType == "TRUCK_ARRIVES") || (value.EventType == "TRUCK_DEPARTS") {
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
