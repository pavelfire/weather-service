go mod init github.com/pavelfire/weather-service
git init
git remote add origin https://github.com/pavelfire/weather-service.git

1. Сделать http эндпоинт, который присылает в path параметре имя города и температуру в нём
2. Реализовать cron, который будет ходить в стороннее API 
3. Подключить PostgreSQL и сохранять в него данные о погоде полученные cron -ом
4. Подключить Redis и реализовать кэширование данных о погоде
5. Вывести все параметры в конфиг

go run ./cmd/server/main.go
go get -u github.com/go-chi/chi/v5
localhost:3000

go get github.com/go-co-op/gocron/v2

