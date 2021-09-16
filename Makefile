.PHONY: all

resolve: scripts/deps.sh
	@sh $<

nggo: cmd/nggo/main/main.go
	@go run $<

head: cmd/nggo/head/head.go
	@go run $<

tst: cmd/nggo/tst/tst.go
	@go run $<

format: cmd/nggo/format/format.go
	@go run $<

decode: cmd/nggo/decode/decode.go
	@go run $<