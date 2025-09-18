DEFAULT := ./cmd
OUTPUT := ./bin/mall
default: build
build: 
	go build -gcflags "-N -l" -o $(OUTPUT) $(DEFAULT)