# Приложение для создания и применения промокодов для игры

 1. [Описание решения](#описание-решения)
	 - [Структура проекта](#структура-проекта)
	 - [Описание сервисов](#описание-сервисов)
		 - [Admin](#admin)
			 - /admin/promocodes/
			 - /admin/promocodes/promocode/
			 - /admin/rewards/
			 - /admin/rewards/reward/
		- [Promocodes](#promocodes)
			- /promocodes/promocode/use
		- [UI](#ui)

 2. [Инструкция для запуска](#инструкция-для-запуска)


# Описание решения

Данный проект представляет собой приложение для создания и применения промокодов для игры (и не только) и состоит из нескольких частей:
- административный сервис
- сервис, выполняющий серверную роль для пользователя
- сервис, сервис выполняющий клиентскую роль для пользователя
- база данных


<img width="1198" alt="image" src="https://github.com/user-attachments/assets/a7fa88d1-f42c-4992-b1db-6492d6a10745">



## Структура проекта

Приложение может запускаться с использованием docker-compose. Каждый сервис находится в своей директории в корне проекта,  каждый запускается в отдельном Docker контейнере и представляет из себя самостоятельную единицу. 
Сервисы взаимодействуют посредством HTTP запросов. 
 
## Описание сервисов

### Promocodes
Сервис получает запросы от клиента и реализует применение промокодов.

Промокод может быть применен игроком только один раз. Также у промокода может быть ограничение по дате применения и есть ограничение на общее число применений. 
Применение каждого промокода сопровождается с получением награды `reward`. 

#### POST /promocodes/promocode/use/
Позволяет применить промокод.

Параметры  - **Body**: 

```
{
	"promocode":"GET_FREE_DIAMONDS24",
	"user_id": 4,
}
```

В случае удачного применения промокода возвращается `reward`:  

```
{
    "id": 2,
    "title": "GET_DIAMONDS_7_REWARD",
    "description": "Добавляет 7 алмазов игроку"
}
```

Ответ в случае, если промокод не был применен:  

```
{
    "message": "promocode has already been used",
    "status": 0
}
```



### Admin
Сервис представляет собой административную часть и позволяет создавать промокоды и награды и присваивать промокодам награды. 

***Запросы для работы с промокодами:***
#### GET /admin/promocodes/
Позволяет получить все существующие промокоды. 
Без параметров.
Ответ: 

```
[
	{
		"id":11,
		"promocode":"GET_FREE_DIAMONDS24",
		"reward_id":4,
		"expires":"2025-12-01T00:30:00Z",
		"max_uses":10,
		"remain_uses":10
	}
]
```
  
#### POST /admin/promocodes/promocode/
Создает новый промокод. 

Параметры  - **Body**: 

```
{
	"promocode":"GET_FREE_DIAMONDS24",
	"reward_id":4,
	"expires":"2025-12-01T00:30:00Z", // опционально
	"max_uses":10
}
```

Ответ: 

```
{
	"id":  11
}
```
  
#### GET /admin/promocodes/promocode/
Получение промокода по **id**. 

Параметры  - **Body**: 

```
{
	"id":  11
}
```

Ответ: 

```
{
	"id":11,
	"promocode":"GET_FREE_DIAMONDS24",
	"reward_id":4,
	"expires":"2025-12-01T00:30:00Z", 
	"max_uses":10,
	"remain_uses":  10
}
```

#### DELETE /admin/promocodes/promocode/
Удаление промокода. 

Параметры  - **Body**: 

```
{
	"promocode":  "GET_FREE_DIAMONDS24"
}
```

Ответ: 

```
{}
```

#### UPDATE /admin/promocodes/promocode/
Обновление промокода. 

Параметры  - **Body**: 

```
{
	"id":11, //обязательно, остальное - опционально
	"promocode":"GET_FREE_COINS",
	"reward_id":4,
	"expires":"2025-12-01T00:30:00Z", 
	"max_uses":10
}
```

Ответ: 

```
{
	"id":  11
}
```
  

***Запросы для работы с наградами:***


#### GET /admin/rewards/
Позволяет получить все существующие награды. 
Без параметров.
Ответ: 

```
[
	{
		"id":11,
		"title":  "DIAMOND5",
		"description":  "Добавляет 5 алмазов игроку"
	}
]
```
  
#### POST /admin/rewards/reward/
Создает новую награду. 

Параметры  - **Body**: 

```
{
	"title":  "DIAMOND5",
	"description":  "Добавляет 5 алмазов игроку"
}
```

Ответ: 

```
{
	"id":  11
}
```
  
  #### GET /admin/rewards/reward/
Получение награды по **id**. 

Параметры  - **Body**: 

```
{
	"id":  11
}
```

Ответ: 

```
{
	"id":11,
	"title":  "DIAMOND5",
	"description":  "Добавляет 5 алмазов игроку"
}
```

  #### DELETE /admin/rewards/reward/
Удаление награды. 

Параметры  - **Body**: 

```
{
	"title":  "DIAMOND5"
}
```

Ответ: 

```
{}
```

  #### UPDATE /admin/promocodes/promocode/
Обновление промокода. 

Параметры  - **Body**: 

```
{
	"id":11, //обязательно, остальное - опционально
	"promocode":"TEST123",
	"reward_id":4,
	"expires":"2025-12-01T00:30:00Z", 
	"max_uses":10
}
```

Ответ: 

```
{
	"id":  11
}
```

### UI
Это небольшой прототип, который запускается локально и позволяет обращаться к сервису `promocodes` для применения промокода. 


<img width="905" alt="image" src="https://github.com/user-attachments/assets/d9e52607-2748-41b8-a031-cae7cf4773ef">



# Инструкция для запуска

Необходимо установить переменные среды для сервисов admin и promocodes. Шаблон .env файла находится в директории `/configs` каждого проекта.

Для запуска сервисов admin, promocodes и базы данных достаточно выполнить команду в корне проекта:

    make run

> [!TIP]
> Для удобства управления сервисом admin в проекте находится JSON файл:  `Promocodes.postman_collection.json`, который можно импортировать в Postman. 


Для запуска UI необходимо перейти в директорию /ui 
```
cd ui
```

и выполнить команды:

```
npm install
```

```
npm start
```
