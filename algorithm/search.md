# 搜索算法详细描述

搜索算法用于在给定的数据结构（如数组、树或图）中查找目标元素或路径。以下是常见搜索算法的分类和详细描述。

## 一、线性搜索

### 1. 顺序搜索（Linear Search）
**描述：**
- 顺序遍历列表中的每个元素，直到找到目标元素或遍历完整个列表。
- 时间复杂度：O(n)。

**适用场景：**
- 数据未排序。

**伪代码：**
```plaintext
function linearSearch(array, target):
    for i = 0 to array.length - 1:
        if array[i] == target:
            return i
    return -1
```

---

## 二、二分搜索

### 1. 二分查找（Binary Search）
**描述：**
- 仅适用于排序数据。
- 每次将搜索范围缩小一半，直到找到目标或搜索范围为空。
- 时间复杂度：O(log n)。

**适用场景：**
- 数据已排序。

**伪代码：**
```plaintext
function binarySearch(array, target):
    left = 0
    right = array.length - 1

    while left <= right:
        mid = left + (right - left) / 2
        if array[mid] == target:
            return mid
        else if array[mid] < target:
            left = mid + 1
        else:
            right = mid - 1

    return -1
```

---

## 三、图搜索算法

### 1. 深度优先搜索（Depth-First Search, DFS）
**描述：**
- 遍历图时优先深入每一个分支，直到无法前进为止，然后回溯。
- 可以通过栈（递归实现或显式栈）实现。
- 时间复杂度：O(V + E)（V 为顶点数，E 为边数）。

**适用场景：**
- 图遍历，路径查找。

**伪代码：**
```plaintext
function DFS(node, visited):
    if node in visited:
        return
    visited.add(node)
    for neighbor in node.neighbors:
        DFS(neighbor, visited)
```

### 2. 广度优先搜索（Breadth-First Search, BFS）
**描述：**
- 遍历图时按层级顺序（从近到远）逐步访问每个节点。
- 使用队列实现。
- 时间复杂度：O(V + E)。

**适用场景：**
- 最短路径查找，无权图遍历。

**伪代码：**
```plaintext
function BFS(startNode):
    queue = [startNode]
    visited = set()

    while queue is not empty:
        node = queue.pop(0)
        if node not in visited:
            visited.add(node)
            for neighbor in node.neighbors:
                queue.append(neighbor)
```

---

## 四、启发式搜索

### 1. A* 算法
**描述：**
- 基于启发式估价函数的搜索算法，通常用于路径规划。
- 通过 \( f(n) = g(n) + h(n) \) 评估节点，其中：
    - \( g(n) \): 从起点到当前节点的实际代价。
    - \( h(n) \): 从当前节点到目标的估计代价。
- 时间复杂度：依赖于启发函数的效率。

**适用场景：**
- 最短路径查找，地图导航。

**伪代码：**
```plaintext
function A*(startNode, goalNode):
    openSet = {startNode}
    gScore = {startNode: 0}
    fScore = {startNode: heuristic(startNode, goalNode)}

    while openSet is not empty:
        current = node in openSet with lowest fScore
        if current == goalNode:
            return reconstructPath(current)

        openSet.remove(current)
        for neighbor in current.neighbors:
            tentative_gScore = gScore[current] + dist(current, neighbor)
            if tentative_gScore < gScore[neighbor]:
                gScore[neighbor] = tentative_gScore
                fScore[neighbor] = gScore[neighbor] + heuristic(neighbor, goalNode)
                if neighbor not in openSet:
                    openSet.add(neighbor)
```

---

## 五、总结

| 算法类别         | 时间复杂度         | 数据要求       | 适用场景               |
|------------------|-------------------|---------------|-----------------------|
| 顺序搜索         | O(n)             | 无要求         | 未排序数据查找         |
| 二分搜索         | O(log n)         | 排序数据       | 高效查找目标元素       |
| 深度优先搜索     | O(V + E)         | 图结构         | 路径搜索，连通性检查   |
| 广度优先搜索     | O(V + E)         | 图结构         | 无权最短路径查找       |
| A* 算法          | 依赖启发函数      | 图结构         | 最优路径规划           |

---

搜索算法的选择取决于具体应用场景和数据结构的特性。在实际应用中，可以根据问题的特点综合使用这些算法，或对其进行优化和改进。
