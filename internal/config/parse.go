package config

import (
	"os"
	"path/filepath"
	"slices"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/creasty/defaults"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

var (
	once sync.Once
	C    = new(Config)
)

func MustLoad(configDir string, names ...string) {
	once.Do(func() {
		if err := load(configDir, names...); err != nil {
			panic(err)
		}
	})
}

func load(configDir string, names ...string) error {
	if err := defaults.Set(C); err != nil {
		return err
	}

	supportExts := []string{".json", ".toml"}

	parseFile := func(name string) error {
		ext := filepath.Ext(name)
		if ext == "" || !slices.Contains(supportExts, ext) {
			return nil
		}
		buf, err := os.ReadFile(name)
		if err != nil {
			return errors.Wrapf(err, "读取配置文件失败: %s", name)
		}

		switch ext {
		case ".json":
			err = jsoniter.Unmarshal(buf, C)
		case ".toml":
			err = toml.Unmarshal(buf, C)
		}
		return errors.Wrapf(err, "处理配置文件失败 %s", name)
	}

	for _, name := range names {
		fullname := filepath.Join(configDir, name)

		stat, err := os.Stat(fullname)
		if err != nil {
			return errors.Wrapf(err, "配置文件不存在: %s", fullname)
		}

		if stat.IsDir() {
			err := filepath.WalkDir(fullname, func(path string, d os.DirEntry, err error) error {
				if err != nil {
					return err
				}

				if d.IsDir() {
					return nil
				}

				return parseFile(path)
			})

			if err != nil {
				return errors.Wrapf(err, "遍历配置文件路径失败: %s", fullname)
			}

			continue
		}

		if err := parseFile(fullname); err != nil {
			return err
		}
	}

	return nil
}
