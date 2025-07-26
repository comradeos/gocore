BINARY=ping_logger
SRC=main.go

all: build run

build:
	go build -o $(BINARY) $(SRC)

run:
	./$(BINARY)

clean:
	rm -f $(BINARY) ping.log