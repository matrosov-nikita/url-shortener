Должны быть установлены `docker` и `docker-compose`

Запуск

    1. Создать docker сеть: `docker network create my-network`
    2. Запустить сервисы для бд: `docker-compose -f docker-compose.db.yml up -d`
    3. Убедиться что все сервисы для бд успешно стартовавли: `docker -f docker-compose.db.yml ps` (поле State должно быть Up)
    4. Запустить микросервисы: `docker-compose -f docker-compose.services.yml up -d`

Использование CLI клиента:
    1. `cd urlsho-cli`
    2. `go build`
    3. `./urlsho-cli --apiURL=<HOST:PORT> --encode(or --decode)=<URL>

    -apiURL string
    	API URL в формате host:port  (default "http://localhost:8080")
    -decode string
    	URL для декодирования
    -encode string
    	URL для кодирования










