watch:
	find . -name '*.go' -o -path './sample/*' | entr -c go run *.go ./sample/
