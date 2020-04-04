package service

//import "github.com/com/invest19/trade/scrip/basket/api/domain/ScripBasket"
import "github.com/com/invest19/trade/scrip/basket/api/domain/ScripBasket"

// Creating an interface
type ScripBasketService interface{

// Methods
	getScripBasket() ScripBasket
	createScripBasket() ScripBasket
}
