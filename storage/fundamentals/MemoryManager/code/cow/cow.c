//
// Created by zhanJames on 2024/12/27.
//

#include "cow.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>


// 创建新共享字符串
SharedString* create_shared_string(const char *str) {
    SharedString *shared = (SharedString *)malloc(sizeof(SharedString));
    shared->data = strdup(str);  // 分配并复制字符串
    shared->ref_count = 1;       // 初始引用计数为 1
    return shared;
}

// 增加引用计数
void add_reference(SharedString *shared) {
    if (shared) {
        shared->ref_count++;
    }
}

// 减少引用计数并释放资源
void release_reference(SharedString *shared) {
    if (shared) {
        shared->ref_count--;
        if (shared->ref_count == 0) {  // 引用计数为 0 时释放资源
            free(shared->data);
            free(shared);
        }
    }
}

// 写时复制：检查引用计数并执行复制
void modify_shared_string(SharedString **shared, const char *new_value) {
    if ((*shared)->ref_count > 1) { // 引用计数大于 1 时复制
        SharedString *new_shared = create_shared_string(new_value);
        release_reference(*shared); // 释放当前引用
        *shared = new_shared;       // 更新为新数据
    } else { // 只有一个引用时直接修改
        free((*shared)->data);
        (*shared)->data = strdup(new_value);
    }
}

// 打印共享字符串状态
void print_shared_string(SharedString *shared) {
    if (shared) {
        printf("String: %s, Ref Count: %d\n", shared->data, shared->ref_count);
    }
}

// 示例测试代码
int main() {
    SharedString *str1 = create_shared_string("Hello, World!");
    SharedString *str2 = str1;  // str2 共享 str1 的数据

    add_reference(str1);  // 增加引用计数
    print_shared_string(str1);
    print_shared_string(str2);

    // 修改 str1（写时复制会生效）
    modify_shared_string(&str1, "Hello, COW!");
    print_shared_string(str1);
    print_shared_string(str2);

    // 释放引用
    release_reference(str1);
    release_reference(str2);

    return 0;
}
