package main

//gp:generate tern migrate --migrations ./internal/pgstore/migrations/ --config ./internal/pgstore/migrations/tern.conf run go generate ./... | run go generate
//go:generate sqlc generate -f ./internal/pgstore/sqlc.yaml
