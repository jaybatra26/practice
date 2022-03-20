package main 

import(
	"fmt"
)


type drink struct {
	water 		iunt16
	Milk  		iunt16
	sugar 		iunt16
	CoffeeBeans	iunt16
	tea beans	 iunt16
	cups 		  iunt16
	Money		 iunt16

}

var (
	drinks = make(map[string]Drink)

)

func GetavailableDrinks() map[string]Drink {
	return Drinks
}

func GetavailableDrinksname []string {
	drinksnamearr := make([]string,0,len(drinks))

	for key := range drinks {
		drinksNamearr = append(drinksnamearr,key)
	}
	return drinksnamearr
}