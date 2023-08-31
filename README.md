# Avito-Task
## Задание выполнил: Корешков Даниил
____
## Описание задания
Нужно реализовать сервис, хранящий пользователя и сегменты, в которых он состоит 
(создание, изменение, удаление сегментов, а также добавление и удаление пользователей в сегмент)

### Что было реализовано
+ Метод создания сегмента
+ Метод удаления сегмента
+ Метод добавления пользователя в сегмент
+ Метод получения активных сегментов пользователя
+ Дополнительно:
  + Реализация попадания/пользователя из сегмента с возможностью получения отчета по пользователю 
  за определенный период
  + Добавление пользователя в сегмент на ограниченный срок
  + Автоматическое распределение пользователей по сегментам
  + Swagger файл для API
  
---

## Начало работы

### Запуск dev-среды

#### Вариант 1
 При помощи `make` выполнить команду `build-containers`. <br>
 После установки docker-контейнеров выполнить команду `add-tables`. <br>

##### Код запуска dev-среды:
```
    make build-containers
    make add-tables
```
#### Вариант 2
Выаолнить последовательно команды:

##### Код запуска dev-среды:
```
 docker-compose up -d --build avito-app
 migrate -database postgres://postgres:task@localhost:5436/?sslmode=disable -path ./schema up
```

---

## Описание методов

### Метод *getAllSegments*
Возвращает список всех существующих сегментов из базы данных

#### Аргументы на вход
Нет
#### Возвращает
Срез объектов Segment

#### Пример работы:

##### Запрос
`http://localhost:8080/api/segments/all`
##### Ответ
```
{
    "data": [
        {
            "id": 12,
            "name": "VK",
            "entirety": {
                "Int16": 0,
                "Valid": false
            }
        },
        {
            "id": 15,
            "name": "AVITO_TEST_VIDEO",
            "entirety": {
                "Int16": 0,
                "Valid": false
            }
        },
        {
            "id": 17,
            "name": "AVITO_TEST_VOICE",
            "entirety": {
                "Int16": 0,
                "Valid": false
            }
        }
    ]
}
```
---
### Метод *getSegmentById*
Возвращает сегмент из базы данных по заданному Id

#### Аргументы на вход

`/api/segments/get/:id`

+ id - идентификатор сегмента в Базе данных

#### Возвращает
Объект Segment

#### Пример работы:

##### Запрос
`http://localhost:8080/api/segments/get/12`
##### Ответ
```
{
    "id": 12,
    "name": "VK",
    "entirety": {
        "Int16": 0,
        "Valid": false
    }
}
```
---
### Метод *createSegment*
Создает новый объект сегмента в Базе данных

#### Аргументы на вход
`/api/segments/add` <br>

**JSON**
```
{
  "id": Value // integer, not required
  "name": Value // string, required
  "entirety": Value // integer, not required
}
```

#### Возвращает
Результат выполнения добавления:
+ true - если все выполненно успешно
+ false, error - если найдена ошибка

#### Пример работы:

##### Запрос
`http://localhost:8080/api/segments/add`
```
{
    "name": "TEST_AVITO_SEGMENT"
}
```
##### Ответ
```
{
    "isInsert": true
}
```
---
### Метод *deleteSegmentByName*
Удаляет выбранный сегмент по его имени

#### Аргументы на вход
`"api/segments/delete/:segmentName"`
+ segmentName - название сегмента (string)
#### Возвращает
+ true - если все выполненно успешно
+ false, error - если найдена ошибка

#### Пример работы:

##### Запрос
`http://localhost:8080/api/segments/delete/TEST_AVITO_SEGMENT`
##### Ответ
```
{
    "isDelete": true
}
```
---
### Метод *getActiveUserSegments*
Возвращает список всех активных сегментов из базы данных, закрепленных за пользователями

#### Аргументы на вход
`"/api/users/active-segments/:id"`
+ id - Идентификатор пользователя (int)
#### Возвращает
Срез объектов string

#### Пример работы:

##### Запрос

`http://localhost:8080/api/users/active-segments/10`

##### Ответ

```
{
    "activeSegments": [
        "AVITO_SEGMENT_VOICE",
        "AVITO_SEGMENT_VIDEO",
        "AVITO_SEGMENT_30"
    ]
}
```

---
### Метод *getHistory*
Генерирует CSV-файл с историей действий пользователя за определенный период и ссылку для скачивания файла

#### Аргументы на вход
`"api/users/history/:id"`
+ id - идентификатор пользователя (int)

**JSON**
```
{
  "month": Value // integer, required
  "year": Value // integer, required
}
```

#### Возвращает
Сгенерированную ссылку для скачивания Link(string)

#### Пример работы:

##### Запрос

`http://localhost:8080/api/users/history/10`

```
{
    "year": 2023,
    "month": 8
}
```
##### Ответ

```
{
    "downloadLink": "localhost:8080/service/download/JrRbBvzc"
}
```

---
### Метод *changeUserSegments*
Добавляет/удаляет сегменту у пользователя

#### Аргументы на вход
`"/api/users/change-segments/:id"`
+ id - идентификатор пользователя

**JSON**
```
{
  "add": Value // []string, required
  "delete": Value // []string, required
}
```

#### Возвращает
Результат работы:
+ true, int, int - успешное добавление/удаление (дополнительно выводи количество добавленных/удаленных сегментов)
+ false, error - ошибка в ходе выполнения метода

#### Пример работы:

##### Запрос

`http://localhost:8080/api/users/change-segments/10`

```
{
    "delete": ["AVITO_SEGMENT_VIDEO", "AVITO_SEGMENT_VOICE"],
    "add": ["AVITO_SEGMENT_30"]
}
```

##### Ответ
```
{
    "added": 1,
    "isSegmentListChanged": true,
    "removed": 2
}
```
---
### Метод *addExpiredSegment*
Добавляет сегмент на ограниченное время

#### Аргументы на вход
`"api/users/expired-segments/:id"`
+ id - идентификатор пользователя (id)

**JSON**
```
{
  "segment": Value // string, required
  "time": Value // int, required
}
```
+ segment - имя сегмента
+ time - время жизни (в часах)

#### Возвращает
Результат работы:
+ true, - успешное добавление
+ false, error - ошибка в ходе выполнения метода

#### Пример работы:

##### Запрос
`http://localhost:8080/api/users/expired-segments/10`

```
{
    "segment": "AVITO_SEGMENT_AUDIO",
    "time": 2
}
```
##### Ответ
```
{
    "isSegmentListChanged": true
}
```

---
### Метод *addUser*
Добавляет пользователя в Базу данных

#### Аргументы на вход
`"api/users/add"`


**JSON**
```
{
  "login": Value // string, required
  "password": Value // string, required
}
```
+ login - Логин пользователя
+ password - Пароль пользователя (пароль в базе хранится в хешированном виде)

#### Возвращает
true - успешное добавление пользователя
false, error - ошибка во время добавления 

#### Пример работы:

##### Запрос

`"api/users/add"`

```
{
    "login": "avito",
    "password": "avito"
}
```

##### Ответ
```
{
    "userAdded": true
}
```

---

### Метод *downloadFile*
Скачивание файла по ссылке

#### Аргументы на вход
`"service/download/:id"`
+ id - Идентификатор файла скачивания (string)
#### Возвращает
CSV-файл

#### Пример работы:

##### Запрос
`http://localhost:8080/service/download/nirUwP3I`
##### Ответ
```
id,user_id,segment_name,action,time
15,10,AVITO_TEST_1,added,2023-08-31T10:36:38.185193Z
14,10,AVITO_TEST_2,2023-08-31T10:29:50.57482Z
```
---

