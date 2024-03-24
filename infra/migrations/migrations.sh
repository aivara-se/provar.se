#!/bin/bash

migrate -database 'postgresql://postgres:postgres@localhost:5432/provar?sslmode=disable' -path db/migrations up
migrate -database 'postgresql://postgres:postgres@localhost:5432/provar-test?sslmode=disable' -path db/migrations up
