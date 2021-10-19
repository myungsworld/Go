package main

import (
	"Go/src/DuckTyping/postKorea"
)


type Sender interface{
	Send(parcel string)
}

//func SendBook(name string, sender *Fedex.FedexSender) {
//	sender.Send(name)
//}

func SendBook(name string , sender Sender) {
	sender.Send(name)
}



func main() {

	var sender Sender = &postKorea.PostKoreaSender{}

	SendBook("어린왕자", sender)
	SendBook("지못미", sender)
}
