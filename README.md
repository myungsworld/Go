# Go

1. [Go Basic](#Garbage-Collector)
2. [Data Struct](#자료구조)
3. [Go Web Programming](#WEB)

## Garbage Collector

Reference count = 0 이 되면 메모리를 지움  
Go의 선업대입문은 모두 copy 되어 메모리에 저장된다.  

## String

문자열은 읽기전용이라 "가나다"를 "각나다"로 바꿀수 없음  
```go
s := "가나다"
s[2]++ // 에러
```

하지만 문자열을 아예 바이트 단위의 슬라이스로 변환 가능  
```go
func Example_modifyBytes(){
    b := []byte("가나다")
    b[2]++
    fmt.Println(string(b))
    // Output:
    // 각나다
}
```

- strings.Split(a,b) []string  
    - 문자열을 b를 기준으로 나눠 슬라이스로 만듬  
- strings.Index(a,b) int  
    - b가 a의 몇번째에 포함되어 있는지 확인, 포함이 되어 있지 않다면 -1을 리턴  

**어떤 문자들이 들어 있는지를 중요시 한다면 string**  
**실제 바이트 표현이 어떤지를 중시한다면 []byte**  

## Slice

go 에서 동적배열은 capacity가 넘어가면 그 두배 크기의 배열을 새로 만들고 기존에 있는 배열을 삭제  
capacity가 남아있으면 값을 추가해도 변수에 할당한 주소는 동일  

값을 append해서 쓰고싶으면 복사를해서 서로 다른 메모리 공간을 확보해서 쓰는게 좋다.  
Slice 한( ex) a = a[1:] 에 a[0] ) 메모리는 보이지 않을 뿐 사라지지 않는다.
Slice 의 주소와 원래의 주소는 다르고 Slice는 본 주소를 가리키고 있음


capacity가 8이고 len이 5인 slice에 새로운 slice 타입을 만들어서 append를 시키게 되면 
```go
    var a []int
    a = make([]int,5,8)
    var b []int
    b = append(a,1,2,3)
```
위에서 a 와 b slice가 가지는 주소값은 동일해서 b[0] 값을 9999로 바꾸게 되면 a와 b의 첫번째 값은 모두 9999가 된다.  
반면에  
```go
    var a []int
    a = make([]int,5,8)
    var b []int
    b = append(a,1,2,3,4)
```
위와 같이 a의 capacity를 넘기는 b slice를 추가하면 a 와 b 의 주소는 바뀌게 되어 b[0] 값을 9999로 바꾸면 b만 바뀌게 된다.  

- [slice의 삽입](https://github.com/myungsworld/Go/blob/main/src/DiscoveryGo/3./prac2/main.go)

## [Map](https://github.com/myungsworld/Go/blob/main/src/DiscoveryGo/3./%EC%97%B0%EC%8A%B5%EB%AC%B8%EC%A0%9C/5./main.go)
Key-Value 값으로 저장되는 배열  
```go
    m := make(map[int]int)
    m[0] = 0
    v, ok := m[0]
    v2, ok2 := m[1]
    fmt.Println(v, ok, v2, ok2)
    //출력 = 0 true 0 false

    delete(m, 0)
    v, ok := m[0]
    fmt.Println(v, ok, v2, ok2)
    //출력 = 0 false 0 false

    //map 순회
    for key, value := range m {
        fmt.Println(key, " : " , value)
    }
```
map의 value 값이 정해지지 않으면 int 0 , string = "" 로 초기화, 확인하는 방법은 파라미터를 2개로 두고 bool형으로 확인  

- Hash
    - 출력 값의 범위가 있다. (넣을수 있는 입력값은 한정되어 있지만 출력값은 동일할수도 있음) 
    - Hash를 만족할수 있는 함수 sign function , Modular (나머지 연산)
    - 공개키 암호화 , Checksum , 블록체인 
    - [RollingHash](https://src/dataStruct/map.go)
    - 다른 입력을 넣었을때 같은 값이 나올수 있는데 이걸 방지하기 위해 해쉬값에다가 또다른 배열을 추가 -> [Hash](https://src/datastruct/map.go)
    - HashMap (Hash 사용)
        - find , add , remove : O(1)
    - SortedMap (BST 사용)
        - find , add , remove : O(log2N)   
- Modular -> One-way function
    - X Mod N = M 일때 X의 후보군이 될수있는 값은 무수히 많음  
    - 이러한 구조때문에 암호화가 가능

- [구조체 내장](https://github.com/myungsworld/Go/tree/main/src/DiscoveryGo/5./%EA%B5%AC%EC%A1%B0%EC%B2%B4%EB%82%B4%EC%9E%A5)
    - 구조체를 내장하게 되면 내장된 구조체에 들어있는 필드들도 바르 접근이 가능하다.
- [직렬화 및 역직렬화]()

##  자료구조

- **Slice**  (cache에 좋음)
    - append -> O(N)  
    - remove -> O(N)  
    - Random Access -> O(1)  
        - stack   
            - append -> O(N)  
            - remove -> O(1)  
        - queue  
            - append -> O(N)  
            - remove -> O(1)  
- [**LinkedList**](https://github.com/myungsworld/Go/blob/main/src/dataStruct/linkedList.go)  (cache 적중률이 낮아짐 (거의 없음))
    - append -> O(1)  
    - remove -> O(1)    
    - Random Access -> O(N)  
        - stack 
            - append -> O(1)
            - remove -> O(1)
        - queue
            - append -> O(1)
            - remove -> O(1)
- **Tree**

    - [DFS(깊이우선검색)](https://github.com/myungsworld/Go/blob/main/src/dataStruct/tree.go)
        1. 재귀호출
        2. 스택  
    - [BFS(넓이우선검색)](https://github.com/myungsworld/Go/blob/main/src/dataStruct/tree.go)
        1. 큐  
    - [BST(이진검색트리)](https://github.com/myungsworld/Go/blob/main/src/dataStruct/binaryTree.go)  
        1. Ex) Parent 기준으로부터 left는 Parent 보다 작은 노드 right는 Parent 보다 큰 노드  
        2. 특정 노드를 찾을 때 좋음(검색)  O(log2N)
    - AVL(최소신장트리)
        1. 회전시켜서 트리의 신장을 최소로 줄임  
 
- [**Heap**](https://github.com/myungsworld/Go/blob/main/src/dataStruct/heap.go) :정렬을 만들 수 있음 (힙정렬)
    - 속도 : O(Nlog2N)
        - Push : O(log2N) 마지막 노드에 추가해 위의 노드와 비교해서 올라감  
        - Pop : O(log2N) 맨 끝 노드를 맨 위로 올린다음 자식노드와 비교해서 큰 노드와 교체  
    - 최대 힙(우선순위 큐) : 부모노드가 자식노드보다 크거나 같아야 함
    - 최소 힙 : 부모노드가 자식노드보다 작거나 같아야 함  



- ## [Thread](https://github.com/myungsworld/Go/blob/main/src/Intermediate/thread/main.go)

    - Go 에서는 OS의 CPU의 개수만큼 쓰레드를 만들고 각 OS의 쓰레드를 짤라서 여러개의 Go Theard 를 할당한다.  
    프로그래머 입장에서 Context switching을 상관쓰지 않아도 됨 (다른 언어보다)  
    - [병렬 프로그래밍](https://lynlab.co.kr/blog/82)
    - Context
        - 동시에 실행되는 고루틴 ex) 요청 값, 취소 신호, 고루틴의 작업 마감시간 등의 작업을 쉽게 하기 위해 context 패키지를 제공
        - Cancel context : 고루틴의 종료를 관리해 줄수 있는 context 이를 통해 불필요한 작업을 취소 가능
고루틴은 백그라운드에서 비동기적으로 실행되며 프로그램은 고루틴이 끝날 때까지 기다려주지 않는다. -> sync.WaitGroup 활용


**레이스 컨디션**
전역 변수가 있을 때 여러 프로세스가 동시에 접근하려 하면서 서로 경쟁하는 상태
여러 쓰레드가 같은 메모리 영역을 건드릴때는 Lock을 사용 ([Mutex](https://github.com/myungsworld/Go/blob/main/src/Intermediate/account/main.go))  
Golang에서는 이걸 쉽게 하기위해 channel을 제공함 => 동기화 제어


## [Channel (Queue)](https://github.com/myungsworld/Go/blob/main/src/Intermediate/factory/main.go)
Thread Safe , fixed size queue  
```go
    var a chan int
    a = make(chan int,1)
    //push to chan a
    a <- 10
    //Pop from chan a to b
    b := <- a 
```
[select](https://github.com/myungsworld/Go/blob/main/src/Intermediate/select/main.go)

## OOP  
- [절차적 프로그래밍](https://github.com/myungsworld/Go/blob/main/src/Intermediate/sandwitch/main.go)  
    - 순서와 절차
- [OOP](https://github.com/myungsworld/Go/blob/main/src/Intermediate/method/main.go)
    - 각 Objects 의 기능
    - 각 Objects 간의 관계
    - 메서드는 object에 속해있는 function
- [Interface](https://github.com/myungsworld/Go/blob/main/src/Intermediate/sandwitch2/main.go)
    - 객체간 상호 관계 정의
    - object = 상태 + 기능 , interface는 obejct의 기능을 따로 정함, decoupling 종속성 제거 
- OOD : object oriented design
    - SOLID 객체 중심 설계를 위해 지향해야할 목표
        - [S : Single Reponsibility principle](https://github.com/myungsworld/Go/blob/main/src/Intermediate/ood/srp.go)
            - 단일 책임 원칙 : 하나의 객체는 하나의 책임을 가져야 한다
        - [O : Open closed principle](https://github.com/myungsworld/Go/blob/main/src/Intermediate/ood/srp.go)
            - 확장에는 열려있고 변경에는 닫혀있다.
        - L : Liskop subtitution principle
            - 리스코프 치환 이론 : Base Type의 기존 함수(동작)를 Extended Type에서도 동일하게 적용이 되어야한다.
        - [I : Interface segregation principle](https://github.com/myungsworld/Go/blob/main/src/Intermediate/ood/isp.go)
            - 인터페이스 분리원칙 : 여러개의 관계를 모아놓은 인터페이스보다 관계 하나씩 정의하는게 더 좋다. 
        - [D : Dependency inversion principle](https://github.com/myungsworld/Go/blob/main/src/Intermediate/ood/dip.go)
            - 관계는 인터페이스에 의존해야 한다. 객체가 아니라
- [ORM]()
    - 객체와 데이터베이스를 연결(맵핑)
    - go get github.com/jinzhu/gorm
- ## Beyond OOP
    - Micro Service
    - Serverless
    - ECS(Entity Component System)
        - 직접도가 올라감
    - MVC(Model View Controller)
        - Data와 View 와 기능을 따로 만들어서 View에서 쓰는 것  

## 계속 공부해야 하는 것들

1. **컴퓨터 원리** -> 하드웨어 -> CPU 설계
2. **프로그래밍언어 & 문법**
3. **자료구조 & 알고리즘**
4. **Thread & 고급기능**
5. **OOP** -> 설계 (아키텍쳐)

# WEB

- [goconvey => test 환경 구축](https://github.com/myungsworld/Go/tree/main/src/github.com/WEB)  
- [Static Web Server(File upload)](https://github.com/myungsworld/Go/tree/main/src/github.com/WEB2)
- [RestfulAPI](https://github.com/myungsworld/Go/tree/main/src/github.com/WEB3)
    - 자원에 대한 CRUD 조작을 URL에다가 표시
    - gorilla/mux => 자동으로 파싱시켜주는 외부 패키지, CRUD 조작을 쉽게 도와줌
    
    ```go
    //Restful 형식의 users/{id:[0-9]+} handling 
     vars := mux.Vars(r) , id := vars("id")
     id := vars("id")
    ```
- [Decorator + WEB4](https://github.com/myungsworld/Go/tree/main/src/github.com/WEB4)
- [Template](https://github.com/myungsworld/Go/tree/main/src/github.com/templateW)
    - Go에는 강력한(?) 내부 template package가 있음
- [Render,Pat,Negroni 외부패키지](ttps://github.com/myungsworld/Go/tree/main/src/github.com/packageW)
    - pat : mux와 비슷한 일을 하는데 좀더 간단한 코드로 구현 가능
    - render : json parsing , html template 간단한 코드로 구현
    - negroni : HTTP Middleware
- [Chat service with eventSource](https://github.com/myungsworld/Go/tree/main/src/github.com/chatW)
    - server-sent events에 대한 웹 콘텐츠 인터페이스(서버에서만 보낼수 있음)
- [Oauth](https://github.com/myungsworld/Go/tree/main/src/github.com/Oauth)
    - go get golang.org/x/oauth2
    - go get cloud.google.com/go  
- [TodoList](https://github.com/myungsworld/Go/tree/main/src/github.com/todolistW)
    - [Refactoring Handler](https://github.com/myungsworld/Go/tree/main/src/github.com/todolistW/model)
    - [SQLiteDB Handler](https://github.com/myungsworld/Go/tree/main/src/github.com/todolistW/model/sqliteHandler.go)
    - [Oauth, Session](https://github.com/myungsworld/Go/tree/main/src/github.com/todolistW/myapp/signin.go)
- [Web Socket](https://github.com/myungsworld/Go/tree/main/src/github.com/socket)    
 
- [배포](https://www.youtube.com/watch?v=Samw30ZlkuU&list=PLy-g2fnSzUTDALoERcKDniql16SAaQYHF&index=24)
    - 도메인 구매 , DNS 등록
    - 개인 IP는 ISP(Internet Service Provider) 가 제공한 동적 IP 할당함으로 Private 이다.  따라서 외부에서 접속이 불가
    - Public IP를 받으려면 호스팅을 받아야 한다. 그중 클라우드 서비스를 많이 사용 IaaS(Infrasturct as a Service)

- ## Design Pattern
    - [Decorator Pattern](https://github.com/myungsworld/Go/blob/main/src/designPattern/decoratePattern/main.go)
        - 객체에 새로운 행동을 동적으로 첨가해 객체의 기능을 유연하게 확장 시킴
    - [Singleton Pattern](https://github.com/myungsworld/Go/blob/main/src/designPattern/singletonPattern/main.go)
        - 프로그램 내부에서 특정 클래스에 대해 하나의 인스턴스만 존재하도록 강제하는 디자인 패턴
    - [Factory method Pattern](https://github.com/myungsworld/Go/blob/main/src/designPattern/factoryPattern/main.go)
        - 메서드의 구현을 서브 클래스에 위임하는 방식
    - [Concurrency Pattern](https://github.com/myungsworld/Go/blob/main/src/designPattern/concurrencyPattern)
        - 
- ## TDD(Test Driven Development(테스트 주도 개발))
    - 기존의 설계 -> 구현 -> 테스트 방식이 아닌 테스트부터 하면서 코딩하는 방법
    - WEB 쪽에있는 모든 코드가 TDD방식 
    - 장점: 유지보수가 뛰어나다, test가 촘촘하게 이루어져 있을수록 에러잡기가 수월하다.

- ## Data Store
    - [In-Memory Storage](https://github.com/myungsworld/Go/blob/main/src/dataStore/inmemory/main.go)
        - 앱이 실행되는 동안에 사용 , 디스크를 사용하지 않아 동작이 빠르고 즉각적이지만 영구적으로 사용이 불가능하다
    - [file Storage](https://github.com/myungsworld/Go/blob/main/src/dataStore/fileStorage)
        - CSV : 데이터의 목적에 맞게 디코딩하여 사용
        - god 패키지 : 인코딩과 디코딩
    - [Database](https://github.com/myungsworld/Go/blob/main/src/dataStore/database)
        - SQL , NoSQL

- ## AWS
    - [Lambda](https://github.com/myungsworld/Go/blob/main/src/aws/lambda-functions/main.go)
        - go get github.com/aws/aws-lambda-go/lambda
        - GOOS=linux go build -o main : 리눅스 바이너리 파일로 빌드
        - 압축한다음 AWS lambda function에 넣고 Test
- ## Framework
    - [echo](https://github.com/myungsworld/Go/blob/main/src/echoframework/main.go)
    - [StockAPI](https://github.com/myungsworld/Go/blob/main/src/stock/controllers/price.go)
    io.ReadWriter의 content를 보존하는 법
```go
        var bodyBytes []byte
	    if c.Request().Body != nil {
		    bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	    }   
	    c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
```
c.Request().Body는 io.ReadCloser type 이라서 bodyBytes는 빈공간이 남게된다.  
따라서 빈공간에 다시 새로운 버퍼를 넣어서 읽을수 있게 만들어줌  

        
        
- ## [JWT](https://github.com/myungsworld/Go/blob/main/src/jwt/login.go)
    - JSON 객체로 당사자간의 정보를 안정하게 전송하기 위한 방법
        - 상태 비저장 => 불투명 토큰과 달리 데이터베이스에 저장할 필요가 없음
        - 일정기간이 지나면 무효로 설정가능 => XSS 공격 방지 
    - 헤더.페이로드.서명 의 형태로 토큰에 담김
        - 헤더
            - 사용된 암호화 알고리즘과 키값을 담고 있음
        - 페이로드
            - 토큰 클라이언트의 정보, 생성일시 등 정보를 저장
        - 서명
            - 헤더와 페이로드를 합친 문자열을 서명한 값 (Base64 URL-Safe로 인코딩)
