package setting

type DbConfig struct {
	DbType    string
	DbName    string
	Host      string
	Username  string
	Pwd       string
	Charset   string
	ParseTime bool
}
type ContractConfig struct {
	ContractAddress string `json:"contractAddress"`
	ContractAbi     string `json:"contractAbi"`
	URL             string `json:"URL"`
}
