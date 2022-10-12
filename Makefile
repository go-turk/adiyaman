test:
	go test -v ./...
fuzz:
	 go test ./calctree -fuzz=FuzzNode_Repair -test.fuzztime 6s
run :
	go run main.go