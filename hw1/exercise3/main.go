package main

import (
	"fmt"
)

type Shopping struct {
	Name string
	Price int
	Quantity int
}

func main() {
	shopList := []Shopping{
		Shopping{"kruh", 8, 2},
		Shopping{"mlijeko", 15, 1},
	}
	solution1, err := mostExpensive(shopList)
	fmt.Println(solution1) // mlijeko
	fmt.Println(err) // nil
	solution2 := totalCost(shopList)
	fmt.Println(solution2) // 31
}

func mostExpensive(shopList []Shopping) (item []Shopping, err error) {
	// tijelo funkcije
	// pripaziti: kad se može dogoditi pogreška?
	var errMsg error

	if len(shopList) == 0{
		errMsg = fmt.Errorf("no data")
		return nil, errMsg
	}
	expensiveItems := make([]Shopping, 1, 5)
	expensiveItems = append(expensiveItems, shopList[0])

	for _, shopItem := range shopList {
		if shopItem.Price > expensiveItems[0].Price {
			expensiveItems = make([]Shopping, 1, 5)
			expensiveItems = append(expensiveItems, shopItem)
		} else if shopItem.Price == expensiveItems[0].Price{
			expensiveItems = append(expensiveItems, shopItem)
		}
	}
	return expensiveItems, errMsg
}

func totalCost(shopList []Shopping) (total int) {
	sum := 0

	for _, shopItem := range shopList{
		sum += shopItem.Price * shopItem.Quantity
	}
	return sum
}