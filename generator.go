package zerswag

import (
	"os"
	"path/filepath"
	"strings"
)


type JsonPath struct {
	JsonFile     string `json:"url"`
	Name         string `json:"name"`
	RealFileName string
	FullPath     string `json:"-"` // 完整文件路径，不序列化到JSON
}

func GenerateApi(docPath string, basePath string) []JsonPath {
	if basePath == "" {
		basePath = "."
	}
	exeDir := filepath.Dir(basePath)
	files, _ := findFilesWithPattern(exeDir, "*api.json")

	var resultMap []JsonPath
	for _, jsonPath := range files {
		// 从完整路径中提取文件名
		fileName := filepath.Base(jsonPath)
		// 移除 .json 扩展名作为显示名称
		displayName := strings.TrimSuffix(fileName, ".json")

		resultMap = append(resultMap, JsonPath{
			JsonFile:     docPath + "/api-" + fileName,
			Name:         displayName,
			RealFileName: fileName,
			FullPath:     jsonPath,
		})
	}

	return resultMap
}


func findFilesWithExt(rootDir, ext string) ([]string, error) {
	var result []string
	err := filepath.WalkDir(rootDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(path) == ext {
			result = append(result, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func findFilesWithPattern(rootDir, pattern string) ([]string, error) {
	var result []string
	err := filepath.WalkDir(rootDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			matched, err := filepath.Match(pattern, d.Name())
			if err != nil {
				return err
			}
			if matched {
				result = append(result, path)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
