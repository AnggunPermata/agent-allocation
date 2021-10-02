package constant

const SECRET_JWT = "AgentAllocation"

var Configuration = map[string]string{
	//"ConnectionString": "root:Teacup21@tcp(localhost:3306)/qiscus?charset=utf8&parseTime=True&loc=Local",
	"ConnectionString": "root:12345@tcp(172.17.0.1:3307)/qiscus?charset=utf8&parseTime=True&loc=Local",
	"PORT":             "8080",
}

const Default_Welcome = "Hello, Welcome to the store. How can I Help you?"
