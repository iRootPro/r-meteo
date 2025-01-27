Архитектурная документация - Проект MeteoWeb

```plantuml
@startuml
!define RECTANGLE
!includeurl https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml

title MeteoWeb - C4 Container Diagram

LAYOUT_LEFT_RIGHT()

Person(user, "User", "Использует веб-интерфейс для просмотра данных метеостанции")

System_Boundary(meteo_system, "MeteoWeb System") {
    Container(frontend, "Frontend", "React/Vue", "Отображает данные метеостанции")
    Container(backend, "Backend", "Golang/Python", "Обрабатывает запросы от фронтенда, взаимодействует с базой данных и внешним API")
    ContainerDb(database, "MariaDB", "SQL Database", "Хранит данные сенсоров и расчетные значения")
    Container(home_assistant, "Home Assistant", "Automation Hub", "Собирает данные с метеостанции и сохраняет их в базу данных")
    Container(redis, "Redis", "In-Memory Cache", "Кэширует частые запросы для повышения производительности")
    Container(sensors, "Weather Sensors", "Hardware", "Собирают данные о погоде (температура, давление, влажность)")
}

System_Ext(api_service, "Weather API", "External API", "Предоставляет текущие данные и прогноз погоды")

Rel(user, frontend, "Запрашивает данные")
Rel(frontend, backend, "Запросы через REST API")
Rel(backend, database, "Запрашивает и сохраняет данные")
Rel(backend, redis, "Часто запрашиваемые данные")
Rel(backend, api_service, "Запрашивает текущие данные и прогноз погоды")
Rel(home_assistant, database, "Сохраняет данные сенсоров")
Rel(sensors, home_assistant, "Отправляют данные о погоде")

@enduml
```

Описание проекта:

MeteoWeb — это система мониторинга метеорологических данных, которая отображает информацию с датчиков метеостанции на веб-дашборде. Все данные собираются через Home Assistant и хранятся в базе данных MariaDB. Дополнительно система может взаимодействовать с внешним API для получения актуальных данных о погоде и прогнозов.
Текущая схема работы:

   - Датчики метеостанции передают данные в Home Assistant.
   - Home Assistant сохраняет эти данные в базу данных MariaDB.
   - Пользователь взаимодействует с веб-интерфейсом (Frontend), который запрашивает данные через Backend.
   - Backend обрабатывает запросы, взаимодействуя с Redis для кэширования часто запрашиваемых данных и с MariaDB для получения/сохранения информации.
   - Внешний API-сервис позволяет Backend получать данные о текущей погоде и прогнозах.

Добавление внешнего API:

В системе предусмотрена возможность обращения к сторонним сервисам через API для получения:

   - Актуальных данных о текущей погоде.
   - Прогнозов погоды на заданный период.

Это взаимодействие реализуется на уровне Backend. Сервер отправляет запросы к внешнему API и обрабатывает полученные данные для отображения на фронтенде или для дальнейшего использования и хранения в базе данных.
Потенциальные задачи, решаемые через внешние API:

   - Получение оперативных данных для улучшения точности мониторинга.
   - Использование прогностических данных для расширения возможностей аналитики системы.

