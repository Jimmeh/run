all: build

build: .
	go build -o built/run .