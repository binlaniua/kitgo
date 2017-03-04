package db


//-------------------------------------
//
//
//
//-------------------------------------
type DataBaseConfig struct {
	Alias    string `json:"alias"`
	Ip       string `json:"ip"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DB       string `json:"db"`
}
