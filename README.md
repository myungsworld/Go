# Go

## Garbage Collector

Reference count = 0 이 되면 메모리를 지움  
   
## Slice

go 에서 동적배열은 capacity가 넘어가면 그 두배 크기의 배열을 새로 만들고 기존에 있는 배열을 삭제  
capacity가 남아있으면 값을 추가해도 변수에 할당한 주소는 동일  

값을 append해서 쓰고싶으면 복사를해서 서로 다른 메모리 공간을 확보해서 쓰는게 좋다.  