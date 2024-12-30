# Bitmap 去重详解

## 背景与问题

对于大规模的数字去重问题（如 40 亿个 10 位数字），直接将所有数据加载到内存进行去重可能不现实，主要受限于内存容量和计算效率。位图（Bitmap）是一种高效的去重方式，利用位的存储特性，可以显著降低内存开销。

## 位图去重的核心思想

1. **位图简介**：
   - 位图是一个按位存储的数组，用于高效地表示大量可能的值。
   - 每个位（bit）对应一个数字的存在状态，`1` 表示数字已出现，`0` 表示数字未出现。

2. **数字范围与位图大小**：
   - 对于 10 位数字（范围 `0 ~ 9,999,999,999`），最多需要标记 \(10^10\) 个值。
   - 位图需要的内存大小为：
     10^10/8/1024/1024/1024 = 1.2G
   - 如果只考虑部分数据（如实际只有 40 亿个数字），内存需求可进一步优化。

3. **存储映射**：
   - 每个数字通过以下映射方式找到其对应的位：
      - 字节索引：byte_index = num / 8
      - 位索引：bit_index = num % 8

4. **去重逻辑**：
   - 初始化位图，所有位为 `0`。
   - 遍历数字列表，将每个数字对应的位设置为 `1`。
   - 最终提取位图中为 `1` 的数字作为去重结果。

## 算法实现

### Python 示例代码
以下是基于位图的去重实现：

```python
import numpy as np

def bitmap_deduplication(numbers):
    max_number = 10**10  # 最大10位数字
    bitmap = np.zeros(max_number // 8 + 1, dtype=np.uint8)  # 位图初始化

    # 设置位图
    for num in numbers:
        byte_index = num // 8
        bit_index = num % 8
        bitmap[byte_index] |= (1 << bit_index)  # 将对应位设置为1

    # 提取去重结果
    result = []
    for i in range(len(bitmap)):
        for j in range(8):
            if bitmap[i] & (1 << j):
                result.append(i * 8 + j)

    return result

# 测试数据
numbers = [1234567890, 1234567890, 9876543210]
unique_numbers = bitmap_deduplication(numbers)
print(unique_numbers)
```

### Go 实现
```go
package main

import (
	"fmt"
)

func bitmapDeduplication(numbers []int) []int {
	const maxNumber = 10000000000
	bitmap := make([]byte, maxNumber/8+1)

	// 设置位图
	for _, num := range numbers {
		byteIndex := num / 8
		bitIndex := num % 8
		bitmap[byteIndex] |= 1 << bitIndex
	}

	// 提取去重结果
	result := []int{}
	for i, b := range bitmap {
		for j := 0; j < 8; j++ {
			if b&(1<<j) != 0 {
				result = append(result, i*8+j)
			}
		}
	}

	return result
}

func main() {
	numbers := []int{1234567890, 1234567890, 9876543210}
	uniqueNumbers := bitmapDeduplication(numbers)
	fmt.Println(uniqueNumbers)
}
```

### C 实现
```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_NUMBER 10000000000ULL

void bitmap_deduplication(unsigned long long *numbers, size_t size) {
    size_t bitmap_size = MAX_NUMBER / 8 + 1;
    unsigned char *bitmap = (unsigned char *)calloc(bitmap_size, sizeof(unsigned char));

    // 设置位图
    for (size_t i = 0; i < size; i++) {
        unsigned long long num = numbers[i];
        size_t byte_index = num / 8;
        size_t bit_index = num % 8;
        bitmap[byte_index] |= (1 << bit_index);
    }

    // 提取去重结果
    for (size_t i = 0; i < bitmap_size; i++) {
        for (int j = 0; j < 8; j++) {
            if (bitmap[i] & (1 << j)) {
                printf("%llu\n", i * 8 + j);
            }
        }
    }

    free(bitmap);
}

int main() {
    unsigned long long numbers[] = {1234567890, 1234567890, 9876543210};
    size_t size = sizeof(numbers) / sizeof(numbers[0]);
    bitmap_deduplication(numbers, size);
    return 0;
}
```

### C++ 实现
```cpp
#include <iostream>
#include <vector>
#include <cstring>

#define MAX_NUMBER 10000000000ULL

void bitmapDeduplication(const std::vector<unsigned long long>& numbers) {
    size_t bitmapSize = MAX_NUMBER / 8 + 1;
    std::vector<unsigned char> bitmap(bitmapSize, 0);

    // 设置位图
    for (auto num : numbers) {
        size_t byteIndex = num / 8;
        size_t bitIndex = num % 8;
        bitmap[byteIndex] |= (1 << bitIndex);
    }

    // 提取去重结果
    for (size_t i = 0; i < bitmapSize; i++) {
        for (int j = 0; j < 8; j++) {
            if (bitmap[i] & (1 << j)) {
                std::cout << i * 8 + j << std::endl;
            }
        }
    }
}

int main() {
    std::vector<unsigned long long> numbers = {1234567890, 1234567890, 9876543210};
    bitmapDeduplication(numbers);
    return 0;
}
```

### Java 实现
```java
import java.util.*;

public class BitmapDeduplication {
    public static void bitmapDeduplication(List<Long> numbers) {
        final long MAX_NUMBER = 10_000_000_000L;
        byte[] bitmap = new byte[(int) (MAX_NUMBER / 8 + 1)];

        // 设置位图
        for (long num : numbers) {
            int byteIndex = (int) (num / 8);
            int bitIndex = (int) (num % 8);
            bitmap[byteIndex] |= (1 << bitIndex);
        }

        // 提取去重结果
        for (int i = 0; i < bitmap.length; i++) {
            for (int j = 0; j < 8; j++) {
                if ((bitmap[i] & (1 << j)) != 0) {
                    System.out.println((long) i * 8 + j);
                }
            }
        }
    }

    public static void main(String[] args) {
        List<Long> numbers = Arrays.asList(1234567890L, 1234567890L, 9876543210L);
        bitmapDeduplication(numbers);
    }
}
```

### Rust 实现
```rust
fn bitmap_deduplication(numbers: &[u64]) {
    const MAX_NUMBER: u64 = 10_000_000_000;
    let bitmap_size = (MAX_NUMBER / 8 + 1) as usize;
    let mut bitmap = vec![0u8; bitmap_size];

    // 设置位图
    for &num in numbers {
        let byte_index = (num / 8) as usize;
        let bit_index = (num % 8) as u8;
        bitmap[byte_index] |= 1 << bit_index;
    }

    // 提取去重结果
    for (i, &byte) in bitmap.iter().enumerate() {
        for j in 0..8 {
            if byte & (1 << j) != 0 {
                println!("{}", i as u64 * 8 + j);
            }
        }
    }
}

fn main() {
    let numbers = vec![1234567890, 1234567890, 9876543210];
    bitmap_deduplication(&numbers);
}
```

### PHP 实现
```php
<?php
function bitmapDeduplication(array $numbers) {
    $maxNumber = 10000000000;
    $bitmap = str_repeat("\0", $maxNumber / 8 + 1);

    // 设置位图
    foreach ($numbers as $num) {
        $byteIndex = intdiv($num, 8);
        $bitIndex = $num % 8;
        $bitmap[$byteIndex] = $bitmap[$byteIndex] | (1 << $bitIndex);
    }

    // 提取去重结果
    $result = [];
    for ($i = 0; $i < strlen($bitmap); $i++) {
        for ($j = 0; $j < 8; $j++) {
            if ((ord($bitmap[$i]) & (1 << $j)) !== 0) {
                $result[] = $i * 8 + $j;
            }
        }
    }

    return $result;
}

// 测试数据
$numbers = [1234567890, 1234567890, 9876543210];
$uniqueNumbers = bitmapDeduplication($numbers);
print_r($uniqueNumbers);
```

## 优化建议

1. **分段处理**：
   - 如果内存不足，可以将数字分区处理，分块加载位图。

2. **文件存储**：
   - 位图可以序列化存储到磁盘，便于处理超大规模数据。

3. **并行化**：
   - 使用多线程或分布式计算框架（如 Hadoop/Spark），对数据进行并行去重。

## 位图的优缺点

### 优点
- **高效**：内存使用量低，处理速度快。
- **可扩展**：适合处理大规模数据。

### 缺点
- **仅适用于整数**：需对数据进行映射，无法直接处理非整数数据。
- **最大值限制**：位图大小取决于数据范围，对于非常大的范围可能不适用。

## 适用场景
- 需要对大规模整数数据去重。
- 数据范围固定且内存有限。
- 需要高效的时间和空间复杂度。


