# Kahn 算法介绍

Kahn 算法是一种用于解决 **拓扑排序** 问题的经典算法。拓扑排序是对一个有向无环图（DAG）进行排序，使得对于图中的每一条有向边 `(u, v)`，顶点 `u` 排在顶点 `v` 之前。Kahn 算法是一个广度优先搜索（BFS）算法，能够高效地求解拓扑排序。

## Kahn 算法的步骤

Kahn 算法的基本步骤如下：

1. **初始化入度表**：首先，计算图中每个节点的入度。入度是指某个节点的所有指向该节点的边的数量。

2. **初始化队列**：将所有入度为 0 的节点加入一个队列中。入度为 0 的节点可以没有任何依赖，可以首先处理。

3. **遍历节点**：
    - 从队列中弹出一个节点，并将其加入拓扑排序结果。
    - 遍历该节点的所有邻接节点，并将这些邻接节点的入度减 1。
    - 如果某个邻接节点的入度变为 0，将其加入队列中，等待进一步处理。

4. **检测循环**：如果拓扑排序中包含所有节点，则排序成功。如果队列为空但仍有未处理的节点，则说明图中有环，无法进行拓扑排序。

## Kahn 算法的时间复杂度

Kahn 算法的时间复杂度为 `O(V + E)`，其中 `V` 是图中节点的数量，`E` 是图中的边的数量。这个复杂度是因为每个节点和每条边都被访问一次。

## Kahn 算法的空间复杂度

空间复杂度为 `O(V + E)`，因为需要存储图的邻接表、入度表以及队列等。

## 示例代码（Python 实现）

```python
from collections import deque, defaultdict

def kahn_algorithm(graph):
    # 计算入度
    in_degree = defaultdict(int)
    for node in graph:
        for neighbor in graph[node]:
            in_degree[neighbor] += 1
    
    # 初始化队列，加入所有入度为 0 的节点
    queue = deque([node for node in graph if in_degree[node] == 0])
    
    topo_order = []
    
    while queue:
        node = queue.popleft()
        topo_order.append(node)
        
        # 遍历邻接节点，减少它们的入度
        for neighbor in graph[node]:
            in_degree[neighbor] -= 1
            if in_degree[neighbor] == 0:
                queue.append(neighbor)
    
    # 如果拓扑排序包含所有节点，返回结果；否则说明有环
    if len(topo_order) == len(graph):
        return topo_order
    else:
        raise ValueError("The graph has a cycle and cannot be topologically sorted")

```
# 示例图
```mermaid
graph = {
    'A': ['C', 'D'],
    'B': ['D'],
    'C': ['E'],
    'D': ['E'],
    'E': []
}

try:
    print("Topological Sort:", kahn_algorithm(graph))
except ValueError as e:
    print(e)
```

# 例子
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。

在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。

- 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。



**示例 1：**

输入：numCourses = 2, prerequisites = [[1,0]]
`输出：true`
解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。

**示例 2：**

输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
`输出：false`
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。


```go
func canFinish(numCourses int, prerequisites [][]int) bool {
    // Step 1: Create the graph and the in-degree array
    graph := make([][]int, numCourses)   // adjacency list
    inDegree := make([]int, numCourses)  // in-degree array
    
    // Build the graph and compute in-degrees
    for _, pre := range prerequisites {
        course, prerequisite := pre[0], pre[1]
        graph[prerequisite] = append(graph[prerequisite], course)
        inDegree[course]++
    }
    
    // Step 2: Initialize the queue with all courses that have in-degree 0
    queue := []int{}
    for i := 0; i < numCourses; i++ {
        if inDegree[i] == 0 {
            queue = append(queue, i)
        }
    }
    
    // Step 3: Perform Kahn's algorithm
    count := 0
    for len(queue) > 0 {
        course := queue[0]
        queue = queue[1:]  // dequeue the course
        count++
        
        // Decrease the in-degree of each neighbor
        for _, neighbor := range graph[course] {
            inDegree[neighbor]--
            if inDegree[neighbor] == 0 {
                queue = append(queue, neighbor)  // if in-degree becomes 0, add to queue
            }
        }
    }
    
    // Step 4: If we have visited all courses, return true
    return count == numCourses
}


```

## 代码解释：
### 1. 创建图和入度数组
```go
graph := make([][]int, numCourses)  // adjacency list
inDegree := make([]int, numCourses) // in-degree array

```
- `graph` 是一个邻接表，表示每个课程的依赖关系（即每个课程依赖哪些课程）。
- `inDegree` 是一个数组，记录每个课程的入度，即有多少个课程依赖于该课程。

### 2. 构建图和计算每个课程的入度
```go
// Build the graph and compute in-degrees
for _, pre := range prerequisites {
course, prerequisite := pre[0], pre[1]
graph[prerequisite] = append(graph[prerequisite], course)
inDegree[course]++
}

```

- 通过遍历 prerequisites 数组，构建有向图：
  - 对于每个 pre = [ai, bi]，表示课程 ai 依赖于课程 bi。
  - 所以在图 graph 中，graph[bi] 会包含课程 ai，表示课程 bi 是课程 ai 的先修课程。
  - 同时，inDegree[ai]++ 表示课程 ai 被依赖的次数增加 1。

