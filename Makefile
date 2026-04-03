include .env

PROJECT_ROOT := $(shell pwd)
export PROJECT_ROOT

up :
	docker compose up -d postgres

down :
	docker compose down postgres