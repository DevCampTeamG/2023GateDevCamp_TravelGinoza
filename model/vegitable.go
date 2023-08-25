package model

import (
	"bytes"
	"encoding/csv"

	"github.com/DevCampTeamG/TwoGateDevCamp2023_TravelGinoza/helper"
	"github.com/jszwec/csvutil"
)

type Vegitable []struct {
	ID         int    `csv:"id"`
	Name       string `csv:"name"`
	Content    string `csv:"content"`
	ImagePlace string `csv:"image_place"`
}

func (v *Vegitable) MarshallCsv() {
	rows := helper.ReadCSVAll("../vegitable.csv")
	//fmt.Println(rows)
	//v2 := []vegitable{}

	b := &bytes.Buffer{}

	err := csv.NewWriter(b).WriteAll(rows)
	if err != nil {
		panic(err)
	}
	csvutil.Unmarshal(b.Bytes(), v)

}
