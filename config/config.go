package config

import (
	"flag"
	"time"
)

var cfg Config

func init() {
	flag.StringVar(&cfg.Addr, "addr", ":8080", "App addr")
	flag.UintVar(&cfg.RetryAttempts, "retry-attempts", 10, "")
	flag.DurationVar(&cfg.RetryDelay, "retry-delay", 200, "millisecond")

	flag.StringVar(&cfg.DB.URI, "db-uri", "postgresql://postgres:123456@localhost:5432/otc-crosschain?sslmode=disable", "")
	flag.StringVar(&cfg.DB.MigrationTable, "migration-table", "schema_migrations", "")

	flag.StringVar(&cfg.Eth.RpcURL, "eth-rpc-url", "https://eth-rinkeby.alchemyapi.io/v2/JnkdVOGnQ-uvjAl2jto1-V1OL3Vquq-H", "")
	flag.StringVar(&cfg.Eth.WsURL, "eth-ws-url", "wss://eth-rinkeby.alchemyapi.io/v2/JnkdVOGnQ-uvjAl2jto1-V1OL3Vquq-H", "")
	flag.DurationVar(&cfg.Eth.ReconnectInterval, "eth-ws-reconnect-interval", 20, "minutes")
	flag.StringVar(&cfg.Eth.SwapAddress, "eth-swap-addr", "0xd5a742aE5e0e98095F53a5Cb935824868A67bbd8", "")

	flag.StringVar(&cfg.Bsc.RpcURL, "bsc-rpc-url", "https://data-seed-prebsc-1-s1.binance.org:8545", "")
	flag.StringVar(&cfg.Bsc.WsURL, "bsc-ws-url", "<wss-url>", "")
	flag.DurationVar(&cfg.Bsc.ReconnectInterval, "ws-reconnect-interval", 20, "minutes")
	flag.StringVar(&cfg.Bsc.SwapAddress, "bsc-swap-addr", "0x58e661a14C0201c7611481397970b06249bD5cdb", "")

	flag.BoolVar(&cfg.IsNeedSync, "sync", true, "")

	flag.Parse()
}

type Config struct {
	Addr          string
	IsNeedSync    bool
	RetryAttempts uint
	RetryDelay    time.Duration

	DB  DB
	Eth Network
	Bsc Network
}

type DB struct {
	URI            string
	MigrationTable string
}

type Network struct {
	RpcURL            string
	WsURL             string
	ReconnectInterval time.Duration
	SwapAddress       string
}

func Get() Config {
	return cfg
}
