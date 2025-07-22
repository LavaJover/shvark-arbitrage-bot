package domain

import (
	"fmt"
)

type DisputeNotification struct {
	DisputeID 			string 	
	OrderID 			string	
	TraderID 			string  
	ProofUrl 			string	
	Reason 				string	
	Status 				string	
	OrderAmountFiat 	float64 
	DisputeAmountFiat 	float64 
	BankName 			string  
	Phone 				string  
	CardNumber 			string  
	Owner 				string  
}

func (n DisputeNotification) String() string {
	if n.Phone != "" {
		return fmt.Sprintf("Диспут %s по сделке %s\nСтатус: %s\nРеквизит: %s / %s / %s\nПричина:%s\nСумма сделки: %f RUB\nСумма диспута: %f RUB", n.DisputeID, n.OrderID, n.Status, n.Phone, n.BankName, n.Owner, n.Reason, n.OrderAmountFiat, n.DisputeAmountFiat)
	}
	return fmt.Sprintf("Диспут %s по сделке %s\nСтатус: %s\nРеквизит: %s / %s / %s\nПричина:%s\nСумма сделки: %f RUB\nСумма диспута: %f RUB", n.DisputeID, n.OrderID, n.Status, n.CardNumber, n.BankName, n.Owner, n.Reason, n.OrderAmountFiat, n.DisputeAmountFiat)
}