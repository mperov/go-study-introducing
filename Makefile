.SILENT: install

all: hello

hello: hello.go
	go build hello.go

run: hello
	./hello man

install: hello
	go install
	echo "go install"
	echo "This app is placed here - `go list -f '{{.Target}}'`"

clean:
	rm -rf hello
