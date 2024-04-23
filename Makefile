default: out/app

clean:
	rm -rf out

test:
	go test ./...

build: cmd/painter/main.go painter/loop.go painter/op.go painter/lang/http.go ui/window.go
	mkdir -p out
	go build -o out/app ./cmd/painter
