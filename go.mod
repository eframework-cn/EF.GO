module eframe

go 1.18

require (
	github.com/eframework-cn/EP.GO.CORE v0.0.0
	github.com/eframework-cn/EP.GO.UTIL v1.0.1
)

replace (
	github.com/eframework-cn/EP.GO.CORE => ./lib/EP.GO.CORE
	github.com/eframework-cn/EP.GO.UTIL => ./lib/EP.GO.UTIL
)

require github.com/golang/protobuf v1.5.3

require (
	github.com/Microsoft/go-winio v0.5.2 // indirect
	github.com/armon/go-metrics v0.0.0-20180917152333-f0300d1749da // indirect
	github.com/fatih/color v1.9.0 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/hashicorp/consul/api v1.21.0 // indirect
	github.com/hashicorp/consul/sdk v0.13.1 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.1 // indirect
	github.com/hashicorp/go-hclog v0.12.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.0.0 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/lxn/win v0.0.0-20210218163916-a377121e959e // indirect
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/mediocregopher/radix.v2 v0.0.0-20181115013041-b67df6e626f9 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	nanomsg.org/go-mangos v1.4.0 // indirect
)
