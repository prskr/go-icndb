#!/usr/bin/env bash

swagger_path=$(which swagger)

if [[ $? -ne 0 ]]; then
    echo "Installing swagger because it's not in PATH"
    go get -u github.com/go-swagger/go-swagger/cmd/swagger
fi

swagger generate server --target ./ --name ICNDB --spec ./assets/api/swagger.yml