package main

import (
	"fmt"
	"os"

	"github.com/tongque0/HowToCook-json/file"
	"github.com/tongque0/HowToCook-json/recipe"
	simpletype "github.com/tongque0/HowToCook-json/recipe/simpleType"
)

var simpleRecipes []simpletype.Simpletype
var simpleParse recipe.ParseRecipe[simpletype.Simpletype] = &simpletype.Simpletype{}

func main() {

	if len(os.Args) < 3 {
		fmt.Println("请指定菜谱文件夹路径和输出文件夹路径")
		return
	}
	dishDir := os.Args[1]
	output := os.Args[2]
	//查找
	recipePaths, err := file.GetAllRecipePaths(dishDir)
	if err != nil {
		fmt.Printf("%v", err)
	}

	//解析
	for _, path := range recipePaths {
		simpleRecipe, _ := simpleParse.FileToRecipeType(path)
		simpleRecipes = append(simpleRecipes, simpleRecipe)
	}

	//保存
	file.SaveRecipesToJson(output+"/simpleType.json", simpleRecipes)

}
