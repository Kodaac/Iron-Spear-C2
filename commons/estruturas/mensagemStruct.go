package estruturas

type Mensagem struct {
	AgentId       string
	AgentHostname string
	AgentCWD      string
	Comandos      []Comando
}
