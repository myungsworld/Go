package singleton

import (
	"fmt"
	"sync"
)

// 구조체의 이름을 소문자로 해서 패키자 밖으로 export 하지 않음
type singleton struct {
}

// 마찬가지로 소문자로 export 하지 않음
var instance *singleton
var singleInstance *singleton
var m = &sync.Mutex{}
var once sync.Once

func GetInstance() *singleton {
	if instance == nil {
		m.Lock()
		defer m.Unlock()
		if instance == nil {
			fmt.Println("singleton 객체 생성중")
			instance = &singleton{}
		} else {
			fmt.Println("singleton 객체가 이미 있음-1")
		}
	} else {
		fmt.Println("singleton 객체가 이미 있음-2")
	}
	return instance
}
