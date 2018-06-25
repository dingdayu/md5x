package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const filechunk = 8192 // we settle for 8KB
const size = 1024 * 1.5 * 1024 * 1024

type FileI struct {
	Name     string
	FileInfo os.FileInfo
	Path     string
	MD5      string
	FileSize string
}

// 保存所有文件的MD5
var DirFiles []*FileI

// 命令参数
var (
	h      bool
	dir    string
	out    string
	export string
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")

	flag.StringVar(&dir, "dir", "", "需要计算MD5的文件路径")
	flag.StringVar(&out, "out", "", "导出重复文件的记录")
	flag.StringVar(&export, "export", "", "导出所有文件的MD5记录")
	// 改变默认的 Usage
	flag.Usage = usage
}

func main() {
	flag.Parse()

	if dir == "" {
		dir, _ = os.Getwd()
	}

	// 扫描文件夹
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() && strings.HasPrefix(f.Name(), ".") {
			return filepath.SkipDir
		}

		if !strings.HasPrefix(f.Name(), ".") {
			DirFiles = append(DirFiles,
				&FileI{
					Name:     f.Name(),
					FileInfo: f,
					Path:     path,
					FileSize: FormatFileSize(f.Size()),
				})
		}
		return nil
	})
	if err != nil {
		fmt.Errorf("Dir Scan Error: %v", err)
	}

	fmt.Printf("文件个数 %v\n", len(DirFiles))

	// 遍历文件MD5
	for i, value := range DirFiles {
		fmt.Printf("\r%d / %d, %d%%", i, len(DirFiles), i/len(DirFiles))

		if value.FileInfo.Size() > size {
			value.MD5, _, _ = md52(value.Path)
		} else {
			value.MD5, _, _ = md51(value.Path)
		}
	}

	if out != "" {
		// 导出所有MD5记录
		ExportMD5(DirFiles, out)
	}

	// 提取重复记录
	t := ExtractRepeat(DirFiles)

	if export != "" {
		// 导出所有MD5记录
		ExportRepeat(t, export)
	}

	// 遍历输出
	for _, value := range t {
		fmt.Print(value[0].MD5, "\t\t")
		for _, v := range value {
			fmt.Print(v.Path, "\t\t")
		}
		fmt.Println()
	}
}

// 提取重复元素
func ExtractRepeat(s []*FileI) map[string][]*FileI {
	m := map[string][]*FileI{}
	for _, v := range s {
		if _, ok := m[v.MD5]; ok {
			m[v.MD5] = append(m[v.MD5], v)
		} else {
			m[v.MD5] = []*FileI{v}
		}
	}

	for k, v := range m {
		if len(v) <= 1 {
			delete(m, k)
		}
	}
	return m
}

// 格式化文件大小
func FormatFileSize(fileBytes int64) string {
	var (
		units []string
		size  string
		i     int
	)
	units = []string{"B", "K", "M", "G", "T", "P"}
	i = 0
	for {
		i++
		fileBytes = fileBytes / 1024
		if fileBytes < 1024 {
			size = fmt.Sprintf("%d", fileBytes) + units[i]
			break
		}
	}
	return size
}

// \r or \f	回到行首
// fmt.Printf("\r%s%d%%", Bar(i/len(DirFiles), 15), i/len(DirFiles))
func Bar(vl int, width int) string {
	return fmt.Sprintf("%s%*c", strings.Repeat("█", vl/10), vl/10-width+1, ([]rune(" ▏▎▍▌▋▋▊▉█"))[vl%10])
}
