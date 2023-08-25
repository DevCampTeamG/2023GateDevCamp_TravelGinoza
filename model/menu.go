package model

import (
	"bytes"
	"encoding/csv"
	"log"

	"github.com/DevCampTeamG/TwoGateDevCamp2023_TravelGinoza/helper"
	"github.com/jszwec/csvutil"
)

type Menu []struct {
	Num        int    `csv:"num"`
	Name       string `csv:"name"`
	ImagePlace string `csv:"image_place"`
}

func (m *Menu) MarshallCsv() {
	rows := helper.ReadCSVAll("../menu.csv")
	b := &bytes.Buffer{}

	err := csv.NewWriter(b).WriteAll(rows)
	if err != nil {
		panic(err)
	}
	csvutil.Unmarshal(b.Bytes(), m)
	log.Println(m)
}
