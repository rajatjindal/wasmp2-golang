using tinygo branch from Damian

export PATH=/Users/rajatjindal/go/src/github.com/tinygo-org/tinygo/build/release/tinygo/bin:$PATH
export TINYGOROOT=~/go/src/github.com/tinygo-org/tinygo/build/release/tinygo/



tinygo build -target=wasip2 -wit-package ./wit -wit-world foo:foo/foo-world -o main.wasm -x -work main.go

wasm-tools dump ./target/wasm32-wasi/release/http_trigger.wasm | grep 'Import.*key-value.*open'                 