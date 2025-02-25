HOME = $(shell pwd)
DOCKER = docker run --net host -it --rm

build:
	GOOS=linux go build -o ./bin/feuer
	docker build -t lordrahl/feuer .
	rm -rf ./bin ./Library

run:
	docker-compose up

server:
	go run *.go

grafana:
	$(DOCKER) -p 3000:3000 -v $(HOME)/datasources.yml:/etc/grafana/provisioning/datasources/datasources.yml grafana/grafana:latest

prometheus:
	$(DOCKER) -p 9090:9090 -v $(HOME)/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus

influxdb:
	$(DOCKER) -p 8086:8086 quay.io/influxdb/influxdb:v2.0.3

telegraf:
	$(DOCKER) -v $(HOME)/telegraf.conf:/etc/telegraf/telegraf.conf telegraf