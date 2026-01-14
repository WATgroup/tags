## tags

## How to test with general coverage:

go test -cover   



## How to inspect cover of functions:

1. Generate binary for metrics:
go test ./... -coverprofile=cover

2. Function Report: 
go tool cover -func=cover

3. Visual Report:
go tool cover -html cover