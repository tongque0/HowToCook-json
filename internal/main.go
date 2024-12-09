package main

import (
	"fmt"

	"github.com/tongque0/HowToCook-json/file"
	"github.com/tongque0/HowToCook-json/recipe"
	simpletype "github.com/tongque0/HowToCook-json/recipe/simpleType"
)

var simpleRecipes []simpletype.Simpletype
var simpleParse recipe.ParseRecipe[simpletype.Simpletype] = &simpletype.Simpletype{}

func main() {
	//查找
	recipePaths, err := file.GetAllRecipePaths("./dishes")
	if err != nil {
		fmt.Printf("%v", err)
	}

	//解析
	for _, path := range recipePaths {
		simpleRecipe, _ := simpleParse.FileToRecipeType(path)
		simpleRecipes = append(simpleRecipes, simpleRecipe)
	}

	//保存
	file.SaveRecipesToJson("./simpleType.json", simpleRecipes)

}
