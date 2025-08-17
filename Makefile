.PHONY: all go rust run clean
all: go rust
go:  ; mkdir -p bin && go build -o bin/minivpn ./cmd/minivpn
rust: ; cd rust-crypto && cargo build
run: ; ./bin/minivpn --config ./configs/example.toml
clean: ; rm -rf bin && (cd rust-crypto && cargo clean)
