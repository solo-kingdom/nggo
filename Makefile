.PHONY: all

resolve: scripts/deps.sh
	@sh $< $(args)

nggo: cmd/nggo/main/main.go
	@go run $< $(args)

head: cmd/nggo/head/head.go
	@go run $< $(args)

tst: cmd/nggo/tst/tst.go
	@go run $< $(args)

format: cmd/nggo/format/format.go
	@go run $< $(args)

decode: cmd/nggo/decode/decode.go
	@go run $< $(args)

sample: cmd/nggo/sample/sample.go
	@go run $< $(args)