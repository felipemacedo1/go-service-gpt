<!--
# GPT Service (Go)

Este microserviço foi desenvolvido em Go com o objetivo de integrar uma interface conversacional com o modelo GPT-4o da OpenAI.

Ele faz parte da plataforma modular de assistente inteligente e é responsável por:
- Validar tokens JWT emitidos pelo backend (Java + Spring Boot)
- Repassar mensagens para a API da OpenAI (GPT-4o)
- Retornar a resposta ao usuário final

## 🚀 Tecnologias
- Go 1.20+
- Echo Framework (router)
- JWT Middleware
- HTTP Client
- OpenAI API

## 📦 Funcionalidades
- Endpoint `/chat`: recebe uma mensagem, valida o token, consulta a OpenAI e retorna a resposta.
- Validação de token JWT com backend Java
- Timeout configurável para requisições
- Controle de logs e erros

## 🔐 Segurança
- Autenticação via JWT obrigatório
- Verificação ativa com o backend Java
- Suporte futuro para limites de uso por plano

## 📁 Estrutura (sugestão)
-->
