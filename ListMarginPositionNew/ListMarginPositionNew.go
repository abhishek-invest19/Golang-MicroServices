package main
import (
    "bufio"
    "fmt"
    "net/http"
)
func main() {

resp, err := http.Get("http://membersim.bseindia.com/stocks/ListMarginPositionNew?USBackOfficeId=Abhishek-t9072&USLOGINID=Abhishek-t9072&SessionKey=1380713849-02287238289293&Thick+Client=Y")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("Response status:", resp.Status)

    scanner := bufio.NewScanner(resp.Body)
    for i := 0; scanner.Scan() && i < 5; i++ {
        fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
