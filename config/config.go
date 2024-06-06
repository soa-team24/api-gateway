package config

type Config struct {
	Address                   string
	BlogServiceAddress        string
	FollowerServiceAddress    string
	StakeholderServiceAddress string
	TourServiceAddress        string
}

func GetConfig() Config {
	return Config{
		BlogServiceAddress:        "8000",
		FollowerServiceAddress:    "8000",
		StakeholderServiceAddress: "8000",
		TourServiceAddress:        "8000",
		Address:                   "8000",
	}
}
