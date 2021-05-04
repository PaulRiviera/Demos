package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var (
	ordersInMemory []Orders
)

func init() {
	ordersInMemory = []Orders{}
}

func OrderStatusHandler(w http.ResponseWriter, r *http.Request) {
	var req StatusRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	ordersWithIDRequested := []Orders{}
	deliveryDate := "soon"

	for _, o := range ordersInMemory {
		if o.ShipmentNumber == req.ID {
			ordersWithIDRequested = append(ordersWithIDRequested, o)
			deliveryDate = o.DeliveryDate
		}
	}

	response := StatusResponse{}
	response.Message = fmt.Sprintf("Order contains %v product(s) and will be delivered %v", len(ordersWithIDRequested), deliveryDate)

	b, _ := json.Marshal(response)
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PING"))
}

func main() {
	ordersInMemory = readInFile()

	fmt.Println("Ready and Listening at Port 80")
	http.HandleFunc("/ping", PingHandler)
	http.HandleFunc("/order/status", OrderStatusHandler)
	panic(http.ListenAndServe(":80", nil))
}

func readInFile() []Orders {
	file, err := os.Open("./chatbot_data.local.csv")
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	rows := []Orders{}
	for _, r := range records {
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

	return rows
}

func convertRow(d []string) *Orders {
	r := Orders{}
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

type StatusResponse struct {
	Message string `json:"message"`
}

type StatusRequest struct {
	ID string `json:"id"`
}

type Orders struct {
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
