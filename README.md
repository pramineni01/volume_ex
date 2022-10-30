# volume_ex
Volume exercise

Make with targets: gen, build, test, run


Test with:
% make run

And then use following commands to send requests:
curl -X 'POST' \
  'http://localhost:8080/calculate' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '[  [ "SFO", "EWR" ] ]'

curl -X 'POST' \
  'http://localhost:8080/calculate' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '[["ATL", "EWR"], ["SFO", "ATL"]]'

curl -X 'POST' \
  'http://localhost:8080/calculate' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]'