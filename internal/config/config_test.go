package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadWithAutoDiscovery(t *testing.T) {
	// 创建临时测试目录
	tempDir, err := os.MkdirTemp("", "config-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// 创建测试配置文件
	setupTestConfigs(t, tempDir)

	// 重置全局配置
	C = new(Config)
	once.Do(func() {})

	// 测试自动加载
	err = load(tempDir, "dev")
	if err != nil {
		t.Fatalf("自动加载配置失败: %v", err)
	}

	// 验证配置是否正确加载
	if C.General.AppName != "code-push-server" {
		t.Errorf("期望 AppName 为 'code-push-server', 实际为 '%s'", C.General.AppName)
	}

	if C.General.Debug != true {
		t.Errorf("期望 Debug 为 true, 实际为 %v", C.General.Debug)
	}

	// 验证 shared 配置是否加载
	if C.Database.Driver != "sqlite" {
		t.Errorf("期望 Database.Driver 为 'sqlite', 实际为 '%s'", C.Database.Driver)
	}

	// 验证 local 配置是否覆盖
	if C.General.HTTP.Addr != ":8042" {
		t.Errorf("期望 HTTP.Addr 为 ':8042' (local 覆盖), 实际为 '%s'", C.General.HTTP.Addr)
	}

	if C.Developer.Name != "测试开发者" {
		t.Errorf("期望 Developer.Name 为 '测试开发者', 实际为 '%s'", C.Developer.Name)
	}
}

func setupTestConfigs(t *testing.T, tempDir string) {
	// 创建目录结构
	dirs := []string{
		"shared",
		"local",
	}
	for _, dir := range dirs {
		err := os.MkdirAll(filepath.Join(tempDir, dir), 0755)
		if err != nil {
			t.Fatal(err)
		}
	}

	// 创建环境配置文件
	envConfig := `[General]
AppName = "code-push-server"
Version = "v0.0.1"
Debug = true

[General.HTTP]
Addr = ":8040"
`
	err := os.WriteFile(filepath.Join(tempDir, "config.dev.toml"), []byte(envConfig), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// 创建 shared 配置文件
	sharedConfig := `[Database]
Driver = "sqlite"
DSN = "data/test.db"
MaxOpenConns = 10

[Database.Migration]
AutoMigrate = true
MigrationPath = "migrations"
`
	err = os.WriteFile(filepath.Join(tempDir, "shared", "database.toml"), []byte(sharedConfig), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// 创建 local 配置文件
	localConfig := `[General.HTTP]
Addr = ":8042"

[Developer]
Name = "测试开发者"
Email = "test@example.com"
EnableDebugRoutes = true
`
	err = os.WriteFile(filepath.Join(tempDir, "local", "override.toml"), []byte(localConfig), 0644)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetConfigFiles(t *testing.T) {
	// 创建临时测试目录
	tempDir, err := os.MkdirTemp("", "config-files-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// 创建测试文件
	files := []string{
		"database.toml",
		"security.toml",
		"cache.json",
		"readme.txt", // 应该被忽略
	}
	for _, file := range files {
		err := os.WriteFile(filepath.Join(tempDir, file), []byte("test"), 0644)
		if err != nil {
			t.Fatal(err)
		}
	}

	// 测试获取配置文件
	configFiles, err := getConfigFiles(tempDir)
	if err != nil {
		t.Fatalf("获取配置文件失败: %v", err)
	}

	expected := []string{"cache.json", "database.toml", "security.toml"}
	if len(configFiles) != len(expected) {
		t.Fatalf("期望 %d 个配置文件, 实际为 %d 个", len(expected), len(configFiles))
	}

	for i, file := range configFiles {
		if file != expected[i] {
			t.Errorf("期望文件 '%s', 实际为 '%s'", expected[i], file)
		}
	}
}
