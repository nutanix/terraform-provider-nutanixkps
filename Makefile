TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=hashicorp.com
NAMESPACE=nutanix
NAME=nutanixkps
BINARY=terraform-provider-${NAME}
VERSION=0.1
GOOS=linux
GOARCH=amd64
default: install

build:
	go build -o ${BINARY}_${VERSION}

installgolang:
	mkdir -p ./golang/installer
	curl -o ./golang/installer/go1.15.7.linux-amd64.tar.gz https://storage.googleapis.com/golang/go1.15.7.linux-amd64.tar.gz
	tar -C ${HOME} -xzf ./golang/installer/go1.15.7.linux-amd64.tar.gz
	ls ${HOME}/go
	echo "PATH=${PATH}"
	echo "GOROOT=${GOROOT}"
	echo "GOBIN=${GOBIN}"
	echo "PATH=${PATH}"
	go version

installkpsclient:
	mkdir -p ./generated/kps_cloud_api_swagger
	mkdir -p ./bin
	curl -o ./generated/kps_api_full.json https://karbon.nutanix.com/kps_api_full.json
	
	# GOOS=darwin GOARCH=amd64 curl -o ./bin/swagger/${GOOS}_${GOARCH} -L'#' https://github.com/go-swagger/go-swagger/releases/download/v0.21.0/swagger_${GOOS}_${GOARCH}
	# chmod +x ./bin/swagger/${GOOS}_${GOARCH}
	# ./bin/swagger/${GOOS}_${GOARCH} generate client -t ./generated/kps_cloud_api_swagger -f ./generated/kps_api_full.json -A sherlock
	
	# GOOS=freebsd GOARCH=386 curl -o ./bin/swagger/${GOOS}_${GOARCH} -L'#' https://github.com/go-swagger/go-swagger/releases/download/v0.21.0/swagger_${GOOS}_${GOARCH}
	# chmod +x ./bin/swagger/${GOOS}_${GOARCH}
	# ./bin/swagger/${GOOS}_${GOARCH} generate client -t ./generated/kps_cloud_api_swagger -f ./generated/kps_api_full.json -A sherlock
	
	# GOOS=freebsd GOARCH=amd64 curl -o ./bin/swagger/${GOOS}_${GOARCH} -L'#' https://github.com/go-swagger/go-swagger/releases/download/v0.21.0/swagger_${GOOS}_${GOARCH}
	# chmod +x ./bin/swagger/${GOOS}_${GOARCH}
	# ./bin/swagger/${GOOS}_${GOARCH} generate client -t ./generated/kps_cloud_api_swagger -f ./generated/kps_api_full.json -A sherlock

	# GOOS=freebsd GOARCH=arm curl -o ./bin/swagger/${GOOS}_${GOARCH} -L'#' https://github.com/go-swagger/go-swagger/releases/download/v0.21.0/swagger_${GOOS}_${GOARCH}
	# chmod +x ./bin/swagger/${GOOS}_${GOARCH}
	# ./bin/swagger/${GOOS}_${GOARCH} generate client -t ./generated/kps_cloud_api_swagger -f ./generated/kps_api_full.json -A sherlock

	# GOOS=linux GOARCH=386 curl -o ./bin/swagger/${GOOS}_${GOARCH} -L'#' https://github.com/go-swagger/go-swagger/releases/download/v0.21.0/swagger_${GOOS}_${GOARCH}
	# chmod +x ./bin/swagger/${GOOS}_${GOARCH}
	# ./bin/swagger/${GOOS}_${GOARCH} generate client -t ./generated/kps_cloud_api_swagger -f ./generated/kps_api_full.json -A sherlock

	curl -o ./bin/swagger_${GOOS}_${GOARCH} -L'#' https://github.com/go-swagger/go-swagger/releases/download/v0.21.0/swagger_${GOOS}_${GOARCH}
	chmod +x ./bin/swagger_${GOOS}_${GOARCH}
	./bin/swagger_${GOOS}_${GOARCH} generate client -t ./generated/kps_cloud_api_swagger -f ./generated/kps_api_full.json -A sherlock

	# GOOS=linux GOARCH=arm curl -o ./bin/swagger/${GOOS}_${GOARCH} -L'#' https://github.com/go-swagger/go-swagger/releases/download/v0.21.0/swagger_${GOOS}_${GOARCH}
	# chmod +x ./bin/swagger/${GOOS}_${GOARCH}
	# ./bin/swagger/${GOOS}_${GOARCH} generate client -t ./generated/kps_cloud_api_swagger -f ./generated/kps_api_full.json -A sherlock

	# GOOS=openbsd GOARCH=386 curl -o ./bin/swagger/${GOOS}_${GOARCH} -L'#' https://github.com/go-swagger/go-swagger/releases/download/v0.21.0/swagger_${GOOS}_${GOARCH}
	# chmod +x ./bin/swagger/${GOOS}_${GOARCH}
	# ./bin/swagger/${GOOS}_${GOARCH} generate client -t ./generated/kps_cloud_api_swagger -f ./generated/kps_api_full.json -A sherlock

	# GOOS=openbsd GOARCH=amd64 curl -o ./bin/swagger/${GOOS}_${GOARCH} -L'#' https://github.com/go-swagger/go-swagger/releases/download/v0.21.0/swagger_${GOOS}_${GOARCH}
	# chmod +x ./bin/swagger/${GOOS}_${GOARCH}
	# ./bin/swagger/${GOOS}_${GOARCH} generate client -t ./generated/kps_cloud_api_swagger -f ./generated/kps_api_full.json -A sherlock

	# GOOS=solaris GOARCH=amd64 curl -o ./bin/swagger/${GOOS}_${GOARCH} -L'#' https://github.com/go-swagger/go-swagger/releases/download/v0.21.0/swagger_${GOOS}_${GOARCH}
	# chmod +x ./bin/swagger/${GOOS}_${GOARCH}
	# ./bin/swagger/${GOOS}_${GOARCH} generate client -t ./generated/kps_cloud_api_swagger -f ./generated/kps_api_full.json -A sherlock

	# GOOS=windows GOARCH=386 curl -o ./bin/swagger/${GOOS}_${GOARCH} -L'#' https://github.com/go-swagger/go-swagger/releases/download/v0.21.0/swagger_${GOOS}_${GOARCH}
	# chmod +x ./bin/swagger/${GOOS}_${GOARCH}
	# ./bin/swagger/${GOOS}_${GOARCH} generate client -t ./generated/kps_cloud_api_swagger -f ./generated/kps_api_full.json -A sherlock

	# GOOS=windows GOARCH=amd64 curl -o ./bin/swagger/${GOOS}_${GOARCH} -L'#' https://github.com/go-swagger/go-swagger/releases/download/v0.21.0/swagger_${GOOS}_${GOARCH}
	# chmod +x ./bin/swagger/${GOOS}_${GOARCH}
	# ./bin/swagger/${GOOS}_${GOARCH} generate client -t ./generated/kps_cloud_api_swagger -f ./generated/kps_api_full.json -A sherlock

release:
	# GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_${GOOS}_${GOARCH}
	# GOOS=freebsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_${GOOS}_${GOARCH}
	# GOOS=freebsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_${GOOS}_${GOARCH}
	# GOOS=freebsd GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_${GOOS}_${GOARCH}
	# GOOS=linux GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_${GOOS}_${GOARCH}
	GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_${GOOS}_${GOARCH}
	# GOOS=linux GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_${GOOS}_${GOARCH}
	# GOOS=openbsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_${GOOS}_${GOARCH}
	# GOOS=openbsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_${GOOS}_${GOARCH}
	# GOOS=solaris GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_${GOOS}_${GOARCH}
	# GOOS=windows GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_${GOOS}_${GOARCH}
	# GOOS=windows GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_${GOOS}_${GOARCH}

install: 
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${GOOS}_${GOARCH}
	mv ./bin/${BINARY}_${VERSION}_${GOOS}_${GOARCH} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${GOOS}_${GOARCH}
	mkdir -p ${HOME}/artifacts
	cp -r ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/* ${HOME}/artifacts

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

docs:
	tfplugindocs generate
