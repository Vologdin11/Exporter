# stp-exporter

### Необходимые переменные окружения для работы
Задаются в docker-compose.yml<br>
**URL** - адрес сервера, по которому находится база данных<br>
**DB** - название базы данных<br>
**LOGIN** - логин для авторизации<br>
**PASSWORD** - пароль для авторизации<br>

**EXAMPLE**
```
environment:
      - URL=http://localhost:8080
      - DB=DB_NAME
      - LOGIN=YOUR_LOGIN
      - PASSWORD=YOUR_PASSWORD
      - AUTHORIZATION=testexample
```

### Сборка Docker Images
```
sudo make build
```

### Добавление/удаление таблиц
Для добавления/удаление таблиц нужно изменить config.yml
1. Добавить/удалить название таблицы
```
tables:
  - name: name_your_table
```
2. Добавить/удалить индекс по которому находится значение метрики в получаемом JSON ответе, на запрос к таблице<br>

**JSON**
```
"values": [
            [
                "0",
                "example",
                "text",
                1, //значение метрики приходит под 3-им индексом
            ]
```
**config.yml**
```
tables:
  - name: name_your_table
    value_index: 3
```
3. Добавить/удалить индексы по которым находятся Labels метрики, которые вы хотите отображать, в получаемом JSON ответе, на запрос к таблице<br>

**JSON**
```
"values": [
            [
                "0",        //1-й label индекс 0
                "example",  //2-й label индекс 1
                "text",     //3-й label индекс 2
                1,
            ]
```
**config.yml**
```
tables:
  - name: name_your_table
    value_index: 3
    label_indexes: [0,1,2]
```
