# Лабораторна робота №2 - Безперервна інтеграція та автоматизація тестування

### Варіант завдання: 5

- Функція: Перетворити постфіксний вираз у префіксний
- Бібліотека для перевірок: `testify`

## Інсталяція

Склонувати репозиторій:

```sh
git clone https://github.com/Bogdan-Zinovij/SA2.git
```

## Запуск програми:

```sh
go run cmd/example/main.go
```

Вказати опції:

- Для вводу виразу:
  - Ввести вираз у консолі:
  ```sh
    -e "expression"
  ```
  АБО
  - Вказати файл з виразом:
  ```sh
    -f "file.txt"
  ```
- Для виводу результату:
  - Вказати назву файлу, в який потрібно записати результат:
  ```sh
   -o "file.txt"
  ```
  - Якщо не вказати опцію `-o` то за замовчуванням результат виведеться в консоль

## Запуск тестів:

```sh
make test
```

### Приклади збірок:

- Неуспішні збірки:

  - Commit/push - https://github.com/Bogdan-Zinovij/SA2/actions/runs/2534476020
  - Pull request - https://github.com/Bogdan-Zinovij/SA2/actions/runs/2531166285

- Успішні збірки:
  - Commit/push - https://github.com/Bogdan-Zinovij/SA2/actions/runs/2534495150a
  - Pull request - https://github.com/Bogdan-Zinovij/SA2/actions/runs/2531264853

### Виконали:

- Василь Храпко - [@vazzz7zzzok](https://t.me/vazzz7zzzok) - khrapko2002@gmail.com
- Артем Матюшенко - [@artemko_m](https://t.me/artemko_m) - artom.matyushenko@gmail.com
- Богдан Зіновий - [@bzinovoy](https://t.me/bzinovoy) - bogdanolexandrov@gmail.com
