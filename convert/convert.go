package convert

import (
	"fmt"
	"github.com/gingfrederik/docx"
	"os"
	"path/filepath"
)

// ConvertFileToDocx 将指定路径的文件内容转换为.docx文件
func ConvertFileToDocx(inputPath string, outputPath string) (bool, error) {
	// 读取输入文件内容
	inputContent, err := os.ReadFile(inputPath)
	if err != nil {
		return false, err
	}

	// 创建一个新的.docx文件
	doc := docx.NewFile()

	// 向文档中添加段落
	p := doc.AddParagraph()
	p.AddText(string(inputContent))

	//创建output所需文件夹
	dir := filepath.Dir(outputPath)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Printf("创建文件夹失败: %v\n", err)
	}

	// 保存为.docx文件
	err = doc.Save(outputPath)
	if err != nil {
		fmt.Printf("保存文件失败: %v\n", err)
		return false, err
	}

	return true, nil
}
