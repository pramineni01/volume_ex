.PHONY: gen
# generate api
gen:
	oapi-codegen -config api/oapi-codegen-config.yml  api/FlightTrackerApi.yml > flightTracker/flightTracker.gen.go

.PHONY: build
# build
build:
	go build ./...


.PHONY: test
# test
test:
	go test ./...

.PHONY: run
# run
run:
	go run main.go