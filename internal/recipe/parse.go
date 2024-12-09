package recipe

type ParseRecipe[T any] interface {
	FileToRecipeType(filepath string) (T, error)
}
