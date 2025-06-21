package domain

import (
	"fmt"
)

type DisputeNotification struct {
	DisputeID 		string
	OrderID  		string
	TraderID 		string
	TraderName 		string
	AmountFiat 		float64
	Reason 			string
	ProofUrl 		string
}

func (n DisputeNotification) String() string {
	return fmt.Sprintf("Открыт диспут %s\nПричина:%s\nСделка: %s\nТрейдер:%s\nСумма:%f\n%s", n.DisputeID, n.Reason, n.OrderID, n.TraderName, n.AmountFiat, n.ProofUrl)
}