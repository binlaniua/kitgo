package rc

//-------------------------------------
//
//
//
//-------------------------------------
type RedisConfig struct {
	Ip        string `json:"ip"`
	Port      string `json:"port"`
	Password  string `json:"password"`
	MaxIdle   int    `json:"maxIdle"`
	MaxActive int    `json:"maxActive"`
}
