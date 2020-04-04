package domain

//import "github.com/com/invest19/trade/scrip/basket/api/domain/ScripBasket"
//import "github.com/com/invest19/trade/scrip/basket/api/domain/ScripBasket"

var basketName string
	
func getbasketName () string {

    return  basketName
}

func setbasketName (string basketNameSource) {

    basketName = basketNameSource
}
