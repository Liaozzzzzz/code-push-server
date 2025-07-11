package config

import (
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"
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

func MustLoad(configDir string, env string) {
	once.Do(func() {
		if err := load(configDir, env); err != nil {
			panic(err)
		}
	})
}

// load 自动加载配置文件，包括 shared 和 local 目录
func load(configDir string, env string) error {
	if err := defaults.Set(C); err != nil {
		return err
	}

	var loadOrder []string

	// 1. 首先加载环境配置文件
	if env != "" {
		envConfig := "config." + env + ".toml"
		if _, err := os.Stat(filepath.Join(configDir, envConfig)); err == nil {
			loadOrder = append(loadOrder, envConfig)
		}
	}

	// 2. 然后加载 shared 目录下的所有配置文件
	sharedDir := filepath.Join(configDir, "shared")
	if stat, err := os.Stat(sharedDir); err == nil && stat.IsDir() {
		sharedFiles, err := getConfigFiles(sharedDir)
		if err != nil {
			return errors.Wrapf(err, "获取 shared 目录配置文件失败")
		}
		for _, file := range sharedFiles {
			loadOrder = append(loadOrder, "shared/"+file)
		}
	}

	// 3. 最后加载 local 目录下的所有配置文件
	localDir := filepath.Join(configDir, "local")
	if stat, err := os.Stat(localDir); err == nil && stat.IsDir() {
		localFiles, err := getConfigFiles(localDir)
		if err != nil {
			return errors.Wrapf(err, "获取 local 目录配置文件失败")
		}
		for _, file := range localFiles {
			loadOrder = append(loadOrder, "local/"+file)
		}
	}

	// 按顺序加载配置文件
	return loadFiles(configDir, loadOrder...)
}

// getConfigFiles 获取目录下的所有配置文件，并按文件名排序
func getConfigFiles(dir string) ([]string, error) {
	var files []string
	supportExts := []string{".json", ".toml"}

	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		ext := filepath.Ext(d.Name())
		if slices.Contains(supportExts, ext) {
			relPath, err := filepath.Rel(dir, path)
			if err != nil {
				return err
			}
			files = append(files, relPath)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// 按文件名排序，确保加载顺序一致
	sort.Strings(files)
	return files, nil
}

// loadFiles 加载指定的配置文件列表
func loadFiles(configDir string, names ...string) error {
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
			// 对于 local 目录的文件，如果不存在则跳过（不是错误）
			if strings.HasPrefix(name, "local/") {
				continue
			}
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
