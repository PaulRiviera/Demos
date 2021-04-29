package main

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strings"
)

type Row struct {
	ShipmentNumber    string `json:"sd_ship_no,omitempty"`
	ShipmentType      string `json:"sh_ship_typ_cd,omitempty"`
	ShipToName        string `json:"sh_ship_to_nam,omitempty"`
	CustomerProductID string `json:"sd_cust_prod_id,omitempty"`
	ExtnDesc          string `json:"sd_extn_desc,omitempty"`
	InvoiceNumber     string `json:"sd_invce_no,omitempty"`
	DeliveryDate      string `json:"sd_prm_dlvry_dte,omitempty"`
	OrderDate         string `json:"sh_ord_dte,omitempty"`
	ShipDate          string `json:"sh_ship_dte,omitempty"`
	ShipCode          string `json:"sh_ship_cd,omitempty"`
	Product           string `json:"sd_prod_desc,omitempty"`
	ProductID         string `json:"sd_prod_id,omitempty"`
}

func main() {

	file, err := os.Open("./chatbot_data.local.csv")
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	rows := []Row{}
	for _, r := range records[1:10] {

		newRow := convertRow(r)
		rows = append(rows, *newRow)
	}

	jsonFile, err := os.Create("data.json")
	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(rows)
	if err != nil {
		panic(err)
	}

	_, err = jsonFile.Write(b)
	if err != nil {
		panic(err)
	}
}

func convertRow(d []string) *Row {
	r := Row{}
	r.ShipmentNumber = strings.Trim(d[0], " ")
	r.ShipmentType = strings.Trim(d[1], " ")
	r.ShipToName = strings.Trim(d[2], " ")
	r.CustomerProductID = strings.Trim(d[3], " ")
	r.ExtnDesc = strings.Trim(d[4], " ")
	r.InvoiceNumber = strings.Trim(d[5], " ")
	r.DeliveryDate = strings.Trim(d[6], " ")
	r.OrderDate = strings.Trim(d[7], " ")
	r.ShipDate = strings.Trim(d[8], " ")
	r.ShipCode = strings.Trim(d[9], " ")
	r.Product = strings.Trim(d[10], " ")
	r.ProductID = strings.Trim(d[11], " ")
	return &r
}
