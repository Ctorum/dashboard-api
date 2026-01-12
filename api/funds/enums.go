package funds

type Asset string
type Category string
type Originator string
type Debtor string

const (
	TradeReceivable Asset = "TRADE_RECEIVABLE"
	TaxInvoice      Asset = "TAX_INVOICE"
)

const (
	Green          Category = "GREEN"
	Agribusiness   Category = "AGRIBUSINESS"
	Securitization Category = "SECURITIZATION"
)

const (
	SingleOriginator   Originator = "SINGLE"
	MultipleOriginator Originator = "MULTIPLE"
)

const (
	SingleDebtor   Debtor = "SINGLE"
	MultipleDebtor Debtor = "MULTIPLE"
)
