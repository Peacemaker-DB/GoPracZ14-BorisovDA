# Практическое задание № 14 Борисов Денис Александрович ЭФМО-01-25

Цели:

1.	Научиться находить «узкие места» в SQL-запросах и устранять их (индексы, переписывание запросов, пагинация, батчинг).
2.	Освоить настройку пула подключений (connection pool) в Go и параметры его тюнинга.
3.	Научиться использовать EXPLAIN/ANALYZE, базовые метрики (pg_stat_statements), подготовленные запросы и транзакции.
4.	Применить техники уменьшения N+1 запросов и сокращения аллокаций на горячем пути.

# Выполнение практического задания.

  Для выполнение практики заа основу был взят из практики 11
  
1. Структура проекта

  Для выполнения практической работы была сделана следующая структура проекта

<img width="248" height="446" alt="image" src="https://github.com/user-attachments/assets/f39c9528-3a8f-4fe9-98f8-fca355aac546" />

   Так же были установлены все необходимые расширения для выполнения практики

<img width="630" height="177" alt="Снимок экрана 2025-12-13 200055" src="https://github.com/user-attachments/assets/a7280373-7ef5-4b01-ae19-a8eb4b423b7a" />

2.	Docker-Compose.
  Для реализации контейнеризации практики был написан файл docker-compose.yml

<img width="470" height="411" alt="image" src="https://github.com/user-attachments/assets/c94acb7d-c301-4041-81fb-d2d3babd90c5" />

3. Миграция.
 Для реализациии миграции в Postgres-контейнера, был написан файл init.sql

<img width="406" height="470" alt="image" src="https://github.com/user-attachments/assets/27df865b-1d36-4a37-8b70-6759208ae0c8" />

4. Конфигурация.

  Был написан файл конфигурации config.go для подлкючения к Postgres-контейнера

<img width="540" height="302" alt="image" src="https://github.com/user-attachments/assets/2fd44002-3787-4ac3-9947-15b4a0749856" />
 
5. Postgres и интерфейс репозиторий.

  Для выполнения практики была написана файл note_pg.go. В котором реализован CRUD для обработки заметок при помощи Postgres-контейнера. Фрагмент фаайла

<img width="862" height="573" alt="image" src="https://github.com/user-attachments/assets/745a99b9-5b57-4c32-9f89-e39733cefb90" />


6. Поиск и пагинация: OFFSET vs keyset

   Для реализации Keyset-пагинация была напиисана функция в note_pg.go
   
<img width="850" height="512" alt="image" src="https://github.com/user-attachments/assets/3bd8292a-ca91-42c7-9cf4-12ff1dbfb7ce" />

7. Устранение N+1 (батчинг)

  Для реализации батчинга была напиисана функция в note_pg.go

  <img width="683" height="338" alt="image" src="https://github.com/user-attachments/assets/c08064af-2f43-4640-beaa-340ece9eaf1d" />

# Проверка работоспособности

  Для проверки работоспособности был запущен контейнер в Docker:

  <img width="815" height="91" alt="image" src="https://github.com/user-attachments/assets/35d44b77-a1ca-471a-a7a1-a7f83a98e7b7" />

  Нагрузочное тестирование

  Тестирование обычной пагинации

<img width="829" height="722" alt="image" src="https://github.com/user-attachments/assets/0b5cfe40-f2e0-44b5-805a-f302f5e8e5c2" />

  Тестирование keyset-пагинации

<img width="873" height="739" alt="image" src="https://github.com/user-attachments/assets/57dfa019-3ba4-4844-b4f2-e83f98497e05" />

  Тестирование батчинга

<img width="865" height="737" alt="image" src="https://github.com/user-attachments/assets/6e86adb8-28d1-4cf4-8658-07d509b938d2" />

  Тестирование получения одной заметки

<img width="741" height="723" alt="image" src="https://github.com/user-attachments/assets/94e4086c-eeea-48a0-8f6f-e707cd49585b" />

# Анализ запросов в PostgreSQL

  Анализ запроса с пагинацией

<img width="816" height="248" alt="image" src="https://github.com/user-attachments/assets/f85bcc27-ae32-481d-a430-48fa25c7ae69" />

  Анализ keyset-пагинациb
  
<img width="815" height="270" alt="image" src="https://github.com/user-attachments/assets/0f7cfbb2-f8bd-4b7f-bf9c-21282f293f60" />
