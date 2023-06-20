package constants

type UserRole string

const (
	DefaultHashCost = 14

	DonatorRole   UserRole = "donator"
	CollectorRole UserRole = "collector"
)

var (
	// has a format of:
	// 	 key   -> points
	//   value -> real money
	RedeemExchangeRate = map[int64]int64{
		20_000:  10_000,
		50_000:  30_000,
		100_000: 70_000,
		200_000: 130_000,
	}
)
