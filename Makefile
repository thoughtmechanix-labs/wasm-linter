.PHONY: all test clean

all: buildwasm buildserver 

clean:
	rm -f assets/main.wasm && rm -f cmd/server/server

buildwasm: 
	cd cmd/wasm && GOOS=js GOARCH=wasm go build  -o ../../assets/main.wasm && cd ../..

buildserver:
	cd cmd/server && go build  -o ../../bin/server && cd ../.. 