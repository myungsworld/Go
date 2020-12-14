# Go

Go의 선업대입문은 모두 copy 되어 메모리에 저장된다.  

## Garbage Collector

Reference count = 0 이 되면 메모리를 지움  
    
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



##  시간복잡도

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
- **ListLinked**  (cache 적중률이 낮아짐 (거의 없음))
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

    - [DFS(깊이우선검색)](https://src/dataStruct/tree.go)
        1. 재귀호출
        2. 스택  
    - [BFS(넓이우선검색)](https://src/dataStruct/tree.go)
        1. 큐  
    - [BST(이진검색트리)](https://src/datasStruct/binaryTree.go)  
        1. Parent 기준으로부터 left는 Parent 보다 작은 노드 right는 Parent 보다 큰 노드  
        2. 특정 노드를 찾을 때 좋음(검색)  O(log2N)