package processtransaction

type TransactionDtoInput struct {
	ID                        string  `json: "id"`
	AccountID                 string  `json: "account_id"`
	CreditCardNumber          string  `json: "credit_card_number"`
	CreditCardName            string  `json: "credit_card_name"`
	CreditCardExpirationMonth int     `json: "credit_expiration_month"`
	CreditCardExpirationYear  int     `json: "credit_expiration_year"`
	CreditCardCVV             int     `json: "credit_cvv"`
	Amount                    float64 `json: "amount"`
}

type TransactionDtoOutput struct {
	ID           string `json: "id"`
	Status       string `json: "status"`
	ErrorMessage string `json: "error_message"`
}
