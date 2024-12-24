# 快速排序（Quick Sort）

快速排序（Quick Sort）是由 Tony Hoare 在 1960 年提出的一个高效的排序算法。它采用分治法（Divide and Conquer）策略，将一个大的数组分成两个子数组，然后递归地对这两个子数组进行排序。

## 算法简介

快速排序的基本思想是：

1. **选择基准元素**：从待排序数组中选取一个元素作为基准元素（pivot）。
2. **分区操作**：将数组重新排列，将小于基准元素的元素放在基准元素的左边，大于基准元素的元素放在右边。经过此操作，基准元素就排到了它最终的位置。
3. **递归排序**：分别对基准元素左边和右边的两个子数组进行快速排序，直到数组长度为 1 或者为空，此时排序完成。

快速排序的平均时间复杂度是 `O(n log n)`，但在最坏情况下（当数组已经有序或逆序时），时间复杂度可能退化为 `O(n^2)`。不过，快速排序通常能较好地工作，且空间复杂度为 `O(log n)`，因此它在许多排序问题中得到广泛应用。

## 快速排序的工作原理

快速排序的核心思想是通过 **分治法** 将大问题转化为小问题。具体步骤如下：

1. **选择基准元素**：通常选取数组的第一个元素、最后一个元素或者随机选取一个元素作为基准。
2. **分区**：通过一个 `partition` 函数将数组分为两个部分：
    - 左边部分的元素都小于等于基准元素。
    - 右边部分的元素都大于等于基准元素。
3. **递归排序**：分别对基准元素左边和右边的两个子数组递归进行排序。

### 快速排序的过程

以数组 `[9, 7, 5, 11, 12, 2, 14, 3, 10, 6]` 为例，假设我们选择数组的最后一个元素作为基准：

#### 第 1 步：选择基准元素并进行分区

选择基准元素 `6`，然后将小于等于 6 的元素移到基准元素左侧，大于 6 的元素移到右侧。分区后的数组如下：
`[5, 2, 3, 6, 12, 11, 14, 7, 10, 9] `↑ 基准元素 6


此时，基准元素 `6` 已经排到正确的位置，它的左侧是所有小于等于 6 的元素，右侧是所有大于等于 6 的元素。

#### 第 2 步：递归对左右部分进行排序

- 对基准元素左侧的 `[5, 2, 3]` 进行快速排序。
- 对基准元素右侧的 `[12, 11, 14, 7, 10, 9]` 进行快速排序。

#### 左侧排序过程

对左侧子数组 `[5, 2, 3]` 进行快速排序：

1. **选择基准元素**：选择最后一个元素 `3` 作为基准。
2. **分区**：将小于等于 `3` 的元素放到左边，大于 `3` 的元素放到右边。

分区后的数组：
`[2, 3, 5]` ↑ 基准元素 3


此时，`3` 已经排到正确的位置。接下来递归对左右部分排序：

- 左侧子数组 `[2]` 只有一个元素，不需要排序。
- 右侧子数组 `[5]` 只有一个元素，不需要排序。

#### 右侧排序过程

对右侧子数组 `[12, 11, 14, 7, 10, 9]` 进行快速排序：

1. **选择基准元素**：选择最后一个元素 `9` 作为基准。
2. **分区**：将小于等于 `9` 的元素放到左边，大于 `9` 的元素放到右边。

分区后的数组：

`[7, 9, 11, 12, 10, 14]` ↑ 基准元素 9


接下来递归对左右部分排序：

- 左侧子数组 `[7]` 只有一个元素，不需要排序。
- 右侧子数组 `[11, 12, 10, 14]` 进行排序。

重复上述步骤，直到所有子数组排序完成。

### 最终结果

通过不断地分区和递归，快速排序最终会将数组排序为：

`[2, 3, 5, 6, 7, 9, 10, 11, 12, 14]`


## 快速排序的时间复杂度分析

- **平均时间复杂度**：`O(n log n)`。由于每次分区操作会将数组大致平分为两部分，递归树的深度为 `log n`，每一层的工作量为 `O(n)`，所以总体复杂度为 `O(n log n)`。
- **最坏时间复杂度**：`O(n^2)`。当每次分区都将数组划分成只有一个元素的子数组时，递归树的深度为 `n`，导致最坏情况的复杂度为 `O(n^2)`。这种情况通常出现在数组已经排序或者是逆序排列的情况下。
- **最好时间复杂度**：`O(n log n)`。理想情况下，每次分区操作都能将数组平分成两部分。

### 空间复杂度

快速排序的空间复杂度为 `O(log n)`，主要用于递归调用栈。由于快速排序是原地排序，不需要额外的存储空间。

## 快速排序的优缺点

### 优点：
- **平均时间复杂度低**：在大多数情况下，快速排序比其他 `O(n log n)` 算法（如归并排序）更快，因为它的常数因素较小。
- **原地排序**：快速排序不需要额外的存储空间，空间复杂度较低。

### 缺点：
- **最坏情况性能差**：如果每次选择的基准元素不合适，导致分区不平衡，性能会退化为 `O(n^2)`。例如，在已经有序或者逆序的情况下，快速排序的效率较低。
- **不稳定排序**：快速排序不是稳定的排序算法，可能改变相同元素的相对顺序。

## 结论

快速排序是一个非常高效的排序算法，尤其适合大规模数据的排序。通过选择合适的基准元素和优化分区策略，能够在大多数情况下保持较高的性能。然而，在某些特殊情况下，它的性能可能退化为 `O(n^2)`，这时可以考虑使用其他排序算法，如归并排序或堆排序。

**go**
```go

type quickSort struc{}

func (q *quickSort) partition(nums []int, left, right int) int {
	i, j := left, right
	for i < j {
        for i<j && nums[i] < nums[left] {
            i++
        }
        for i < j && nums[j] > nums[left] {
            j--
        }
        nums[i],nums[j] = nums[j], nums[i]
    }
	nums[i], nums[left] = nums[left], nums[i]
	return i
}

func (q *quickSort)quickSort(nums []int, left, right int) {
	if left > right {
		return
    }
	pivot := q.partition(nums, left, right)
	q.quickSort(nums, left, pivot-1)
	q.quickSort(nums, pivot+1, right)
	
}

```

**c++**
```cpp
#include <iostream>
#include <vector>
using namespace std;

int partition(vector<int>& nums, int left, int right) {
    int pivotNum = nums[left];  // 选择第一个元素作为 pivot
    int i = left + 1, j = right;

    while (i <= j) {
        // 从左边找第一个大于 pivot 的元素
        while (i <= right && nums[i] <= pivotNum) {
            i++;
        }
        // 从右边找第一个小于 pivot 的元素
        while (j >= left + 1 && nums[j] > pivotNum) {
            j--;
        }

        if (i < j) {
            swap(nums[i], nums[j]);  // 交换
        }
    }

    // 将 pivot 元素放到正确的位置
    swap(nums[left], nums[j]);
    return j;  // 返回 pivot 的最终位置
}

void quickSort(vector<int>& nums, int left, int right) {
    if (left >= right) {
        return;  // 递归终止条件
    }

    int pivot = partition(nums, left, right);  // 获取 pivot 的最终位置
    quickSort(nums, left, pivot - 1);  // 递归排序左边部分
    quickSort(nums, pivot + 1, right);  // 递归排序右边部分
}


```

**c**
```c
#include <stdio.h>

/* 元素交换 */
void swap(int nums[], int i, int j) {
    int tmp = nums[i];
    nums[i] = nums[j];
    nums[j] = tmp;
}

int partition(int nums[], int left, int right) {
    int pivot = nums[left];
    int i = left, j = right;
    while (i < j) {
        while (i < j && nums[j] >= pivot) {
            j--;
        }
        while(i < j && nums[i] <= pivot) {
            i++;
        }

        swap(nums, i, j);
    }
    swap(nums, i, left);

    return i;

}

void quickSort(int nums[], int left, int right) {
    if ( left >= right) {
        return;
    }

    int pivot = partition(nums, left , right);
    quickSort(nums, left, pivot-1);
    quickSort(nums, pivot+1, right);
}



```
**python**
```python
def quick_sort(nums, left, right):
    if left >= right:
        return

    # 选择第一个元素作为 pivot
    pivot = partition(nums, left, right)

    # 递归排序左右部分
    quick_sort(nums, left, pivot - 1)
    quick_sort(nums, pivot + 1, right)


def partition(nums, left, right):
    pivot = nums[left]  # 选择第一个元素作为 pivot
    i = left + 1
    j = right
    while i < j:
        while i < j and nums[j] >= pivot:
            j -= 1
        while i < j and nums[i] <= pivot:
            i += 1
        nums[i], nums[j] = nums[j], nums[i]

    nums[i], nums[left] = nums[left], nums[i]

    return i

```

**rust**
```rust
fn partition(nums: &mut [i32], left: usize, right usize)-> usize {
   let pivot = nums[left];
   let mut i = left;
   let mut j = right;
   while i < j {
      while i < j && nums[j] >= pivot {
         j-=1;
      }
      while i < j && nums[i] <= pivot {
         i+=1;
      }
      nums.swap(i, j);
   }
   nums.swap(i, left);
   i
}

fn quick_sort(nums: &mut [i32], left: i32, right: i32) {
   if left >= right {
      return;
   }
   let pivot = Self::partition(nums, left as usize, right as usize) as i32;
   Self::quick_sort(nums, left, pivot-1);
   Self::quick_sort(nums, pivot+1, right);
}


```