# Volume coding challenge

#### API format:
Defined using OpenAPI format. Can be found in api/FlightTrackerApi.yml
Generated code using oapi-codegen tool (configuration is in api/oapi-codegen-config.yml)

#### API exposed: http://localhost:8080/calculate
#### HTTP method: POST

#### Tools/Frameworks used:
    API spec defined using: OpenAPI
    Boilerplate code generation using: oapi-codegen
    Server framerwork: Echo

#### Makefile targets:
```
gen, build, test, run
```

#### Test with:
```
% make run
```

#### And then use following commands to send requests:

```
curl -X 'POST' \
  'http://localhost:8080/calculate' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '[  [ "SFO", "EWR" ] ]'
```

```
curl -X 'POST' \
  'http://localhost:8080/calculate' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '[["ATL", "EWR"], ["SFO", "ATL"]]'
```
```
curl -X 'POST' \
  'http://localhost:8080/calculate' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]'
```
