## Array
### Initialization
```go
1, nums := [3]int{1,2,3}
2, nums := [...]int{1,2,3}
```

### Traversal
```go
for i, v := range nums {
    //i is index, and v is element
}
```

## Map

### Initialization
```go
1, hashTable := map[int]int{}
2, hashTable := make(map[int]int)
hashTable[0] = 0
```

### Traversal
```go
for k, v := range hashTable {
	// k is key, v is value
}
for _, v := range hashTable {
	
}
for key := range hashTable {
	fmt.Println(hashTable[key])
}
```

### Judeg whether a key exists
```go
if _, ok := map[key]; ok {
    // exist
}

if _, ok := map[key]; !ok {
    // not exist
}

if value, ok := hashTable[key]; ok {
    return key, value
}
```
