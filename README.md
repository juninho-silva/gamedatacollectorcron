# 🕹️ GameDataCollectorCron

Este projeto é um serviço em Go responsável por coletar e registrar dados de jogos periodicamente através de jobs agendados com `robfig/cron`

---

## 🚀 Funcionalidades

- Agendamento de tarefas com [`robfig/cron`](https://github.com/robfig/cron)
- Coleta de dados de APIs de jogos 
- Suporte ao uso de variáveis de ambiente via `.env`

---

## 📦 Tecnologias

- Go 1.24+
- robfig/cron v3
- godotenv
- MongoDB (opcional, para persistência de dados)

---

## ⚙️ Instalação

1. Clone o repositório:

```bash
git clone https://github.com/juninho-silva/gamedatacollectorcron.git
cd gamedatacollectorcron
```
2. Instalar os pacotes:

```bash
go mod tidy
```

3. Crie um arquivo .env:

```bash
APIKEY="fa80ea3611ef4484a6475bccfb634ec8"
URL="https://api.rawg.io/api/games"
MONGOURI="mongodb://localhost:27017/"
```

4. Executando o projeto:

```bash
go run main.go
```