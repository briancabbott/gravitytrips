package main

import "fmt"
import "io"
import "encoding/csv"
import "bufio"
import "os"
import "strconv"

/*
	Good Names:
		- Gravity Trips
		- Gravity Travels
*/

type TDTBounds struct {
	min_t	float64
	max_t 	float64

	min_x 	float64
	max_x 	float64
	
	min_y 	float64
	max_y 	float64
	
	min_z	float64
	max_z 	float64
}

type TemporalDisplacementTrips struct {
	ID  string
	T	float64
	X 	float64
	Y 	float64
	Z	float64
}

func computeTDTBoundaries(tdts []TemporalDisplacementTrips) TDTBounds {

	bounds := TDTBounds{}

	for _, tdt := range tdts {
		// min/max T
		if tdt.T < bounds.min_t || bounds.min_t == 0.0 {
			bounds.min_t = tdt.T
		}
		if tdt.T > bounds.max_t || bounds.max_t == 0.0 {
			bounds.max_t = tdt.T
		}

		// min/max X
		if tdt.X < bounds.min_x || bounds.min_x == 0.0 {
			bounds.min_x = tdt.X
		}
		if tdt.X > bounds.max_x || bounds.max_x == 0.0 {
			bounds.max_x = tdt.X
		}

		// min/max Y
		if tdt.Y < bounds.min_y || bounds.min_y == 0.0 {
			bounds.min_y = tdt.Y
		}
		if tdt.Y > bounds.max_y || bounds.max_y == 0.0 {
			bounds.max_y = tdt.Y
		}

		// min/max Z
		if tdt.Z < bounds.min_z || bounds.min_z == 0.0 {
			bounds.min_z = tdt.Z
		}
		if tdt.Z > bounds.max_z || bounds.max_z == 0.0 {
			bounds.max_z = tdt.Z
		}
	}

	return bounds
}

func computeSimpleScores(bounds TDTBounds, tdts []TemporalDisplacementTrips) {
	for _, tdt := range tdts { 
		computeSimpleScore(bounds.max_t, bounds.min_t, tdt.T)
	}

	for _, tdt := range tdts { 
		computeSimpleScore(bounds.max_x, bounds.min_x, tdt.X)
	}

	for _, tdt := range tdts { 
		computeSimpleScore(bounds.max_y, bounds.min_y, tdt.Y)
	}

	for _, tdt := range tdts { 
		computeSimpleScore(bounds.max_z, bounds.min_z, tdt.Z)
	}
}

func computeSimpleScore(max float64, min float64, value float64) {

	incrementalUnit := (max - min) / 100
	position := value / incrementalUnit 

	fmt.Printf("incunit: %f\n", incrementalUnit)
	fmt.Printf("max: %f, min: %f, value: %f, position: %f \n", max, min, value, position)
}

// computeSpatialMinTemporalMaxScores() 

func loadTDTs() []TemporalDisplacementTrips {
	csvFile, _ := os.Open("dat_points.io")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var isFirstLine bool = true

	var tdt_arrs []TemporalDisplacementTrips

	for {
        line, err := reader.Read()
		if err == io.EOF {
            break
		}
		
		if !isFirstLine {
			ID := line[0]
			T, _ := strconv.ParseFloat(line[1], 64)
			X, _ := strconv.ParseFloat(line[2], 64)
			Y, _ := strconv.ParseFloat(line[3], 64)
			Z, _ := strconv.ParseFloat(line[4], 64)

			tdt := TemporalDisplacementTrips {
				ID: ID, T: T, X: X, Y: Y, Z: Z}
			tdt_arrs = append(tdt_arrs, tdt)
		}
		isFirstLine = false
	}
	
	return tdt_arrs
}

func main() {
	/*
		weighting logic:
			- closest spatial points 
			- farthest temporal displacement

		in other words.... 

		- imagine a cubical rectangle  Y in 1s, y in 10s, z in mils, temporal in fractional seconds
		- compute scores for each quadrant. then, apply logic over the top of the computed quadrant values...

		- A TDT - We favor spatial mins, temporal maxs (so we need to inverse the scoring for spatial attribs)
	*/
	tdts := loadTDTs()
	printTDTs(tdts)

	tdtBounds := computeTDTBoundaries(tdts)
	printTDTBounds(tdtBounds)

	computeSimpleScores(tdtBounds, tdts)
}

func printTDTs(tdts []TemporalDisplacementTrips) {
	for _, tdt := range tdts {
		fmt.Printf("TDT { ID: %v, T: %v, X: %v, Y: %v, Z: %f }\n", tdt.ID, tdt.T, tdt.X, tdt.Y, tdt.Z)
	}
}

func printTDTBounds(tdtBounds TDTBounds) {
	fmt.Printf("TDT-Bounds\n")
	fmt.Printf("\t min_t: %v\n", tdtBounds.min_t)
	fmt.Printf("\t max_t: %v\n", tdtBounds.max_t)
	fmt.Printf("\t min_x: %v\n", tdtBounds.min_x)
	fmt.Printf("\t max_x: %v\n", tdtBounds.max_x)
	fmt.Printf("\t min_y: %v\n", tdtBounds.min_y)
	fmt.Printf("\t max_y: %v\n", tdtBounds.max_y)
	fmt.Printf("\t min_z: %v\n", tdtBounds.min_z)
	fmt.Printf("\t max_z: %v\n", tdtBounds.max_z)
}