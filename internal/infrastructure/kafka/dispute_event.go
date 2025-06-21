package kafka

type DisputeEvent struct {
	DisputeID 		string	`json:"dispute_id"`
	OrderID  		string 	`json:"order_id"`
	TraderName 		string	`json:"trader_name"`
	AmountFiat 		float64 `json:"amount_fiat"`
	Reason		 	string	`json:"reason"`
	ProofUrl 		string	`json:"proof_url"`
}