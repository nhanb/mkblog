watch:
	find . -name '*.go' -o -path './sample/*' | entr -c go run parsing.go mkblog.go ./sample/

test:
	go test
