# 堆排序（Heap Sort）详细介绍

堆排序是一种基于比较的排序算法，它利用了堆这一数据结构来实现排序。堆是一棵完全二叉树，它满足堆的性质：

- **最大堆（Max Heap）**：每个父节点的值都大于或等于其子节点的值。
- **最小堆（Min Heap）**：每个父节点的值都小于或等于其子节点的值。

堆排序的主要思想是通过构建一个最大堆来将最大的元素不断提取出来，形成有序的数组。

## 堆排序的步骤

1. **构建最大堆**：
    - 通过调整数组元素，使其满足最大堆的性质。最大堆是一个二叉树，其中每个父节点的值都大于或等于它的左右子节点。

2. **交换堆顶元素与最后一个元素**：
    - 将堆顶的最大元素与堆的最后一个元素交换，将最大元素放置在已排序的部分。

3. **重新调整堆**：
    - 由于交换了堆顶和最后一个元素，可能破坏了堆的结构，需要对堆进行重新调整（即“堆化”），确保剩下的部分依然满足最大堆的性质。

4. **重复**：
    - 重复步骤 2 和 3，直到堆的大小为 1。

## 堆排序的时间复杂度

- **构建堆的时间复杂度**：`O(n)`。虽然堆化的时间复杂度是 `O(log n)`，但在构建堆的过程中，堆化的节点数逐渐减少，因此总体是 `O(n)`。

- **排序过程的时间复杂度**：`O(n log n)`。每次交换后需要重新堆化，堆化的时间复杂度是 `O(log n)`，总共执行 `n` 次交换，所以时间复杂度是 `O(n log n)`。

- **空间复杂度**：`O(1)`。堆排序是原地排序算法，只需要常数级别的额外空间。

## 堆排序的优缺点

**优点**：
- **时间复杂度稳定**：堆排序的时间复杂度始终是 `O(n log n)`，无论数据的初始顺序如何。
- **不需要额外的空间**：堆排序是一个原地排序算法，只需要常数级别的额外空间。

**缺点**：
- **不稳定排序**：堆排序是一个不稳定的排序算法，这意味着相同值的元素在排序过程中可能会改变原来的相对顺序。
- **比快速排序慢**：尽管堆排序的时间复杂度与快速排序相同，但实际操作中，堆排序通常比快速排序要慢，因为堆排序涉及更多的交换和堆化操作。

## 堆排序的实现

以下是堆排序的具体实现步骤：


### 堆排序（Heap Sort）算法实现

#### 1. 构建最大堆（Max Heap）

首先，我们需要将输入数组构建成最大堆，使得每个父节点的值都大于或等于其子节点的值。可以从最后一个非叶子节点开始，逐步调整每个节点。

#### 2. 交换堆顶和最后一个元素

将堆顶的最大元素与堆的最后一个元素交换。交换后，堆的大小减少1，然后重新调整堆，使其保持最大堆的性质。

#### 3. 重复步骤 2 和 3

重复执行交换和堆化过程，直到堆的大小为 1，此时数组已经排序完成。


### 具体堆排序的步骤：
1. **构建最大堆**：将给定的数组构建成最大堆。最大堆中，每个父节点都大于等于其子节点。
2. **交换根节点和最后一个元素**：根节点（最大元素）与堆的最后一个元素交换。此时，根节点已经排好序。
3. **堆化**：交换后，堆的结构可能被破坏，需要重新调整堆来恢复堆的性质。
4. **重复**：对剩余的元素重复步骤 2 和步骤 3，直到所有元素都排序完成。

### 时间复杂度：
- 最佳/平均/最坏时间复杂度：**O(n log n)**
- 空间复杂度：**O(1)**（原地排序）

## 堆排序示例：

考虑以下数组：`[4, 10, 3, 5, 1]`。

### 步骤 1：构建最大堆
从最后一个非叶节点开始，调整每个子树以形成最大堆：

初始数组：
```text
      4
    /   \
   10    3
  /  \   
 5    1

```
调整后的最大堆：
```text
      10
    /    \
   5      3
  /  \   
 4    1

```

### 步骤 2：交换根节点和最后一个元素
将根节点（10）与最后一个元素（1）交换
`[1, 5, 3, 4, 10]`
### 步骤 3：堆化剩余元素
堆化剩余的堆：

```text
      5
    /   \
   4     3
  /  
 1

```

堆化后的结果：
`[5, 4, 3, 1, 10]`


### 步骤 4：重复该过程
重复步骤 2 和 3，直到数组完全排序。

### 最终排序后的数组：
`[1, 3, 4, 5, 10]`


```mermaid
graph TD;
    A[初始数组: [4, 10, 3, 5, 1]] --> B[构建最大堆];
    B --> C[交换根节点和最后一个元素];
    C --> D[堆化剩余数组];
    D --> E[重复直到排序完成];
    
    subgraph Build_Max_Heap
        A1[从最后一个非叶节点开始]
        A2[调整每个子树]
        A1 --> A2
    end

    subgraph Swap_Step
        B1[交换根节点和最后一个元素]
    end

    subgraph Heapify_Step
        C1[恢复堆的结构]
    end

```

### 代码实现
**go**

```go

func siftDown(nums *[]int, n , i int) {
	for true {
		l := 2*i +1
		r := 2*i+2
		m := i
        if (l < n && (*nums)[l] > (*nums)[m]) {
            m  = l
        }
        if (r < n && (*nums)[l] > (*nums)[m]) {
            m = r
            
        }
        if (m == i) {
            break
        }

        (*nums)[i], (*nums)[m] = (*nums)[m], (*nums)[i]
        i = ma
    }
	
}

func heapSort(nums *[]int) {
    for i := len(*nums)/2 -1; i >= 0 ; i-- {
        siftDown(nums, i, len(*nums))
    }

	for i := len(*nums)-1; i > 0; i-- {
        (*nums)[0], (*nums)[i] = (*nums)[i], (*nums)[0]
        siftDown(nums, i, 0)
    }
	
	
}








```


