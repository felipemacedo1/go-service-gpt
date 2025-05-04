
# GPT Service (Go)

Este microserviço foi desenvolvido em Go com o objetivo de integrar uma interface conversacional com o modelo GPT-4o da OpenAI.

Ele faz parte da plataforma modular de assistente inteligente e é responsável por:
- Validar tokens JWT emitidos pelo backend (Java + Spring Boot)
- Repassar mensagens para a API da OpenAI (GPT-4o)
- Retornar a resposta ao usuário final

---

## 🚀 Tecnologias
- Go 1.20+
- Echo Framework (router)
- JWT Middleware
- HTTP Client
- OpenAI API

---

## 📦 Funcionalidades
- Endpoint `/chat`: recebe uma mensagem, valida o token, consulta a OpenAI e retorna a resposta.
- Validação de token JWT com backend Java
- Timeout configurável para requisições
- Controle de logs e erros
- Modularização para fácil escalabilidade

---

## 🔐 Segurança
- Autenticação via JWT obrigatório
- Verificação ativa com o backend Java (validação do token em cada requisição)
- Suporte futuro para:
  - Limites de uso por plano
  - Rate limiting
  - Auditoria por usuário

---

## 📁 Estrutura do Projeto (sugestão)

```
gpt-service-go/
├── main.go
├── routes/
│   └── chat.go
├── controller/
│   └── chat_controller.go
├── service/
│   └── openai_service.go
├── middleware/
│   └── auth.go
├── config/
│   └── env.go
├── model/
│   └── request.go
├── utils/
│   └── logger.go
├── go.mod
├── Dockerfile
└── README.md
```

---

## ⚙️ Variáveis de Ambiente

Você deve criar um arquivo `.env` na raiz com as seguintes variáveis:

```env
PORT=8080
OPENAI_API_KEY=your_openai_api_key
AUTH_BACKEND_URL=http://backend-java:8081/auth/validate
REQUEST_TIMEOUT=10
```

---

## 🐳 Docker

Para rodar o serviço usando Docker, siga os passos abaixo:

1. **Construa a imagem Docker:**
```bash
docker build -t gpt-service-go .
```

2. **Rode o container:**
```bash
docker run -p 8080:8080 --env-file .env gpt-service-go
```

3. **Com Docker Compose (opcional):**
```yaml
# docker-compose.yml
version: '3.8'

services:
  gpt-service:
    build: .
    ports:
      - "8080:8080"
    env_file:
      - .env
```

Execute com:
```bash
docker-compose up --build
```

---

## ✅ Futuras Melhorias
- Armazenamento de histórico de conversas em banco de dados
- Integração com sistema de planos e pagamentos (Stripe)
- Suporte a arquivos e imagens (usando OpenAI Assistants ou Azure)
- Cache de resposta para prompts comuns

---

## 👨‍💻 Autor
Felipe Macedo – [GitHub](https://github.com/felipemacedo-dev) | Projeto Open Source Modular
