module github.com/maticnetwork/heimdall

go 1.12

require (
	github.com/allegro/bigcache v1.2.1 // indirect
	github.com/apilayer/freegeoip v3.5.0+incompatible // indirect
	github.com/aristanetworks/goarista v0.0.0-20191206003309-5d8d36c240c9 // indirect
	github.com/btcsuite/btcd v0.0.0-20190629003639-c26ffa870fd8 // indirect
	github.com/cbergoon/merkletree v0.2.0
	github.com/cespare/cp v1.1.1 // indirect
	github.com/cosmos/cosmos-sdk v0.38.1
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/edsrzf/mmap-go v1.0.0 // indirect
	github.com/elastic/gosigar v0.10.5 // indirect
	github.com/fatih/color v1.9.0 // indirect
	github.com/fjl/memsize v0.0.0-20190710130421-bcb5799ab5e5 // indirect
	github.com/gballet/go-libpcsclite v0.0.0-20191108122812-4678299bea08 // indirect
	github.com/go-kit/kit v0.9.0
	github.com/gogo/protobuf v1.3.1
	github.com/gorilla/mux v1.7.3
	github.com/howeyc/fsnotify v0.9.0 // indirect
	github.com/huin/goupnp v1.0.0 // indirect
	github.com/jackpal/go-nat-pmp v1.0.2 // indirect
	github.com/karalabe/usb v0.0.0-20191104083709-911d15fe12a9 // indirect
	github.com/maticnetwork/bor v0.0.0-20191204165821-bd9cd503a1b3
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/olekukonko/tablewriter v0.0.4 // indirect
	github.com/oschwald/maxminddb-golang v1.6.0 // indirect
	github.com/pborman/uuid v1.2.0
	github.com/peterh/liner v1.2.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/tsdb v0.10.0 // indirect
	github.com/prysmaticlabs/prysm v0.0.0-20190507024903-1be950f90cad
	github.com/rakyll/statik v0.1.6
	github.com/rcrowley/go-metrics v0.0.0-20190706150252-9beb055b7962 // indirect
	github.com/rjeczalik/notify v0.9.2 // indirect
	github.com/robertkrimen/otto v0.0.0-20191219234010-c382bd3c16ff // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.6.2
	github.com/status-im/keycard-go v0.0.0-20191119114148-6dd40a46baa0 // indirect
	github.com/steakknife/bloomfilter v0.0.0-20180922174646-6819c0d2a570 // indirect
	github.com/steakknife/hamming v0.0.0-20180906055917-c99c65617cd3 // indirect
	github.com/streadway/amqp v0.0.0-20190404075320-75d898a42a94
	github.com/stretchr/testify v1.4.0
	github.com/syndtr/goleveldb v1.0.1-0.20190923125748-758128399b1d
	github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.33.0
	github.com/tendermint/tm-db v0.4.0
	github.com/tyler-smith/go-bip39 v1.0.2 // indirect
	github.com/wsddn/go-ecdh v0.0.0-20161211032359-48726bab9208 // indirect
	github.com/xsleonard/go-merkle v0.0.0-20190513010722-059a904a0479
	golang.org/x/crypto v0.0.0-20191206172530-e9b2fee46413
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	gopkg.in/olebedev/go-duktape.v3 v3.0.0-20190709231704-1e4459ed25ff // indirect
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
	gopkg.in/urfave/cli.v1 v1.20.0 // indirect
	gopkg.in/yaml.v2 v2.2.8
)

// replace github.com/tendermint/tendermint => github.com/maticnetwork/tendermint v0.26.0-dev0.0.20200105181728-9b2d312d5d14
replace github.com/tendermint/tendermint => github.com/maticnetwork/tendermint v0.26.0-dev0.0.20200215173139-de3db64e4b11

// replace golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15
