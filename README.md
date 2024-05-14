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