# 去重算法

去重算法的目的是从一个数据集合中移除重复的元素，保留唯一值。以下是几种常见的去重算法及其应用场景：

---

## **1. 使用哈希表去重**
哈希表（如 Python 的 `set`）具有快速查找的特性，能高效地实现去重。

**步骤**：
1. 遍历原数据集合。
2. 使用哈希表存储每个元素，如果已经存在则跳过，否则添加。

**代码示例（Python）**：
```python
def remove_duplicates(arr):
    return list(set(arr))

arr = [1, 2, 2, 3, 4, 4, 5]
print(remove_duplicates(arr))  # 输出: [1, 2, 3, 4, 5]
```

**优点**：
- 时间复杂度为 \(O(n)\)。
- 实现简单。

**缺点**：
- 无法保证结果的顺序（无序集合）。

---

## **2. 使用双指针法（适用于有序数组）**
双指针法适合对已排序的数组进行去重，效率更高。

**步骤**：
1. 排序数组。
2. 使用两个指针，一个遍历数组，另一个指向结果数组末尾。

**代码示例（Python）**：
```python
def remove_duplicates_sorted(arr):
    if not arr:
        return []
    
    arr.sort()  # 确保有序
    result = [arr[0]]
    
    for i in range(1, len(arr)):
        if arr[i] != result[-1]:  # 只添加与结果数组最后一个元素不同的值
            result.append(arr[i])
    
    return result

arr = [4, 2, 1, 2, 4, 3]
print(remove_duplicates_sorted(arr))  # 输出: [1, 2, 3, 4]
```

**优点**：
- 空间复杂度低。
- 保留原始顺序（排序后）。

**缺点**：
- 需要先排序，时间复杂度为 \(O(n \log n)\)。

---

## **3. 使用索引查找去重**
适合对较小数据集进行去重，并保留原顺序。

**步骤**：
1. 遍历原数组。
2. 如果当前元素未出现过，则将其加入结果数组。

**代码示例（Python）**：
```python
def remove_duplicates_preserve_order(arr):
    seen = set()
    result = []
    for x in arr:
        if x not in seen:
            result.append(x)
            seen.add(x)
    return result

arr = [4, 2, 1, 2, 4, 3]
print(remove_duplicates_preserve_order(arr))  # 输出: [4, 2, 1, 3]
```

**优点**：
- 保留原始顺序。
- 时间复杂度为 \(O(n)\)。

**缺点**：
- 需要额外的空间存储哈希表。

---

## **4. 基于位运算的去重（针对整数）**
对于范围较小的整数集合，可以使用位运算记录某个数是否已经存在。

**步骤**：
1. 用一个位数组记录每个数是否已经出现。
2. 遍历集合，通过位运算判断并记录。

**代码示例（Python）**：
```python
def remove_duplicates_with_bits(arr):
    bit_set = 0
    result = []
    for num in arr:
        if not (bit_set & (1 << num)):  # 检查第 num 位是否为 1
            result.append(num)
            bit_set |= (1 << num)  # 将第 num 位置为 1
    return result

arr = [1, 2, 3, 2, 1, 4]
print(remove_duplicates_with_bits(arr))  # 输出: [1, 2, 3, 4]
```

**优点**：
- 内存占用小。
- 适合小范围整数。

**缺点**：
- 仅适用于整数。
- 无法处理过大范围的整数。

---

## **5. SQL 去重**
如果数据存储在数据库中，可以使用 SQL 的 `DISTINCT` 关键字进行去重。

**示例**：
```sql
SELECT DISTINCT column_name
FROM table_name;
```

---

## **应用场景总结**
- 数据量较大且无序：**哈希表去重**。
- 数据有序：**双指针法**。
- 数据量小，需保留顺序：**索引查找去重**。
- 整数范围固定：**位运算去重**。
- 数据存储在数据库：**SQL 去重**。
