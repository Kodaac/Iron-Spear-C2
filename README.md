# Iron Spear - C2 em Go

Este repositório contém a implementação de um servidor de Comando e Controle (C2) desenvolvido em **Go** como parte de um curso de evasão de defesas. O objetivo deste projeto é puramente educacional, com foco no aprendizado de técnicas e estratégias defensivas e ofensivas em segurança da informação.

## 🚨 AVISO LEGAL

**Este projeto foi desenvolvido para fins educacionais e de pesquisa em segurança da informação.**

- **Não deve ser utilizado para atividades ilegais, invasões, ou qualquer outro tipo de uso malicioso.**
- O autor não se responsabiliza por qualquer dano ou uso indevido deste código.
- Ao utilizar este repositório, você concorda em seguir todas as leis e regulamentos aplicáveis.

Caso tenha dúvidas sobre o uso permitido, consulte as leis de sua jurisdição antes de utilizar este código.

## 📚 Objetivo

O projeto foi criado com os seguintes propósitos:
1. Aprender a desenvolver e entender a arquitetura de um servidor de Comando e Controle.
2. Compreender como técnicas de evasão são aplicadas em ambientes controlados.
3. Praticar a identificação e mitigação de ameaças em um ambiente defensivo.

## 🛠️ Funcionalidades

- Comunicação entre agentes e o servidor.
- Suporte a múltiplos clientes conectados simultaneamente.
- Funções de controle remoto, incluindo execução de comandos.
- Técnicas básicas de evasão para evitar detecção.

## 📋 Pré-requisitos

- **Go**: Versão 1.19 ou superior instalada.
- Ambiente Windows por enquanto.

## 🚀 Configuração

1. Clone o repositório:
   git clone https://github.com/Kodaac/Iron-Spear-C2.git
   cd seu-repositorio

2. Compile o projeto:
    go build -o i2s.go

3. Inicie o servidor:
    go run i2s.go
    ou i2s.exe

## 📖 Como funciona
O servidor aguarda conexões dos agentes.
Uma vez conectado, o agente pode receber comandos específicos enviados pelo servidor.
Todas as interações são logadas para análise e estudo.

## 💻 Comandos disponíveis

### Comandos do Agente
```bash
ls           - Lista os arquivos no diretório atual.
pwd          - Exibe o diretório atual.
cd [path]    - Altera o diretório de trabalho para o especificado.
whoami       - Retorna o usuário atual do sistema.
ps           - Lista os processos em execução na máquina.
send         - Salva no disco um arquivo enviado pelo servidor.
get          - Envia ao servidor um arquivo do sistema de arquivos local.
sleep [s]    - Define o intervalo de tempo (em segundos) para a próxima execução.
[Outros]     - Executa o comando diretamente no shell do agente.
```
### Comandos do Servidor
```bash
show         - Exibe informações sobre os agentes conectados.
sleep [s]    - Define o intervalo de espera do agente selecionado (em segundos).
select [id]  - Seleciona um agente para interagir diretamente.
send [file]  - Envia um arquivo para o agente selecionado.
get [file]   - Solicita um arquivo do agente selecionado.
[Outros]     - Envia o comando para ser executado no shell do agente selecionado.
```
⚠️ Nota: Antes de usar comandos específicos (sleep, send, get, etc.), certifique-se de selecionar um agente com select [id].

## ⚠️ Disclaimer de Responsabilidade
Este repositório foi criado apenas para fins acadêmicos e não deve ser usado para comprometer sistemas sem autorização explícita. Usar este software de maneira inadequada pode violar leis locais, estaduais ou internacionais.