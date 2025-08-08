#!/bin/bash
GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap cmd/aws_lambda/main.go
zip tv_frontal_gw.zip bootstrap