keeb: **/*.go
	go build cmd/keeb.go

install: **/*.go
	go install cmd/keeb.go

run: keeb
	./$<

clean:
	rm -f keeb
