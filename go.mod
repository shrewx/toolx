module github.com/shrewx/toolx

go 1.16

require (
	github.com/go-courier/packagesx v1.0.2
	github.com/spf13/cobra v1.5.0
)

require (
	github.com/Microsoft/go-winio v0.5.2 // indirect
	github.com/docker/distribution v2.8.1+incompatible // indirect
	github.com/docker/docker v20.10.17+incompatible
	github.com/docker/go-connections v0.4.0
	github.com/docker/go-units v0.4.0 // indirect
	github.com/gin-gonic/gin v1.8.1
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/moby/term v0.0.0-20220808134915-39b0c02b01ae // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/shrewx/enum v0.0.0-20220817101904-3ffdba5b63db
	github.com/shrewx/ginx v0.0.0-20220818022447-742dc9e502a7
	github.com/shrewx/statuserror v0.0.0-20220813053810-856d0205fd39
	golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab // indirect
	golang.org/x/time v0.0.0-20220722155302-e5dcc9cfc0b9 // indirect
	gotest.tools/v3 v3.3.0 // indirect
)

replace github.com/shrewx/ginx v0.0.0-20220818022447-742dc9e502a7 => ../ginx
