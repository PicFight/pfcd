module github.com/picfight/pfcd/fees

require (
	github.com/btcsuite/goleveldb v1.0.0
	github.com/decred/slog v1.0.0
	github.com/golang/protobuf v1.2.0 // indirect
	github.com/jessevdk/go-flags v1.4.0
	golang.org/x/sys v0.0.0-20180831094639-fa5fdf94c789 // indirect
)

replace (
	github.com/picfight/pfcd/chaincfg => ../chaincfg
	github.com/picfight/pfcd/chaincfg/chainhash => ../chaincfg/chainhash
	github.com/picfight/pfcd/pfcec => ../pfcec
	github.com/picfight/pfcd/pfcec/edwards => ../pfcec/edwards
	github.com/picfight/pfcd/pfcec/secp256k1 => ../pfcec/secp256k1
	github.com/picfight/pfcd/pfcutil => ../pfcutil
	github.com/picfight/pfcd/wire => ../wire
)