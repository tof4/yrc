build:
	go build -o bin/serv cmd/serv/main.go
	go build -o bin/channeladd cmd/channeladd/main.go
	go build -o bin/channeldel cmd/channeldel/main.go
	go build -o bin/useradd cmd/useradd/main.go
	go build -o bin/userdel cmd/userdel/main.go
	go build -o bin/usermod cmd/usermod/main.go

clean:
	mkdir -p bin
	rm -r bin/*