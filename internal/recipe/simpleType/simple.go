package simpletype

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tongque0/HowToCook-json/file"
)

// type Calculation struct {
// 	Item   string `json:"item"`   //名称
// 	Amount string `json:"amount"` //用量
// }

type Simpletype struct {
	Title        string   `json:"title"`        //名称
	Category     string   `json:"category"`     //分类
	Image        []string `json:"image"`        //图片
	Introduction string   `json:"introduction"` //介绍
	Difficulty   string   `json:"difficulty"`   //难度
	Ingredients  []string `json:"ingredients"`  //原料
	Calculations []string `json:"calculations"` //计算
	Steps        []string `json:"steps"`        //步骤
	Notes        string   `json:"notes"`        //附加内容
}

func (r *Simpletype) FileToRecipeType(name string) (Simpletype, error) {
	recipeContent, err := os.ReadFile(name)
	if err != nil {
		log.Fatalf("无法读取文件: %v", err)
	}
	simpleRecipe := parseMarkdown(recipeContent)

	simpleRecipe.Category = file.GetRecipeGetCategory(name)
	fmt.Println(simpleRecipe.Category)
	simpleRecipe.Image, err = file.GetGithubImagePath(name)
	if err != nil {
		log.Fatalf("无法获取图片: %v", err)
	}

	return simpleRecipe, nil
}

func parseMarkdown(recipeContent []byte) Simpletype {
	var recipe Simpletype
	section := ""
	lines := strings.Split(string(recipeContent), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(line, "# "):
			recipe.Title = strings.TrimPrefix(line, "# ")
			recipe.Title = strings.Replace(recipe.Title, "的做法", "", -1)
			section = "introduction"
		case strings.HasPrefix(line, "!["):
			break
		case strings.HasPrefix(line, "预估烹饪难度："):
			recipe.Difficulty = strings.TrimPrefix(line, "预估烹饪难度：")
		case line == "## 必备原料和工具":
			section = "ingredients"
		case line == "## 计算":
			section = "calculations"
		case line == "## 操作":
			section = "steps"
		case line == "## 附加内容":
			section = "notes"
		case strings.HasPrefix(line, "如果您遵循本指南的制作流程而发现有问题或可以改进的流程，请提出 Issue 或 Pull request"):
			break
		case strings.HasPrefix(line, "<!--"):
			break
		default:
			switch section {
			case "introduction":
				if line != "" {
					recipe.Introduction += line
				}
			case "ingredients":
				if strings.HasPrefix(line, "- ") {
					ingredient := strings.TrimPrefix(line, "- ")
					recipe.Ingredients = append(recipe.Ingredients, ingredient)
				}
			case "calculations":
				if strings.HasPrefix(line, "- ") {
					calculations := strings.TrimPrefix(line, "- ")
					recipe.Calculations = append(recipe.Calculations, calculations)
				}
			case "steps":
				if strings.HasPrefix(line, "- ") {
					step := strings.TrimPrefix(line, "- ")
					recipe.Steps = append(recipe.Steps, step)
				}
			case "notes":
				if line != "" {
					recipe.Notes += line
				}
			}
		}
	}

	return recipe
}
