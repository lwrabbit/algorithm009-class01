学习笔记

## 递归模版
```go
func recursion(problem,...params){
	//recursion terminator 
	if last_problem {
		print_result
		return
	}
	//process cur problem
	process...
	//drill down
	recursion(next_problem)
	//revert state
	revert...
}
```

## 分治模版
```go
func divide_conquer(problem,...params){
	//recursion terminator
	if problem is nil{
		print_result
		return
	}
	//prepare data
	data := prepare_data(problem)
	subproblems := split_problem(problem, data)
    // conquer subproblems 
    var res 
    for subproblem := range subproblems{
    	subres := divide_conquer(subproblem,...params)
    	// process and generate the final result
    	res = append(res,subres)
    }
    // revert the current level states
}
```
## 重点掌握
- 熟练套用递归模版
- 注意全局变量/传参 的状态恢复