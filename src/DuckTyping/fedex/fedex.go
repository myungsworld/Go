package Fedex

import "fmt"

type FedexSender struct {

}

func (f *FedexSender) Send(parcel string) {
	fmt.Printf("Fedex에서 %s를 보냅니다.\n", parcel)
}
