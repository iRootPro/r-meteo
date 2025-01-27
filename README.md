# MeteoWeb

MeteoWeb — это дашборд-сайт для метеостанции, который собирает и отображает данные о погоде с различных сенсоров. Проект включает в себя бэкенд на Go и фронтенд на JavaScript, обеспечивая возможность обработки и визуализации данных о погоде.

## Архитектура

Система состоит из следующих компонентов:
- **Сенсоры**: Измеряют параметры погоды (температура, влажность и т.д.) и отправляют данные в Home Assistant.
- **Home Assistant**: Управляет сенсорами и хранит данные в MariaDB.
- **Бэкенд**: Написан на Go, предоставляет REST API для взаимодействия с фронтендом и базой данных. Поддерживает различные команды:
  - **web**: Веб-приложение для отображения данных.

- **Фронтенд**: Написан на JavaScript, отображает данные и графики о погоде, полученные через REST API.

## Структура проекта


