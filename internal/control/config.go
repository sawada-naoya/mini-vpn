// config.go: stub (Step0)
package control

import (
	"errors"
	"fmt"
	"net"
	"os"

	toml "github.com/pelletier/go-toml/v2"
)

// config.go では設定ファイルの読み込みとマッピングを行う
// TOML/YAMLで記述された設定を Go の構造体へ変換し、
// バリデーションを通して他レイヤーで安全に利用できるようにする
//
// このファイルに置くのは「設定ファイルに対応する構造体」と「その構造体をロード/検証する関数」のみ
// 実行時の状態管理やドメイン固有の構造体はここには置かない

type Config struct {
	Net      Net      `toml:"net"`
	Tun      Tun      `toml:"tun"`
	Security Security `toml:"security"`
	Session  Session  `toml:"session"`
	Metrics  Metrics  `toml:"metrics"`
}

type Net struct {
	Local string `toml:"local"` // e.g. ":51820" or "0.0.0.0:51820"
	Peer  string `toml:"peer"`  // e.g. "203.0.113.10:51820"
}

type Tun struct {
	Name   string `toml:"name"`    // tun0
	IPCidr string `toml:"ip_cidr"` // "10,10,0,1/24"
	MTU    int    `toml:"mtu"`     // 1300
}

type Security struct {
	PSKFile string `toml:"psk_file"` // Path to hex-encoded 32-byte key
}

type Session struct {
	KeepaliveSec int `toml:"keepalive_sec"`
	TimeoutSec   int `toml:"timeout_sec"`
}

type Metrics struct {
	Listen string `toml:"listen"`
}

func LoadConfig(path string) (*Config, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config %s: %w", path, err)
	}

	var cfg Config
	if err := toml.Unmarshal(raw, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse toml %s: %w", path, err)
	}

	if err := cfg.Validate(path); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (c *Config) Validate(cfgPath string) error {
	// net
	if _, _, err := net.SplitHostPort(normalizeAdr(c.Net.Local)); err != nil {
		return fmt.Errorf("net.local invalid: %w", err)
	}
	if c.Net.Peer == "" {
		return errors.New("net.peer required")
	}
	if _, _, err := net.SplitHostPort(c.Net.Peer); err != nil {
		return fmt.Errorf("net.peer invalid: %v", err)
	}

	return nil
}

func normalizeAdr(adr string) string {
	println("dammy")
	return adr
}
