## 总结

#### 递归

代码模版

```python
class Solution:
    def recursion(level,param1,param2...):
        #递归终止条件
        if level>MAX_LEVEL:
        	process_result
        	return
        #处理当前层逻辑
        process(level,data...)
        #下探到下一层
        self.recursion(level+1,p1,p2...)
        #清理当前层
```

注意事项

1. 不要人肉进行递归
2. 找到最近最简方法，拆解成可重复解决的子问题
3. 数学归纳思想



#### 分治

一种特殊的递归，本质上个是找重复性以及分解问题和最后组合每个子问题的结果

代码模版

```python
def divide_conquer(problem, param1, param2, ...): 
  # recursion terminator 
  if problem is None: 
	print_result 
	return 

  # prepare data 
  data = prepare_data(problem) 
  subproblems = split_problem(problem, data) 

  # conquer subproblems 
  subresult1 = self.divide_conquer(subproblems[0], p1, ...) 
  subresult2 = self.divide_conquer(subproblems[1], p1, ...) 
  subresult3 = self.divide_conquer(subproblems[2], p1, ...) 
  …

  # process and generate the final result 
  result = process_result(subresult1, subresult2, subresult3, …)
	
  # revert the current level states
```

模版与递归模版的区别在于分解子问题后组合在一起，递归分治和回溯本质是一样的



#### 回溯

提前判断，如果发现不行，可以节省执行时间。
