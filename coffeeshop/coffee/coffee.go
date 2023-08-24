package coffee

import (
	"fmt"

	"github.com/spf13/viper"
)

type CoffeeDetails struct {
	Name  string
	Price float32
}

type CoffeeList struct {
	CoffeeList []CoffeeDetails
}

var Coffees CoffeeList

func GetCoffees() (*CoffeeList, error) {
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("error: %w", err)
		return nil, err
	}

	err = viper.Unmarshal(&Coffees)
	if err != nil {
		return nil, err
	}

	return &Coffees, nil
}

func IsCoffeeAvailable(coffeetype string) bool {
	for _, element := range Coffees.CoffeeList {
		if element.Name == coffeetype {
			result := fmt.Sprintf("%s for $%v", element.Name, element.Price)
			fmt.Println(result)
			return true
		}
	}
	return false
}
