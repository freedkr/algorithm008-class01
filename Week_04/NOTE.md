# 笔记

1. ##### 深度优先搜索

模版代码：递归

```python
visited = set()
def DFS(node, visted):
    # terminator
    if node in visited:
        # already visited
        return

    visited.add(node)

    # process current node here
    ...
    for next_node in node_children():
        if not next_node in visited:
            DFS(next_node, visited)
```

非递归

```python
def DFS(self, tree): 

    if tree.root is None: 
     return [] 

    visited, stack = [], [tree.root]

    while stack: 
     node = stack.pop() 
     visited.add(node)

     process (node) 
     nodes = generate_related_nodes(node) 
     stack.push(nodes) 

    # other processing work 
		...
```



2. ##### 广度优先搜索

```python
def BFS(graph, start, end):
    queue = []
    queue.append([start])
    visited.add(start)

    while queue:
        node = queue.pop()
        visited.add(node)

        process(node)
        nodes = generate_related_nodes(node)
        queue.push(nodes)
    
    # other processing work
    ...
```



3. ##### 贪心算法

   定义：**贪心算法**是一种在每一步选择中都采取在当前状态下最好或最优的选择，从而希望结果是全局最好或者最优的算法。 

   1. 贪心算法会对每一个子问题的解决方案都做出选择，不能回退；
   2. 动态规划会保存以前的运算结果，并根据以前的结果对当前进行选择， 有回退功能。
   3. 能用贪心算法解决的问题，也可以用动态规划解决

   适用场景：具备贪心选择性质的问题。即所求问题的整体最优解可以通过一系列局部最优的选择，即贪心选择来达到。

4. ##### 二分查找

   二分查找也称折半查找（Binary Search），它是一种效率较高的查找方法。但是，折半查找要求线性表必须采用顺序存储结构，而且表中元素按关键字有序排列。

   寻找某数

   ```c
   int binarySearch(int[] nums, int target) {
       int left = 0; 
       int right = nums.length - 1; // 注意
   
       while(left <= right) {
           int mid = left + (right - left) / 2;
           if(nums[mid] == target)
               return mid; 
           else if (nums[mid] < target)
               left = mid + 1; // 注意
           else if (nums[mid] > target)
               right = mid - 1; // 注意
       }
       return -1;
   }
   ```

   