# 回溯算法详解

回溯算法是一种系统性地搜索问题解空间的算法，它通过构造一个搜索树，逐步尝试所有可能的选择，直到找到满足条件的解或者遍历所有可能性。

## 基本概念

回溯算法是一种试探法，主要用于解决以下类型的问题：

1. **组合问题**：从给定的元素集合中选择满足条件的子集。
2. **排列问题**：确定元素的所有可能顺序。
3. **划分问题**：将集合划分为满足特定条件的子集。
4. **路径问题**：如迷宫问题、八皇后问题等。

### 核心思想

回溯算法的核心思想是：
1. 将问题分解为多个阶段，每个阶段都有若干可选择的状态。
2. 通过搜索树的方式逐步尝试不同的选择。
3. 如果某个选择导致不满足条件，则回退到上一步继续尝试其他选择。

## 算法框架

回溯算法的伪代码如下：

```python
# 回溯框架伪代码
def backtrack(路径, 选择列表):
    if 满足结束条件:
        记录结果
        return

    for 选择 in 选择列表:
        做出选择
        backtrack(路径, 更新后的选择列表)
        撤销选择
```

### 关键步骤

1. **选择列表**：当前可供选择的选项。
2. **路径**：当前已经做出的选择。
3. **结束条件**：判断路径是否满足问题要求。
4. **回溯**：当路径不满足条件时，撤销最近的选择，回到上一步继续探索。

## 示例问题

### 示例 1：全排列问题

给定一个不含重复数字的数组 `nums`，返回其所有可能的全排列。

#### 解法代码
```python
def permute(nums):
    res = []

    def backtrack(path, choices):
        if not choices:
            res.append(path[:])  # 保存结果
            return

        for i in range(len(choices)):
            path.append(choices[i])  # 做出选择
            backtrack(path, choices[:i] + choices[i+1:])  # 更新选择列表
            path.pop()  # 撤销选择

    backtrack([], nums)
    return res
```

#### 执行过程
以输入 `nums = [1, 2, 3]` 为例：
1. 初始状态：路径 `[]`，选择列表 `[1, 2, 3]`。
2. 选择 `1`，路径更新为 `[1]`，选择列表更新为 `[2, 3]`。
3. 递归选择 `2`，路径更新为 `[1, 2]`，选择列表更新为 `[3]`。
4. 最终路径为 `[1, 2, 3]`，保存结果并回溯。

### 示例 2：N 皇后问题

N 皇后问题要求在 N\*N 的棋盘上放置 N 个皇后，使得它们互相不威胁。

#### 解法代码
```python
def solveNQueens(n):
    res = []
    board = ["." * n for _ in range(n)]

    def is_valid(row, col):
        for i in range(row):
            if board[i][col] == "Q" or \
               (col - (row - i) >= 0 and board[i][col - (row - i)] == "Q") or \
               (col + (row - i) < n and board[i][col + (row - i)] == "Q"):
                return False
        return True

    def backtrack(row):
        if row == n:
            res.append(["".join(row) for row in board])
            return

        for col in range(n):
            if not is_valid(row, col):
                continue
            board[row] = board[row][:col] + "Q" + board[row][col+1:]
            backtrack(row + 1)
            board[row] = board[row][:col] + "." + board[row][col+1:]

    backtrack(0)
    return res
```

## 应用场景

1. **组合优化问题**：如背包问题、旅行商问题等。
2. **博弈问题**：如数独、数独求解等。
3. **路径搜索**：如迷宫问题、找最短路径等。
4. **排列组合问题**：如电话号码组合、字母排列等。

## 优化方法

1. **剪枝**：通过提前判断某些选择是否可能满足条件，减少搜索空间。
2. **位运算**：对状态进行编码，利用位运算加速计算。
3. **记忆化搜索**：保存已经计算过的子问题结果，避免重复计算。

## 总结

回溯算法是一种强大的搜索工具，通过系统性地探索所有可能性，可以解决许多复杂的问题。但其时间复杂度较高，因此在实际应用中通常需要配合剪枝等优化技术。

