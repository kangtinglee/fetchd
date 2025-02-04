module github.com/CosmWasm/wasmd

go 1.13

require (
	// Note: update ENV GO_COSMWASM in Dockerfile when updating this
	github.com/CosmWasm/go-cosmwasm v0.9.1
	github.com/btcsuite/btcd v0.0.0-20190807005414-4063feeff79a // indirect
	github.com/cosmos/cosmos-sdk v0.38.3
	github.com/golang/mock v1.4.3 // indirect
	github.com/google/gofuzz v1.0.0
	github.com/gorilla/mux v1.7.4
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/otiai10/copy v1.0.2
	github.com/otiai10/curr v0.0.0-20190513014714-f5a3d24e5776 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rcrowley/go-metrics v0.0.0-20190706150252-9beb055b7962 // indirect
	github.com/snikch/goodman v0.0.0-20171125024755-10e37e294daa
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v0.0.6
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.5.1
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.33.3
	github.com/tendermint/tm-db v0.5.1
	go.etcd.io/bbolt v1.3.4 // indirect
	golang.org/x/net v0.0.0-20190827160401-ba9fcec4b297 // indirect
	golang.org/x/sys v0.0.0-20200602225109-6fdc65e7d980 // indirect
)

replace github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4

// this include a few extra debug helpers on top of cosmos v0.38.3 but original also works fine
replace github.com/cosmos/cosmos-sdk => github.com/fetchai/cosmos-sdk v0.16.4

replace github.com/tendermint/tendermint => github.com/fetchai/cosmos-consensus v0.16.1
