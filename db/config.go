package db


//-------------------------------------
//
//
//
//-------------------------------------
type DataBaseConfig struct {
	Alias    string `json:"mysql.alias"`
	Ip       string `json:"mysql.ip"`
	Port     string `json:"mysql.port"`
	User     string `json:"mysql.user"`
	Password string `json:"mysql.password"`
	DB       string `json:"mysql.db"`
}
