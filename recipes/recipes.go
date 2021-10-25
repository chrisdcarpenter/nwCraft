package recipes

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

type Recipe struct {
	Name        string         `json:"name"`
	Ingredients map[string]int `json:"ingredients"`
}

type CraftData struct {
	GameName string   `json:"game_name"`
	Recipes  []Recipe `json:"recipes"`
}

func NewRecipes(filePath string) (*CraftData, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	log.Info().Str("file", filePath).Msg("Successfully opened")
	defer jsonFile.Close()

	var byteValue, _ = ioutil.ReadAll(jsonFile)

	var craftData CraftData
	json.Unmarshal(byteValue, &craftData)

	log.Debug().Int("length", len(craftData.Recipes)).Msg("Number of recipes")
	return &craftData, nil
}

func (c *CraftData) GetIngredients(item string) map[string]int {
	ingredients := make(map[string]int)

	itemIndex := c.FindRecipes(item)
	if itemIndex == -1 {
		log.Debug().Str("item", item).Msg("Couldn't find recipe")
		return map[string]int{item: 1}
	}

	for k, v := range c.Recipes[itemIndex].Ingredients {
		log.Debug().Str("item", k).Msg("Looking up ingredients")
		ingredients = combineMaps(ingredients, multiplyIngredients(v, c.GetIngredients(k)))
	}
	return ingredients
}

func (c *CraftData) FindRecipes(item string) int {
	for i := range c.Recipes {
		if normalizeItemName(item) == normalizeItemName(c.Recipes[i].Name) {
			log.Debug().Str("item", item).Int("index", i).Msg("Recipe found")
			return i
		}
	}
	log.Warn().Str("item", item).Msg("Could not find recipe")
	return -1
}

func multiplyIngredients(count int, ingredients map[string]int) map[string]int {
	for i := range ingredients {
		ingredients[i] *= count
	}
	return ingredients
}

func combineMaps(a map[string]int, b map[string]int) map[string]int {
	for k, v := range b {
		if _, ok := a[k]; ok {
			a[k] += v
		} else {
			a[k] = v
		}
		a[k] = v
	}
	return a
}

func normalizeItemName(item string) string {
	return strings.ToLower(spaceStringsBuilder(item))
}

func spaceStringsBuilder(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}