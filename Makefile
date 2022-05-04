watch:
	find . -name '*.go' -o -path './example/*' | entr -c go run parsing.go mkblog.go ./example/

test:
	go test

clean:
	rm -f mkblog
