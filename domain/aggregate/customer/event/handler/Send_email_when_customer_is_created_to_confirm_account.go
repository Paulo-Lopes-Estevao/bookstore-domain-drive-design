package handler

import (
	"bookstore/domain/event"
	"fmt"
)

func SendEmailWhenCustomerIsCreatedToConfirmAccount() func(*event.EventType) {
	fmt.Println("Sending email to .....")
	return func(m *event.EventType) {

		fmt.Println(m.EventData)
	}
}

//type name struct {
//	Name string
//}
//productOutput := name{}
//		productJson, _ := json.Marshal(m.EventData)
//		json.Unmarshal(productJson, &productOutput)
//		fmt.Println(productOutput.Name)
