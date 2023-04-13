#!/bin/zsh

go build -o build/bookings cmd/web/*.go && ./build/bookings
