# Makefile для контрольной работы вариант 1
# Цели:
# all - выполняет build и test
# build - компилирует проект
# test - запускает тесты
# run - запускает программу
# clean - очищает артефакты сборки
# deploy - создает исполняемый файл (опционально)

.PHONY: all build test run clean deploy

all: build test

build:
	@echo "Building..."
	go build -o bin/main main.go

test1:
	@echo "Running tests..."
	go test -v ./algo

run:
	@echo "Running program..."
	go run main.go

clean:
	@echo "Cleaning..."
	rm -rf bin/

deploy: build
	@echo "Deploying..."
	cp bin/main /usr/local/bin/fibonacci-secondmax