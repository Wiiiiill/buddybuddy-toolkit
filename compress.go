package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ImageFileInfo 描述一个待压缩的图片文件
type ImageFileInfo struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
}

// isCompressibleImage 判断是否为可压缩的图片格式（仅 JPG/PNG，PSD 等跳过）
func isCompressibleImage(name string) bool {
	switch strings.ToLower(filepath.Ext(name)) {
	case ".jpg", ".jpeg", ".png":
		return true
	}
	return false
}

// ListImageFiles 列出目录中超过目标大小的 JPG/PNG（不递归，跳过子目录）
func (a *App) ListImageFiles(path string) ([]ImageFileInfo, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	target := a.config.CompressTargetBytes
	result := make([]ImageFileInfo, 0)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if !isCompressibleImage(name) {
			continue
		}
		info, err := entry.Info()
		if err != nil {
			return nil, err
		}
		if info.Size() <= target {
			continue
		}
		result = append(result, ImageFileInfo{Name: name, Size: info.Size()})
	}
	return result, nil
}

// ReadImageBase64 读取图片文件并返回 base64，供前端 wasm 编码器使用
func (a *App) ReadImageBase64(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

// ReplaceImageBase64 用压缩结果原子替换原图：
// 先写入同目录临时文件，成功后再 rename 覆盖，避免中途失败损坏原图。
// 返回替换后的文件大小；若结果没有比原图更小则返回错误、不替换。
func (a *App) ReplaceImageBase64(path string, dataBase64 string) (int64, error) {
	if !isCompressibleImage(path) {
		return 0, fmt.Errorf("不支持的文件类型: %s", filepath.Base(path))
	}

	data, err := base64.StdEncoding.DecodeString(dataBase64)
	if err != nil {
		return 0, err
	}

	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	if int64(len(data)) >= info.Size() {
		return info.Size(), fmt.Errorf("压缩结果没有更小，已保留原图")
	}

	dir := filepath.Dir(path)
	tmp, err := os.CreateTemp(dir, ".compress-*.tmp")
	if err != nil {
		return 0, err
	}
	tmpName := tmp.Name()

	committed := false
	defer func() {
		tmp.Close()
		if !committed {
			os.Remove(tmpName)
		}
	}()

	if _, err := tmp.Write(data); err != nil {
		return 0, err
	}
	if err := tmp.Sync(); err != nil {
		return 0, err
	}
	if err := tmp.Close(); err != nil {
		return 0, err
	}
	if err := os.Rename(tmpName, path); err != nil {
		return 0, err
	}
	committed = true
	return int64(len(data)), nil
}
