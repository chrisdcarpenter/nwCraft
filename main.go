package main

import (
	"encoding/json"
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/chris-carpenter/nwCraft/recipes"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
	"time"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	var parser = argparse.NewParser("nwCraft", "Provides shopping list for desired craft")
	item := parser.String("i", "item", &argparse.Options{Required: true, Help: "Item to craft"})
	file := parser.String("f", "file", &argparse.Options{Required: false, Help: "File of craftData to load", Default: "sampleData.json"})
	debugLevel := parser.Selector("d", "debug-level", []string{"INFO", "DEBUG"}, &argparse.Options{Required: false, Help: "Logging debug level"})
	pretty := parser.Flag("p", "pretty", &argparse.Options{Required: false, Help: "Pretty output"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(parser.Usage(err))
		return
	}

	switch *debugLevel {
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}

	if *pretty {
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		output.FormatLevel = func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
		}
		output.FormatMessage = func(i interface{}) string {
			return fmt.Sprintf("%s", i)
		}
		output.FormatFieldName = func(i interface{}) string {
			return fmt.Sprintf("| %s:", i)
		}
		output.FormatFieldValue = func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("%s", i))
		}

		log.Logger = zerolog.New(output).With().Timestamp().Logger()
	}

	log.Info().Str("item", *item).Str("file", *file).Msg("Crafting params")

	craftData, err := recipes.NewRecipes(*file)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Error loading recipes file.")
	}

	//Actual output
	baselineRecipes := recipes.Recipe{Name: *item, Ingredients: craftData.GetIngredients(*item)}

	result, err := jsonMarshall(baselineRecipes, *pretty)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Error getting ingredients")
	}
	fmt.Println(result)
}

func jsonMarshall(data interface{}, pretty bool) (string, error) {
	var value []byte
	if pretty {
		val, err := json.MarshalIndent(data, "", "    ")
		if err != nil {
			return "", err
		}
		value = val
	} else {
		val, err := json.Marshal(data)
		if err != nil {
			return "", err
		}
		value = val
	}
	return string(value), nil
}
