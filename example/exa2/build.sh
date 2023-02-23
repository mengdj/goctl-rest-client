#!/bin/sh
goctl api go -api exa.api -dir . --style go_zero
goctl api plugin -p goctl-rest-client="rest-client" -api exa.api -dir .