package models

import(
  "fmt"
  _ "strconv"
  "time"
)

type Row_iface interface {
    Assemble(event Event_iface)
    String() string
    GetVin() string
    GetLocationTs() time.Time
}

type Row struct {
    Vin string
    Latitude float64
    Longitude float64
    LocationTs time.Time
    Mileage int
    MileageAtOilChange int
}

type SomeRow struct {
    Row
}

func (r Row) String() string {
    return fmt.Sprintf("Row: %s - %f - %f - %s - %x - %x",
      r.Vin,
      r.Latitude,
      r.Longitude,
      r.LocationTs,
      r.Mileage,
      r.MileageAtOilChange,
    )
}

// Interfcae Methods
func (row *SomeRow) Assemble(event Event_iface) {
  	fmt.Println("Assembling components for SomeRow")

    row.Vin = event.GetVin()
    row.Latitude = event.GetLatitude()
    row.Longitude = event.GetLongitude()
    row.LocationTs = event.GetTimestamp()
    row.Mileage = event.GetMileage()
    var _ Row_iface = row // Enforce interface compliance
}

func (r Row) GetVin() string {
    return r.Vin
}

func (r Row) GetLocationTs() time.Time {
    return r.LocationTs
}

func (r Row) GetLatitude() float64 {
    return r.Latitude
}

func (r Row) GetLongitude() float64 {
    return r.Longitude
}

func (r Row) GetMileage() int {
    return r.Mileage
}

func (r Row) GetMileageAtOilChange() int{
    return r.MileageAtOilChange
}
