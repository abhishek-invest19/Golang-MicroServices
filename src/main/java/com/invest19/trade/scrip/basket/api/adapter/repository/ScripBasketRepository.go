package repository

//import "github.com/com/invest19/trade/scrip/basket/api/domain/ScripBasket"
import "github.com/com/invest19/trade/scrip/basket/api/domain/ScripBasket"

// Creating an interface
type ScripBasketRepository interface{

// Methods
	getScripBasket() ScripBasket
	createScripBasket() ScripBasket
}
