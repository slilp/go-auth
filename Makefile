run:
	go run cmd/main.go
generate: 
	go generate	./...
test:
	go	test	./...	-cover