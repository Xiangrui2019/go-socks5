default: deps
	cd ./app/gosocks5-local/;go install
	cd ../../
	cd ./app/gosocks5-server/;go install
	cd ../../

deps:
	go mod tidy
	go mod vendor

clean:
	rm -rf vendor/
	rm -rf go.sum

.PYHOY: go test