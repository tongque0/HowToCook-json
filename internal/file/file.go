package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/duke-git/lancet/v2/strutil"
)

func GetAllRecipePaths(root string) ([]string, error) {
	var paths []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".md" {
			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			paths = append(paths, absPath)
		}
		return nil
	})
	return paths, err
}

func GetRecipeImageName(path string) ([]string, error) {
	dir := filepath.Dir(path)
	entries, err := os.ReadDir(filepath.Join(dir))
	if err != nil {
		return nil, nil
	}

	var files []string

	// 遍历文件列表
	for _, entry := range entries {
		if !entry.IsDir() {
			fileName := entry.Name()
			// 因为目录中只会有md文件和菜品图片，所以只需要判断不是.md文件即可
			if !strings.HasSuffix(fileName, ".md") {

				files = append(files, fileName)
			}
		}
	}

	return files, nil
}
func GetRecipeGetCategory(path string) string {
	dir := filepath.Dir(path)

	normalizedPath := strings.ReplaceAll(dir, `\`, `/`)
	basepath := strutil.After(normalizedPath, "/dishes/")
	categorypath := strutil.Before(basepath, "/")
	fmt.Println(categorypath)
	return categorypath
}
func GetGithubImagePath(path string) ([]string, error) {
	dir := filepath.Dir(path)

	normalizedPath := strings.ReplaceAll(dir, `\`, `/`)
	basepath := "https://raw.githubusercontent.com/Anduin2017/HowToCook/refs/heads/master/dishes/" + strutil.After(normalizedPath, "/dishes/")

	dishes, err := GetRecipeImageName(path)
	if err != nil {
		return nil, err
	}

	var githubPaths []string
	for _, dish := range dishes {
		githubPath := basepath + "/" + dish + "?raw=true"
		githubPaths = append(githubPaths, githubPath)
	}

	return githubPaths, nil
}

func SaveRecipesToJson(path string, T any) {

	file, err := os.Create(path)
	if err != nil {
		return
	}
	defer file.Close()

	TJson, err := convertor.ToJson(T)
	if err != nil {
		return
	}

	err = fileutil.WriteStringToFile(path, TJson, false)
	if err != nil {
		return
	}
}
