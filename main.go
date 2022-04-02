package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/manifoldco/promptui"
)

type FoodCategories struct {
	Categories []struct {
		IDCategory             string `json:"idCategory"`
		StrCategory            string `json:"strCategory"`
		StrCategoryThumb       string `json:"strCategoryThumb"`
		StrCategoryDescription string `json:"strCategoryDescription"`
	} `json:"categories"`
}

type Meals struct {
	Meals []struct {
		StrMeal      string `json:"strMeal"`
		StrMealThumb string `json:"strMealThumb"`
		IDMeal       string `json:"idMeal"`
	} `json:"meals"`
}

type MealsInfo struct {
	MealName string `json:"mealName"`
	MealId   string `json:"mealId"`
}

type Recipe struct {
	Meals []struct {
		IDMeal                      string `json:"idMeal"`
		StrMeal                     string `json:"strMeal"`
		StrDrinkAlternate           string `json:"strDrinkAlternate"`
		StrCategory                 string `json:"strCategory"`
		StrArea                     string `json:"strArea"`
		StrInstructions             string `json:"strInstructions"`
		StrMealThumb                string `json:"strMealThumb"`
		StrTags                     string `json:"strTags"`
		StrYoutube                  string `json:"strYoutube"`
		StrIngredient1              string `json:"strIngredient1"`
		StrIngredient2              string `json:"strIngredient2"`
		StrIngredient3              string `json:"strIngredient3"`
		StrIngredient4              string `json:"strIngredient4"`
		StrIngredient5              string `json:"strIngredient5"`
		StrIngredient6              string `json:"strIngredient6"`
		StrIngredient7              string `json:"strIngredient7"`
		StrIngredient8              string `json:"strIngredient8"`
		StrIngredient9              string `json:"strIngredient9"`
		StrIngredient10             string `json:"strIngredient10"`
		StrIngredient11             string `json:"strIngredient11"`
		StrIngredient12             string `json:"strIngredient12"`
		StrIngredient13             string `json:"strIngredient13"`
		StrIngredient14             string `json:"strIngredient14"`
		StrIngredient15             string `json:"strIngredient15"`
		StrIngredient16             string `json:"strIngredient16"`
		StrIngredient17             string `json:"strIngredient17"`
		StrIngredient18             string `json:"strIngredient18"`
		StrIngredient19             string `json:"strIngredient19"`
		StrIngredient20             string `json:"strIngredient20"`
		StrMeasure1                 string `json:"strMeasure1"`
		StrMeasure2                 string `json:"strMeasure2"`
		StrMeasure3                 string `json:"strMeasure3"`
		StrMeasure4                 string `json:"strMeasure4"`
		StrMeasure5                 string `json:"strMeasure5"`
		StrMeasure6                 string `json:"strMeasure6"`
		StrMeasure7                 string `json:"strMeasure7"`
		StrMeasure8                 string `json:"strMeasure8"`
		StrMeasure9                 string `json:"strMeasure9"`
		StrMeasure10                string `json:"strMeasure10"`
		StrMeasure11                string `json:"strMeasure11"`
		StrMeasure12                string `json:"strMeasure12"`
		StrMeasure13                string `json:"strMeasure13"`
		StrMeasure14                string `json:"strMeasure14"`
		StrMeasure15                string `json:"strMeasure15"`
		StrMeasure16                string `json:"strMeasure16"`
		StrMeasure17                string `json:"strMeasure17"`
		StrMeasure18                string `json:"strMeasure18"`
		StrMeasure19                string `json:"strMeasure19"`
		StrMeasure20                string `json:"strMeasure20"`
		StrSource                   string `json:"strSource"`
		StrImageSource              string `json:"strImageSource"`
		StrCreativeCommonsConfirmed string `json:"strCreativeCommonsConfirmed"`
		DateModified                string `json:"dateModified"`
	} `json:"meals"`
}

func getFoodCategories() FoodCategories {
	var categories FoodCategories
	res, err := http.Get("https://themealdb.com/api/json/v1/1/categories.php")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		log.Fatalf("Failed to fetch data with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(body, &categories); err != nil { // Parse []byte to the go struct pointer
		log.Fatal("unable to parse JSON")
	}
	return categories
}

func makeCategoriesNameArr(catgories FoodCategories) []string {
	var arr []string
	for i := range catgories.Categories {
		arr = append(arr, catgories.Categories[i].StrCategory)
	}

	return arr
}

func getMealsInCategory(category string) Meals {
	var meals Meals
	url := fmt.Sprintf("https://themealdb.com/api/json/v1/1/filter.php?c=%s", category)
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		log.Fatalf("Failed to fetch data with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(body, &meals); err != nil { // Parse []byte to the go struct pointer
		log.Fatal("unable to parse JSON")
	}
	return meals
}

func makeMealsInCategoryArr(meals Meals) []MealsInfo {
	var arr []MealsInfo

	for i := range meals.Meals {
		details := MealsInfo{meals.Meals[i].StrMeal, meals.Meals[i].IDMeal}
		arr = append(arr, details)
	}

	return arr
}

func getMealId(selectedMealName string, mealsArr []MealsInfo) string {

	var id string
	for i := range mealsArr {
		if mealsArr[i].MealName == selectedMealName {
			id = mealsArr[i].MealId
		}
	}

	return id
}

func getRecipeById(id string) Recipe {
	var recipe Recipe

	url := fmt.Sprintf("https://themealdb.com/api/json/v1/1/lookup.php?i=%s", id)
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		log.Fatalf("Failed to fetch data with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(body, &recipe); err != nil { // Parse []byte to the go struct pointer
		log.Fatal("unable to parse JSON")
	}
	return recipe
}

func displayRecipeDetails(recipe Recipe) {
	instructions := recipe.Meals[0].StrInstructions
	ingredientsArr := []string{recipe.Meals[0].StrIngredient1, recipe.Meals[0].StrIngredient2, recipe.Meals[0].StrIngredient3, recipe.Meals[0].StrIngredient4, recipe.Meals[0].StrIngredient5, recipe.Meals[0].StrIngredient6, recipe.Meals[0].StrIngredient7, recipe.Meals[0].StrIngredient8, recipe.Meals[0].StrIngredient9, recipe.Meals[0].StrIngredient10, recipe.Meals[0].StrIngredient11, recipe.Meals[0].StrIngredient12, recipe.Meals[0].StrIngredient13, recipe.Meals[0].StrIngredient14, recipe.Meals[0].StrIngredient15, recipe.Meals[0].StrIngredient16, recipe.Meals[0].StrIngredient17, recipe.Meals[0].StrIngredient18, recipe.Meals[0].StrIngredient19, recipe.Meals[0].StrIngredient20}
	measuresArr := []string{recipe.Meals[0].StrMeasure1, recipe.Meals[0].StrMeasure2, recipe.Meals[0].StrMeasure3, recipe.Meals[0].StrMeasure4, recipe.Meals[0].StrMeasure5, recipe.Meals[0].StrMeasure6, recipe.Meals[0].StrMeasure7, recipe.Meals[0].StrMeasure8, recipe.Meals[0].StrMeasure9, recipe.Meals[0].StrMeasure10, recipe.Meals[0].StrMeasure11, recipe.Meals[0].StrMeasure12, recipe.Meals[0].StrMeasure13, recipe.Meals[0].StrMeasure14, recipe.Meals[0].StrMeasure15, recipe.Meals[0].StrMeasure16, recipe.Meals[0].StrMeasure17, recipe.Meals[0].StrMeasure18, recipe.Meals[0].StrMeasure19, recipe.Meals[0].StrMeasure20}

	type IngredientsStruct struct {
		Name    string
		Measure string
	}
	var usedIngredients []IngredientsStruct
	for i := range ingredientsArr {
		if ingredientsArr[i] != "" {
			ingredient := IngredientsStruct{ingredientsArr[i], measuresArr[i]}
			usedIngredients = append(usedIngredients, ingredient)
		}
	}

	fmt.Printf("Recipe: %s\n\n", recipe.Meals[0].StrMeal)
	fmt.Printf("YouTube Link: %s\n\n", recipe.Meals[0].StrYoutube)
	fmt.Printf("Ingredients:\n")
	for i := range usedIngredients {
		fmt.Printf("\t%s:\t%s\n", usedIngredients[i].Name, usedIngredients[i].Measure)
	}
	fmt.Printf("\nInstructions:\n\n%s", instructions)

}

func main() {
	// GET CATEGORIES
	categoriesRes := getFoodCategories()
	categoriesNamesArr := makeCategoriesNameArr(categoriesRes)

	catgPrompt := promptui.Select{
		Label: "Select a food category:",
		Items: categoriesNamesArr,
	}

	_, catgResult, err := catgPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// GET MEALS
	meals := getMealsInCategory(catgResult)
	var mealsNameArr []string
	mealsArr := makeMealsInCategoryArr(meals)
	for i := range mealsArr {
		mealsNameArr = append(mealsNameArr, mealsArr[i].MealName)
	}

	mealsPrompt := promptui.Select{
		Label: "Select a food category:",
		Items: mealsNameArr,
	}
	_, mealsResult, err := mealsPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	// GET RECIPE BY ID
	mealId := getMealId(mealsResult, mealsArr)
	recipe := getRecipeById(mealId)
	displayRecipeDetails(recipe)

}
