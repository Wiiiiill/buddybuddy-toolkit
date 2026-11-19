package main

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
// NewApp 创建一个新的 App 应用程序
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
// startup 在应用程序启动时调用
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	// 在这里执行初始化设置
	a.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
// domReady 在前端Dom加载完毕后调用
func (a *App) domReady(ctx context.Context) {
	// Add your action here
	// 在这里添加你的操作
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue,
// false will continue shutdown as normal.
// beforeClose在单击窗口关闭按钮或调用runtime.Quit即将退出应用程序时被调用.
// 返回 true 将导致应用程序继续，false 将继续正常关闭。
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
// 在应用程序终止时被调用
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
	// 在此处做一些资源释放的操作
}
func alertErr(ctx context.Context, err error) {
	runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:    runtime.ErrorDialog,
		Title:   "something wrong",
		Message: err.Error(),
	})
}
func alertSuccess(ctx context.Context) {
	runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Title:   "Info",
		Message: "Done"})
}

// isImageFile 判断文件名是否为支持的图片格式（大小写不敏感）
func isImageFile(name string) bool {
	lower := strings.ToLower(name)
	return strings.HasSuffix(lower, ".psd") ||
		strings.HasSuffix(lower, ".jpg") ||
		strings.HasSuffix(lower, ".png")
}

func (a *App) OpenFolder() string {

	// runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
	// 	Type:  runtime.WarningDialog,
	// 	Title: "asd",
	// })

	// runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
	// 	Type:  runtime.QuestionDialog,
	// 	Title: "asd",
	// })

	directory, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		alertErr(a.ctx, err)
		// 处理错误
		return ""
	}
	// 返回选择的目录
	return directory
}

func (a *App) CopyFilesBasedOnPrefix(path string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		alertErr(a.ctx, err)
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		if !isImageFile(fileName) {
			continue
		}

		// 按下划线取前缀作为子目录名；没有下划线或下划线在开头时无法归档，跳过
		idx := strings.Index(fileName, "_")
		if idx <= 0 {
			continue
		}
		prefix := fileName[:idx]

		newDir := filepath.Join(path, prefix)
		if _, err := os.Stat(newDir); os.IsNotExist(err) {
			if err := os.MkdirAll(newDir, 0755); err != nil {
				alertErr(a.ctx, err)
				return err
			}
		}

		oldFile := filepath.Join(path, fileName)
		newFile := filepath.Join(newDir, fileName)
		if err := os.Rename(oldFile, newFile); err != nil {
			alertErr(a.ctx, err)
			return err
		}
		fmt.Printf("Moved: %s -> %s\n", oldFile, newFile)
	}
	alertSuccess(a.ctx)
	return nil
}

func (a *App) DecodeFilenames(path string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		alertErr(a.ctx, err)
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		if !isImageFile(fileName) {
			continue
		}

		newName, err := url.QueryUnescape(fileName)
		if err != nil {
			alertErr(a.ctx, err)
			return err
		}
		// 解码后无变化的文件无需重命名，避免 rename 到自身报错
		if newName == fileName {
			continue
		}

		oldFile := filepath.Join(path, fileName)
		newFile := filepath.Join(path, newName)
		if err := os.Rename(oldFile, newFile); err != nil {
			alertErr(a.ctx, err)
			return err
		}
		fmt.Printf("Renamed: %s -> %s\n", oldFile, newFile)
	}

	alertSuccess(a.ctx)
	return nil
}

func (a *App) ReplaceFilenames(path, before, after string) error {
	// before 为空时替换无意义，直接返回
	if before == "" {
		return nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		alertErr(a.ctx, err)
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		if !isImageFile(fileName) || !strings.Contains(fileName, before) {
			continue
		}

		newName := strings.ReplaceAll(fileName, before, after)
		if newName == fileName {
			continue
		}

		oldFile := filepath.Join(path, fileName)
		newFile := filepath.Join(path, newName)
		if err := os.Rename(oldFile, newFile); err != nil {
			alertErr(a.ctx, err)
			return err
		}
		fmt.Printf("Replaced: %s -> %s\n", oldFile, newFile)
	}

	alertSuccess(a.ctx)
	return nil
}
