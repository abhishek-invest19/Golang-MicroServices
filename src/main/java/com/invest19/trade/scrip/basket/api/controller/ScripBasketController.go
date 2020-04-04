package service

//import "github.com/com/invest19/trade/scrip/basket/api/domain/ScripBasket"
import "github.com/com/invest19/trade/scrip/basket/api/adapter/service/ScripBasketService"
import "github.com/com/invest19/trade/scrip/basket/api/domain/ScripBasket"

//adapter/service

// Creating an interface
type ScripBasketService interface{

// Methods
	getScripBasket() ScripBasket
	createScripBasket() ScripBasket
}
