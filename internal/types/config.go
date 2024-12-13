package types

type Config struct {
	Mysql   Mysql   `json:"mysql" yaml:"host"`
	BaiLian BaiLian `json:"baiLian" yaml:"baiLian"`
}

type Mysql struct {
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	DB       string `json:"db" yaml:"db"`
	Dir      string `json:"dir" yaml:"dir"`
}

type BaiLian struct {
	Apikey string `json:"apikey" yaml:"apikey"`
}
