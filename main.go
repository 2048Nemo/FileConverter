package main

import (
	"demo/v1/convert"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// 检查命令行参数数量
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_directory>")
		os.Exit(1)
	}

	// 获取输入文件夹路径
	inputPath := os.Args[1]

	// 检查是否指定了目录
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		fmt.Printf("Directory not found: %s\n", inputPath)
		os.Exit(1)
	}

	// 创建输出文件夹
	outputPath := inputPath + "_docx"
	if err := os.MkdirAll(outputPath, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// 开始递归遍历目录
	err := filepath.WalkDir(inputPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// 检查是否为 .java 文件
		if strings.HasSuffix(d.Name(), ".java") {
			inputFile := path
			relativePath, err := filepath.Rel(inputPath, inputFile)
			if err != nil {
				return err
			}
			outputFile := filepath.Join(outputPath, strings.Replace(relativePath, ".java", ".docx", 1))

			// 检查是否为.java文件，如果不是则转换
			if inputFile != outputFile {
				fmt.Printf("Converting %s to %s ...\n", inputFile, outputFile)
				if ok, err := convert.ConvertFileToDocx(inputFile, outputFile); !ok {
					log.Printf("Failed to convert %s: %v", inputFile, err)
				}
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("File conversion successful.")
}
