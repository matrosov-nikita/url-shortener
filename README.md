Должны быть установлены `docker`, `docker-compose`, `go`

### Запуск:
Запустить билд:
```
$ go build
```     
Создать docker сеть:
```
$ docker network create my-network
```  
Запустить сервисы для бд:
```
$ docker-compose -f docker-compose.db.yml up -d
```
Убедиться что все сервисы для бд успешно стартовавли (поле State должно быть Up):  
```
$ docker-compose -f docker-compose.db.yml ps
```
Запустить микросервисы
```
$ docker-compose -f docker-compose.services.yml up -d
```  

### Использование CLI клиента:  
```
    1. cd urlsho-cli  
    2. go build  
    3. ./urlsho-cli --apiURL=<HOST:PORT> --encode(or --decode)=<URL>
```  

    -apiURL string
    	API URL в формате host:port  (default "http://localhost:8080")
    -decode string
    	URL для декодирования
    -encode string
    	URL для кодирования 

### Использование Web-клиента:   
См. https://github.com/matrosov-nikita/url-shortener-web-client  

### API

Endpoint: URLHandler.Encode  

```
// Request содержит URL для кодирования
Request: {
	url string 
}

// Response содержит сокращенную ссылку
Response: {
	short_url string
}
```


Endpoint: URLHandler.Decode  

```
// Request содержит сокращенную ссылку для декодирования
Request: {
    url string
}

// Response содержит оригинальный URL
Response: {
	origin_url string
}
```









