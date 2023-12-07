package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Cars struct {
	cars []CarTypes
}

type CarTypes struct {
	Manufacturer        string
	Model               string
	Sales_in_thousands  float64
	__year_resale_value float64
	Vehicle_type        string
	Price_in_thousands  float64
	Engine_size         float64
	Horsepower          uint16
	Wheelbase           float64
	Width               float64
	Length              float64
	Curb_weight         float64
	Fuel_capacity       float64
	Fuel_efficiency     uint16
	Latest_Launch       string
	Power_perf_factor   float64
}

func main() {

	carSales, carSalesErr := os.Open("./assets/Car_sales.csv")
	if carSalesErr != nil {
		fmt.Println("ERROR: File error")
		panic(carSalesErr)
	}
	defer carSales.Close()

	convertCars := ReadCsv(carSales)
	carList := Cars{}

	for _, car := range convertCars {
		carList.InsertHeap(car)
	}
	carList.DisplayCarList()
}

func NewCars(carArr []CarTypes) Cars {
	return Cars{
		cars: carArr,
	}
}

func (c *Cars) DisplayCarList() {
	for _, car := range c.cars {
		fmt.Println(car.Price_in_thousands)
	}
}

func ReadCsv(cs *os.File) []CarTypes {
	csvReader := csv.NewReader(cs)
	data, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	c := []CarTypes{}
	for di, d := range data {
		car := CarTypes{}
		if len(d) < 16 {
			fmt.Printf("Warning index %d is missing values.", di)
			continue
		}

		car.Manufacturer = d[0]
		car.Model = d[1]
		car.Sales_in_thousands = StringToFloat(d[2])
		car.__year_resale_value = StringToFloat(d[3])
		car.Vehicle_type = d[4]
		car.Price_in_thousands = StringToFloat(d[5])
		car.Engine_size = StringToFloat(d[6])
		car.Horsepower = StringToInt(d[7])
		car.Wheelbase = StringToFloat(d[8])
		car.Width = StringToFloat(d[9])
		car.Length = StringToFloat(d[10])
		car.Curb_weight = StringToFloat(d[11])
		car.Fuel_capacity = StringToFloat(d[12])
		car.Fuel_efficiency = StringToInt(d[13])
		car.Latest_Launch = d[14]
		car.Power_perf_factor = StringToFloat(d[15])

		c = append(c, car)
	}

	return c
}

func StringToFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println("cant parse float", s)
	}

	return f
}

func StringToInt(s string) uint16 {
	ui, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		fmt.Println("cant parse uint", s)
	}

	return uint16(ui)
}
