vendor-deps: 
	go mod vendor

fmt:
	go fmt flags/*.go

tools: 	
	# @GOPATH=$(GOPATH) go build -o bin/wof-staticd cmd/wof-staticd.go
