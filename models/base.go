package models

import (
  
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
