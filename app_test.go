package main

import (
	"os"
	"path/filepath"
	"testing"
)

func testApp() *App {
	return &App{config: defaultConfig()}
}

func writeFile(t *testing.T, dir, name string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, name), []byte("test"), 0644); err != nil {
		t.Fatal(err)
	}
}

func checkExists(t *testing.T, path string) {
	t.Helper()
	if _, err := os.Stat(path); err != nil {
		t.Errorf("expected file to exist: %s", path)
	}
}

func TestCopyFilesBasedOnPrefix(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "首图_001.jpg")   // 有下划线，应归档
	writeFile(t, dir, "plain.jpg")       // 无下划线，应跳过且不 panic
	writeFile(t, dir, "属性_002.PNG")    // 大写后缀，应处理

	app := testApp()
	count, err := app.CopyFilesBasedOnPrefix(dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if count != 2 {
		t.Errorf("want 2, got %d", count)
	}
	checkExists(t, filepath.Join(dir, "首图", "首图_001.jpg"))
	checkExists(t, filepath.Join(dir, "属性", "属性_002.PNG"))
	checkExists(t, filepath.Join(dir, "plain.jpg"))
}

func TestDecodeFilenames(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "a%20b.jpg") // URL 编码空格，应解码为 "a b.jpg"
	writeFile(t, dir, "plain.jpg")  // 无编码，应跳过

	app := testApp()
	count, err := app.DecodeFilenames(dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if count != 1 {
		t.Errorf("want 1, got %d", count)
	}
	checkExists(t, filepath.Join(dir, "a b.jpg"))
	checkExists(t, filepath.Join(dir, "plain.jpg"))
}

func TestReplaceFilenames(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "abc.jpg")
	writeFile(t, dir, "xyz.png")

	app := testApp()

	// 空 before 应直接返回 0
	count, err := app.ReplaceFilenames(dir, "", "zzz")
	if err != nil || count != 0 {
		t.Errorf("empty before: want 0/nil, got %d/%v", count, err)
	}

	// 正常替换 .jpg -> _首图.jpg
	count, err = app.ReplaceFilenames(dir, ".jpg", "_首图.jpg")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if count != 1 {
		t.Errorf("want 1, got %d", count)
	}
	checkExists(t, filepath.Join(dir, "abc_首图.jpg"))
	checkExists(t, filepath.Join(dir, "xyz.png"))
}
