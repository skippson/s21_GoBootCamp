package main

import (
	"flag"
	"fmt"
	"log"
	models "secondDay/models"
	reader "secondDay/reader"
	"strings"
)

type Pair struct {
	count, unit string
}

func compare(first, second models.Recipes) {
	oldRecipeNames := make(map[string]bool)
	newRecipeNames := make(map[string]bool)

	for _, recipe := range first.Cake {
		oldRecipeNames[recipe.Name] = true

	}
	for _, recipe := range second.Cake {
		newRecipeNames[recipe.Name] = true
	}
	for _, new_recipe := range second.Cake {
		if _, inMap := oldRecipeNames[new_recipe.Name]; !inMap {
			fmt.Printf("ADDED cake \"%s\"\n", new_recipe.Name)
		}

	}
	for _, old_recipe := range first.Cake {
		if _, inMap := newRecipeNames[old_recipe.Name]; !inMap {
			fmt.Printf("REMOVED cake \"%s\"\n", old_recipe.Name)
		}

	}
	for _, old_recipe := range first.Cake {
		for _, new_recipe := range second.Cake {
			if old_recipe.Name == new_recipe.Name {
				if old_recipe.Time != new_recipe.Time {
					fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", new_recipe.Name, new_recipe.Time, old_recipe.Time)
				}
				oldIngName := make(map[string]Pair) //для названия ингридиентов
				newIngName := make(map[string]Pair)
				for _, ing := range old_recipe.Ingridients {
					oldIngName[ing.Name] = Pair{ing.Count, ing.Unit}
				}
				for _, ing := range new_recipe.Ingridients {
					newIngName[ing.Name] = Pair{ing.Count, ing.Unit}
				}
				for _, new_ing := range new_recipe.Ingridients {
					if _, inMap := oldIngName[new_ing.Name]; !inMap {
						fmt.Printf("ADDED ingredient \"%s\" for cake  \"%s\"\n", new_ing.Name, new_recipe.Name)
					} else {
						if new_ing.Unit != oldIngName[new_ing.Name].unit && new_ing.Unit != "" && oldIngName[new_ing.Name].unit != "" {
							fmt.Printf("CHANGED unit for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n", new_ing.Name, new_recipe.Name, new_ing.Unit, oldIngName[new_ing.Name].unit)

						} else if new_ing.Unit != oldIngName[new_ing.Name].unit && new_ing.Unit == "" && oldIngName[new_ing.Name].unit != "" {
							fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", oldIngName[new_ing.Name].unit, new_ing.Name, new_recipe.Name)

						} else if new_ing.Unit != oldIngName[new_ing.Name].unit && new_ing.Unit != "" && oldIngName[new_ing.Name].unit == "" {
							fmt.Printf("ADDED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", new_ing.Unit, new_ing.Name, new_recipe.Name)
						}
					}
				}

				for _, old_ing := range old_recipe.Ingridients {
					if _, inMap := newIngName[old_ing.Name]; !inMap {
						fmt.Printf("REMOVED ingredient \"%s\" for cake  \"%s\"\n", old_ing.Name, old_recipe.Name)
					} else {
						if old_ing.Count != newIngName[old_ing.Name].count && old_ing.Count != "" && newIngName[old_ing.Name].count != "" {
							fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n", old_ing.Name, new_recipe.Name, newIngName[old_ing.Name].count, old_ing.Count)

						} else if old_ing.Count != newIngName[old_ing.Name].count && old_ing.Count == "" && newIngName[old_ing.Name].count != "" {
							fmt.Printf("ADDED unit count \"%s\" for ingredient \"%s\" for cake  \"%s\"\n", newIngName[old_ing.Name].count, old_ing.Name, new_recipe.Name)

						} else if old_ing.Count != newIngName[old_ing.Name].count && old_ing.Count != "" && newIngName[old_ing.Name].count == "" {
							fmt.Printf("REMOVED unit count \"%s\" for ingredient \"%s\" for cake  \"%s\"\n", newIngName[old_ing.Name].count, old_ing.Name, new_recipe.Name)
						}
					}
				}

			}
		}
	}
}

func main() {
	oldFilename := flag.String("old", "", "input file name")
	newFilename := flag.String("new", "", "input file name")
	flag.Parse()
	if *oldFilename == "" || *newFilename == "" {
		log.Fatal("no input file specified")
	}
	fileOld := strings.HasSuffix(*oldFilename, ".xml")
	fileNew := strings.HasSuffix(*newFilename, ".json")
	db := reader.DB{}
	if fileOld && fileNew {
		_, old, err := db.DBReader(*oldFilename, "xml")
		if err != nil {
			log.Fatal(err)
		}
		_, new, err := db.DBReader(*newFilename, "json")
		if err != nil {
			log.Fatal(err)
		}
		compare(*old, *new)
	} else if !fileOld && !fileNew {
		_, old, err := db.DBReader(*oldFilename, "json")
		if err != nil {
			log.Fatal(err)
		}
		_, new, err := db.DBReader(*newFilename, "xml")
		if err != nil {
			log.Fatal(err)
		}
		compare(*old, *new)
	} else {
		fmt.Println("invalid filename")
	}
}
