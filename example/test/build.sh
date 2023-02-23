#!/bin/sh
goctl api go -api test.api -dir . --style go_zero
goctl api plugin -p goctl-rest-client="rest-client" -api test.api -dir .