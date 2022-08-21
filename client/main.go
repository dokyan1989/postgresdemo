package main

import (
	"encoding/json"
	"fmt"

	modelgen "github.com/dokyan1989/postgresdemo/model/gen"
)

func main() {
	order := modelgen.GenOrder()
	// pOrder, _ := PrettyStruct(order)
	// fmt.Println(pOrder)

	orderRawData := modelgen.GenOrderRawData(order)
	// pOrderRawdata, _ := PrettyStruct(orderRawData)
	// fmt.Println(pOrderRawdata)
	fmt.Println(PrettyStruct(orderRawData.OrderInfo.Items[0]))
	fmt.Println(len(orderRawData.OrderInfo.Items))

	// shipment := modelgen.GenShipment("123")
	// s, _ := PrettyStruct(shipment)
	// fmt.Println(s)

	// invoice := modelgen.GenInvoice("123", "234", 1)
	// s, _ := PrettyStruct(invoice)
	// fmt.Println(s)
}

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(val), nil
}
