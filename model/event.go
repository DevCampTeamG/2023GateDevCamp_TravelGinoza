package model

import (
	"bytes"
	"encoding/csv"

	"github.com/DevCampTeamG/TwoGateDevCamp2023_TravelGinoza/helper"
	"github.com/jszwec/csvutil"
)

type Event []struct {
	WeekDay int    `csv:"week_day"`
	Content string `csv:"content"`
}

func (e *Event) MarshallCsv() {
	rows := helper.ReadCSVAll("../GinozaEvent.csv")
	b := &bytes.Buffer{}

	err := csv.NewWriter(b).WriteAll(rows)
	if err != nil {
		panic(err)
	}
	csvutil.Unmarshal(b.Bytes(), e)
}
