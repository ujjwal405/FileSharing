.PHONY: build-UploadSignedURLFunction 

build-UploadSignedURLFunction:
	go mod vendor
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap cmd/main.go
	chmod +x bootstrap
	cp bootstrap $(ARTIFACTS_DIR)/bootstrap

clean:
	rm -f bootstrap
