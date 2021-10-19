package postKorea

import "fmt"

type PostKoreaSender struct {

}

func (k *PostKoreaSender) Send(parcel string) {
	fmt.Printf("우체국에서 %s를 보냅니다.\n", parcel)
}