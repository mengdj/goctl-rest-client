#!/bin/sh
goctl api go -api test.api -dir . --style go_zero
goctl api plugin -p goctl-rest-client="rest-client --package=client --destination=exa_api" -api test.api -dir .