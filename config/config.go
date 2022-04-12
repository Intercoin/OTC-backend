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

	flag.StringVar(&cfg.DB.URI, "db-uri", "postgresql://postgres:123456@localhost:5432/otc?sslmode=disable", "")
	flag.StringVar(&cfg.DB.MigrationTable, "migration-table", "schema_migrations", "")

	flag.StringVar(&cfg.Eth.RpcURL, "rpc-url", "https://eth-rinkeby.alchemyapi.io/v2/JnkdVOGnQ-uvjAl2jto1-V1OL3Vquq-H", "")
	flag.StringVar(&cfg.Eth.WsURL, "ws-url", "wss://eth-rinkeby.alchemyapi.io/v2/JnkdVOGnQ-uvjAl2jto1-V1OL3Vquq-H", "")
	flag.DurationVar(&cfg.Eth.ReconnectInterval, "ws-reconnect-interval", 20, "minutes")
	flag.StringVar(&cfg.Eth.SwapAddress, "swap-addr", "<swap-addr>", "")

	flag.StringVar(&cfg.Bsc.RpcURL, "rpc-url", "<rpc-url>", "")
	flag.StringVar(&cfg.Bsc.WsURL, "ws-url", "<wss-url>", "")
	flag.DurationVar(&cfg.Bsc.ReconnectInterval, "ws-reconnect-interval", 20, "minutes")
	flag.StringVar(&cfg.Bsc.SwapAddress, "swap-addr", "<swap-addr>", "")

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
