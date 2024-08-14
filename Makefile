# Указываем переменные
GOPATH := $(shell go env GOPATH)
GOARCH := arm64
GOOS := linux
TARGET := hw
MAIN := cmd/main.go

# Цель по умолчанию (build и run)
all: pull build run

# Выполняем git pull для обновления репозитория
pull:
	git pull

# Сборка Go программы
build:
	GOARCH=$(GOARCH) GOOS=$(GOOS) go build -o $(TARGET) $(MAIN)

# Запуск программы
run:
	./$(TARGET)

# Чистка скомпилированных файлов
clean:
	rm -f $(TARGET)
