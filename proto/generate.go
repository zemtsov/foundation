package proto

//go:generate protoc -I=. --go_out=paths=source_relative:. task.proto
//go:generate protoc -I=. --go_out=paths=source_relative:. batch.proto
//go:generate protoc -I=. --go_out=paths=source_relative:. chtransfer.proto
//go:generate protoc -I=. --go_out=paths=source_relative:. report.proto
//go:generate protoc -I=. --go_out=paths=source_relative:. locks.proto
//go:generate protoc -I=. -I=./validate --go_out=paths=source_relative:. --validate_out=lang=go,paths=source_relative:. transfer_request.proto

// Chaincode configuration
//go:generate protoc -I=. -I=./validate --go_out=paths=source_relative:. --validate_out=lang=go,paths=source_relative:. foundation_config.proto
