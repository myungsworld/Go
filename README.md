# Go

1. [Go Basic](#Garbage-Collector)
2. [Go Web Programming](#WEB)

## Garbage Collector

Reference count = 0 이 되면 메모리를 지움  
Go의 선업대입문은 모두 copy 되어 메모리에 저장된다.  
    
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


## [Map](https://github.com/myungsworld/Go/blob/main/src/dataStruct/map.go)
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



## [Thread](https://github.com/myungsworld/Go/blob/main/src/Intermediate/thread/main.go)

Go 에서는 OS의 CPU의 개수만큼 쓰레드를 만들고 각 OS의 쓰레드를 짤라서 여러개의 Go Theard 를 할당한다.  
프로그래머 입장에서 Context switching을 상관쓰지 않아도 됨 (다른 언어보다)  

여러 쓰레드가 같은 메모리 영역을 건드릴때는 Lock을 사용 ([Mutex](https://github.com/myungsworld/Go/blob/main/src/Intermediate/account/main.go))  
Golang에서는 이걸 쉽게 하기위해 channel을 제공함  

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
            - 리스코프 치환 이론 : Base Type의 기존 함수(동작)를 Exprended Type에서도 동일하게 적용이 되어야한다.
        - [I : Interface segregation principle](https://github.com/myungsworld/Go/blob/main/src/Intermeditae/ood/isp.go)
            - 인터페이스 분리원칙 : 여러개의 관계를 모아놓은 인터페이스보다 관계 하나씩 정의하는게 더 좋다. 
        - [D : Dependency inversion principle](https://github.com/myungsworld/Go/blob/main/src/Intermediate/ood/dip.go)
            - 관계는 인터페이스에 의존해야 한다. 객체가 아니라

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

- [goconvey => test 환경 구축](https://go/src/github.com/WEB/myapp/app_test.go)  
- [Static Web Server(File upload)](https://go/src/github.com/com/WEB2)
- [RestfulAPI](https://go/src/github.com/WEB3)
    - 자원에 대한 CRUD 조작을 URL에다가 표시
    - gorilla/mux => 자동으로 파싱시켜주는 외부 패키지, CRUD 조작을 쉽게 도와줌
    
    ```go
    //Restful 형식의 users/{id:[0-9]+} handling 
     vars := mux.Vars(r) , id := vars("id")
     id := vars("id")
    ```
- [Decorator + WEB4](https://go/src/github.com/WEB4)
- [Template](https://go/src/github.com/templateW)
    - Go에는 강력한(?) 내부 template package가 있음
- [Render,Pat,Negroni 외부패키지](https://go/src/github.com/packageW)
    - pat : mux와 비슷한 일을 하는데 좀더 간단한 코드로 구현 가능
    - render : json parsing , html template 간단한 코드로 구현
    - negroni : HTTP Middleware
- [Chat service with eventSource](https://go/src/github.com/chatW)
    - server-sent events에 대한 웹 콘텐츠 인터페이스(서버에서만 보낼수 있음)
- [Oauth](https://go/src/github.com/Oauth)
    - go get golang.org/x/oauth2
    - go get cloud.google.com/go  
- [TodoList](https://go/src/github.com/todolistW)

- 배포
    - 도메인 구매 , DNS 등록
    - 개인 IP는 ISP(Internet Service Provider) 가 제공한 동적 IP 할당함으로 Private 이다.  따라서 외부에서 접속이 불가
    - Public IP를 받으려면 호스팅을 받아야 한다. 그중 클라우드 서비스를 많이 사용 IaaS(Infrasturct as a Service)

- ## Design Pattern
    - [Decorator Pattern](https://go/src/decoratePattern/main.go)
        - 객체에 새로운 행동을 동적으로 첨가해 객체의 기능을 유연하게 확장 시킴

- ## TDD(Test Driven Development(테스트 주도 개발))
    - 기존의 설계 -> 구현 -> 테스트 방식이 아닌 테스트부터 하면서 코딩하는 방법
    - WEB 쪽에있는 모든 코드가 TDD방식 
    - 장점: 유지보수가 뛰어나다, test가 촘촘하게 이루어져 있을수록 에러잡기가 수월하다.
