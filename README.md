# Consulta de CEP em Go: Escolhendo a API Mais Rápida com Multithreading

Este projeto realiza consultas de CEP utilizando duas APIs diferentes simultaneamente: [BrasilAPI](https://brasilapi.com.br) e [ViaCEP](https://viacep.com.br). A aplicação retorna a resposta da API que responder primeiro, descartando a mais lenta, e limita o tempo de resposta a 1 segundo.

## Índice

- [Sobre o Projeto](#sobre-o-projeto)
- [Tecnologias Utilizadas](#tecnologias-utilizadas)
- [Como Executar o Projeto](#como-executar-o-projeto)
- [Funcionamento](#funcionamento)
- [Estrutura do Código](#estrutura-do-código)
- [Contribuindo](#contribuindo)
- [Licença](#licença)
- [Contato](#contato)

## Sobre o Projeto

O objetivo deste projeto é demonstrar como implementar consultas eficientes a serviços web externos em Go, utilizando goroutines para executar requisições simultâneas e canais para capturar a resposta mais rápida.

## Tecnologias Utilizadas

- **Go** (Golang) versão 1.16 ou superior
- **APIs Consumidas**:
    - [BrasilAPI](https://brasilapi.com.br)
    - [ViaCEP](https://viacep.com.br)

## Como Executar o Projeto

1. **Pré-requisitos**:

    - Ter o Go instalado em sua máquina. Você pode baixar a versão mais recente em [https://golang.org/dl/](https://golang.org/dl/).

2. **Clonar o repositório**:

   ```bash
   git clone https://github.com/seu-usuario/nome-do-repositorio.git
