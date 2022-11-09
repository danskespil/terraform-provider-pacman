HOSTNAME=danskespil.dk
NAMESPACE=prod
NAME=pacman
BINARY=terraform-provider-${NAME}
VERSION=1.0.0
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
OS_ARCH=${GOOS}_${GOARCH}

default: install

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

build: 
	go build
