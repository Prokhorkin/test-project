 1. Запуск приложения - go run main.go
    - Приложение запускается на локальном хосте, порт 5505 
 2. Http API приложения:
   - Метод Get
     - /ping - запрос доступности приложения, возвращает код 200
        - Пример запроса из консоли - curl localhost:5505/ping
   - Метод Post
     - /putUserData - отправка данных пользователя в JSON формате и вывод в разобранном виде в логах
        - Формат JSON:
            - {"FIO":"Иванов Иван Иванович","age":26,"passportData":{"seria":"0950","number":"0759679"}}
        - Пример запроса из консоли - curl localhost:5505/putUserData -d '{"FIO":"Иванов Иван Иванович","age":26,"passportData":{"seria":"0950","number":"0759679"}}'
 3. Остановка приложения - сочетания клавиш Ctrl+C    
   