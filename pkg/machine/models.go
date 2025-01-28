package machine

type Balance struct {
	OneRupeeCoin int64
	TenRupeeBill int64
}

func NewBalance() *Balance {
	return &Balance{
		OneRupeeCoin: 0,
		TenRupeeBill: 0,
	}
}
