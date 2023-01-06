#!/bin/bash

docker pull quay.io/goswagger/swagger

go get github.com/go-openapi/errors
go get github.com/go-openapi/loads
go get github.com/go-openapi/runtime
go get github.com/go-openapi/runtime/flagext
go get github.com/go-openapi/runtime/middleware
go get github.com/go-openapi/runtime/security
go get github.com/go-openapi/spec
go get github.com/go-openapi/strfmt
go get github.com/go-openapi/swag
go get github.com/go-openapi/validate
go get github.com/jessevdk/go-flags
go get golang.org/x/net/netutil

PROJECT_PATH=/go/src/github.com/edyknopfler/apifirst

docker run --rm -it \
  -e GOPATH=/go \
  -v $GOPATH:/go \
  -w $PROJECT_PATH \
  quay.io/goswagger/swagger \
  generate server -A medidas -f $PROJECT_PATH/swagger.yml
