package main

import (
	"gopkg.in/yaml.v3"
	"io"
	"io/ioutil"
	"os"
)

type Configuration struct {
	Log            string `yaml:"log"`
	LogLevel       string `yaml:"log_level"`
	Proxy          string `yaml:"proxy"`
	UserAgent      string `yaml:"user_agent"`
	UseImageCache  bool   `yaml:"use_image_cache"`
	MaxCacheDbSize int    `yaml:"max_cache_db_size"`
	CacheDbDir     string `yaml:"cache_db_dir"`
	CacheImageDir  string `yaml:"cache_image_dir"`
}

func LoadConfig(file string) *Configuration {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	configuration := Configuration{
		Log:            "stdout",
		LogLevel:       "info",
		Proxy:          "",
		UserAgent:      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36",
		UseImageCache:  true,
		MaxCacheDbSize: 1024 * 1024 * 100, // 100MB
		CacheDbDir:     "cache/db",
		CacheImageDir:  "cache/image",
	}
	err = yaml.Unmarshal(data, &configuration)
	if err != nil {
		panic(err)
	}
	return &configuration
}

func (_this *Configuration) GetLogFile() io.WriteCloser {
	if _this.Log == "stdout" {
		return os.Stdout
	} else {
		logFile, err := os.Open(_this.Log)
		if err != nil {
			panic(err)
		}
		return logFile
	}
}

func (_this *Configuration) PrepareDirs() {
	_ = os.MkdirAll(_this.CacheDbDir, 0777)
	_ = os.MkdirAll(_this.CacheImageDir, 0777)
}