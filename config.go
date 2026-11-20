package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Preset 表示一条预填充替换规则
type Preset struct {
	Name   string `json:"name"`
	Before string `json:"before"`
	After  string `json:"after"`
}

// Config 是应用的可配置项
type Config struct {
	SupportedFormats []string `json:"supportedFormats"`
	Presets          []Preset `json:"presets"`
	Locale           string   `json:"locale"`
	Theme            string   `json:"theme"`
}

// defaultConfig 返回默认配置
func defaultConfig() Config {
	return Config{
		SupportedFormats: []string{".psd", ".jpg", ".png"},
		Presets: []Preset{
			{Name: "首图", Before: ".jpg", After: "_首图.jpg"},
		},
		Locale: "zh-Hans",
		Theme:  "light",
	}
}

// configPath 返回配置文件路径（用户配置目录下）
func configPath() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		dir = "."
	}
	return filepath.Join(dir, "buddybuddy-toolkit", "config.json")
}

// loadConfig 读取配置，文件不存在或解析失败时回退到默认配置
func loadConfig() Config {
	cfg := defaultConfig()
	data, err := os.ReadFile(configPath())
	if err != nil {
		return cfg
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return defaultConfig()
	}
	return cfg
}

// saveConfig 将配置写入文件
func saveConfig(cfg Config) error {
	path := configPath()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
