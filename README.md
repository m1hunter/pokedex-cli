# Pokedex-CLI
Простая консольная REPL мини-игра по поимке покемонов, написанная на языке GO.
## Команды
- help - информация о всех командах;
- сatch <pokemon_name> - попробовать поймать покемона;
- inspect <pokemon_name> - узнать информацию о пойманном покемоне;
- explore <location_name> - исследовать локацию на наличие покемонов;
- map - листать список с локациями вперед;
- mapb - листать список с локациями назад;
- exit - выход из игры.

## Дальнейшая разработка

- Будет добавлена команда pokedex, которая выведет всех пойманных покемонов.

- Сохранение прогресса в базу данных.

## Геймплей
Вы - охотник на покемонов. Вам нужно поймать 
понравившегося покемона. Выберите подходящую локацию и киньте покебол,
и если удача на Вашей стороне, желанный покемон станет Вашим! Но учтите,
они тоже не пальцем деланы, и могут убежать.

## Структура проекта

`cmd/main.go` - точка запуска игры;

`internal/commands/command_catch.go` - модуль поимки покемона;

`internal/commands/command_exit.go` - модуль завершения игры;

`internal/commands/command_explore.go` - модуль исследования локации и находящихся на ней покемонов;

`internal/commands/command_help.go` - модуль отображения информации;

`internal/commands/command_inspect.go` - модуль отображения информации о покемоне;

`internal/commands/command_map.go` - модуль пролистывания страниц локаций;

`internal/commands/repl.go` - модуль, который отвечает за выполение команд пользователя в консоли;

`internal/pokeapi/...` - модули, связанные с API PokeAPI;

`internal/pokecache/pokecache.go` - модуль, который отвечает за создание кэша и его очистку.

## Используемые технологии
 - REST API. Мы получаем данные в JSON-формате по endpoint'ам в PokeAPI по
адресу https://pokeapi.co/api/v2/.
 - Используется кэширование. Данные хранятся в хэш-таблицах, и каждые 5 минут из них удаляются старые записи с помощью
функции `reap()` в `internal/pokecache.go`.

## Запуск

Загрузите содержимое репозитория:

`git clone https://github.com/m1hunter/pokedex-cli.git`

Перейдите в папку `cmd` и выполните команду:

`go run main.go`

