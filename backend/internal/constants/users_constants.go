package constants

type UserRole string

const (
	DefaultHashCost = 14

	DonatorRole   UserRole = "donator"
	CollectorRole UserRole = "collector"
)