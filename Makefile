build: tracker.go controllers/TrackerController.go models/Event.go
	go build -o tracker . 

run: build
	./tracker


all: build
