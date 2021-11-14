keeb: **/*.go
	go build cmd/keeb.go

.PHONY: firmware
firmware:
	(cd firmware && make)

install: **/*.go
	go install cmd/keeb.go

run: keeb
	./$<

.PHONY: clean
clean:
	rm -rf keeb firmware/.build
