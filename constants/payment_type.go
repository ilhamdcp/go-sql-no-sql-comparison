package constants

type PaymentType string

const (
	Debit      PaymentType = "DEBIT"
	CreditCard             = "CREDIT_CARD"
	Gopay                  = "GOPAY"
	ShopeePay              = "SHOPEE_PAY"
	Ovo                    = "OVO"
	Dana                   = "DANA"
	Qris                   = "QRIS"
	Cash                   = "CASH"
)
