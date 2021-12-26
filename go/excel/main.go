package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type City struct {
	Address  string
	Lat, Lng float32
}

func main() {
	f, err := excelize.OpenFile("worldcities.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	// cell, err := f.GetCellValue("Sheet1", "B2")
	// if err != nil {
	//     fmt.Println(err)
	//     return
	// }
	// fmt.Println(cell)
	// Get all the rows in the Sheet1.
	list := []City{}

	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		count := 0
		city := City{}
		for _, colCell := range row {
			if count > 4 {
				list = append(list, city)
				fmt.Println(city)
				break
			}

			switch count {
			case 0:
				city.Address = colCell
			case 2:
				value, err := strconv.ParseFloat(colCell, 32)
				if err != nil {
					// do something sensible
				}
				city.Lat = float32(value)
			case 3:
				value, err := strconv.ParseFloat(colCell, 32)
				if err != nil {
					// do something sensible
				}
				city.Lng = float32(value)
			case 4:
				city.Address = city.Address + ", " + colCell
			}

			count++
		}
		fmt.Println()
	}

	file, _ := json.MarshalIndent(list, "", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)
}
