all: hello

hello: hello.go
	go build hello.go

run: hello
	./hello man

clean:
	rm -rf hello
