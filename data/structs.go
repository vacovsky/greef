package data

import "time"

// EnvironmentalParameters encapsulates all params we can measure and care about relating to the aquarium
type EnvironmentalParameters struct {
	ID                    int `gorm:"AUTO_INCREMENT;column:id;primary_key"`
	ReadingDateTime       time.Time
	AmbientAirTemperature float64
	AmbientHumidity       float64
	WaterTemperature      float64
	AmbientAirCO2         float64
	AmbientSumpLight      float64
	AmbientLight          float64
	WaterPH               float64
	WaterTdsPpm           float64
}

// Livestock is a catalog of commnity members within the aquarium
type Livestock struct {
	ID             int `gorm:"AUTO_INCREMENT;column:id;primary_key"`
	CommonName     string
	AquiredDate    time.Time
	Alive          bool
	ScientificName string
	Description    string
	AnimalType     string // fish, invert, coral, etc
}

// SystemDetails contains information about the establishment of the aquatic system
type SystemDetails struct {
	StartedDate             time.Time
	DisplayTankGallons      float64
	DisplayTankHeightInches float64
	DisplayTankWidthInches  float64
	DisplayTankDepthInches  float64
	SumpGallons             float64
	SumpWidthInches         float64
	SumpDepthInches         float64
	SumpHeightInches        float64
	ReturnPumpGPH           int
	OverflowBoxGPH          int
}

type Chart struct {
	Series []string
	Labels []int64
	Data   [][]string
}
