package coffee

import (
"fmt"
"github.com/spf13/viper"
)


type CoffeeDetails struct {
	Name string 
	Price float32 
}

type CoffeeList struct {
	List []CoffeeDetails 
}


var Coffees CoffeeList

// unmashal coffee data from a file to a list of objects
func GetCoffees() (*CoffeeList, error) {
	viper.AddConfigPath("./") // optionally look for config in the working directory : /Users/yuanyuanji/workspace/go 
	err := viper.ReadInConfig()            
	if err != nil { 
		fmt.Println("fatal error config file: %w", err)
		return nil, err
	}
	
	err = viper.Unmarshal(&Coffees)
	if err != nil {
		return nil, err
	}
	
	return &Coffees, nil
}