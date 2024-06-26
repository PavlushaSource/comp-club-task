### Тестовое задание на должность инженера по разработке ПО для базовых станций Golang.

---

С техническим заданием можно ознакомиться
по [ссылке](https://docs.google.com/document/d/1jAU5xgP2K9iGsxDtyCkhKECdjs7-oNC_/edit?usp=sharing&ouid=105079011505262905999&rtpof=true&sd=true)

## Installation

```bash
git clone https://github.com/PavlushaSource/comp-club-task.git
```

## Usage

### Makefile

```bash
make
./task -input <input_log_path>
```

### Build and run with Go

```bash
go build ./cmd/app -o task
./task -input <input_log_path>
```

### Docker

**Сборка**

```bash
docker build --tag compclub .
```

**Запуск образа**

```bash
docker run compclub -input <input_log_path>
```

---

## Тестирование и краевые случаи

К сожалению, из-за отсутствия времени на зачётной неделе в университете, успел написать лишь интеграционные тесты.

**Кроме предложенного сценария в ТЗ, были рассмотрены следующие:**

- Время окончания работы клуба происходит в начале следующего дня, то есть работа всю ночь (closeLessOpenHours.txt).
- Ни одного события за рабочий день не произошло (emptyEvents.txt).
- Прибыль со стола округляется в большую сторону, даже если клиент просидел лишних 3 минуты (roundProfit.txt).
- Cобытия, произошедшие после завершения рабочего дня, не рассматриваются (eventsAfterClose.txt).

**Написаны сценарии для проверки корректной валидации входного файла, а именно:**

- Корректное имя клиента (uncorrectedName.txt).
- События сортированы не по порядку (uncorrectedSortEvents.txt).
- Ввод удовлетворяет формату времени из ТЗ (uncorrectedTimeFormat.txt).

**Для некоторых совсем простых валидаций интеграционные тесты не были написаны:**

- Положительное целое число для количества столов и стоимости часа в компьютерном клубе
- Каждый стол пронумерован от 1 до N.
- Сесть за несуществующий стол.