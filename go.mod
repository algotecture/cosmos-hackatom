module github.com/cosmos/sdk-application-tutorial

go 1.12

require (
	github.com/cosmos/cosmos-sdk v0.28.2-0.20190605232616-5f9c3fdf8895
	github.com/gorilla/mux v1.7.0
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/mattn/go-isatty v0.0.8 // indirect
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/prometheus/client_golang v0.9.4 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	github.com/syndtr/goleveldb v1.0.0 // indirect
	github.com/tendermint/go-amino v0.15.0
	github.com/tendermint/tendermint v0.31.5
	golang.org/x/crypto v0.0.0-20190611184440-5c40567a22f8 // indirect
	golang.org/x/net v0.0.0-20190613194153-d28f0bde5980 // indirect
	golang.org/x/sys v0.0.0-20190614160838-b47fdc937951 // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/appengine v1.4.0 // indirect
	google.golang.org/genproto v0.0.0-20190611190212-a7e196e89fd3 // indirect
	google.golang.org/grpc v1.21.1 // indirect
)

replace golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5

replace github.com/cosmos/sdk-application-tutorial/x/nameservice => ./x/nameservice
