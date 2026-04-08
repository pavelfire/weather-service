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

geocoding-api.open-meteo.com/v1/search?name=Moscow&count=1&language=ru&format=json

api.open-meteo.com/v1/forecast?latitude=55.7558&longitude=37.6176&current=temperature_2m,apparent_temperature,relative_humidity_2m,weather_code,surface_pressure,wind_speed_10m,wind_direction_10m&timezone=auto

https://api.open-meteo.com/v1/forecast?latitude=55.7558&longitude=37.6176&current=temperature_2m

https://wttr.in/55.72,52.41?format=j1