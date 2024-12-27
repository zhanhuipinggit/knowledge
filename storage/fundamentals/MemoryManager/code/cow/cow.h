//
// Created by zhanJames on 2024/12/27.
//

#ifndef KNOWLEDGE_COW_H
#define KNOWLEDGE_COW_H

#endif //KNOWLEDGE_COW_H
// 定义共享字符串结构体
typedef struct {
    char *data;       // 实际存储的字符串
    int ref_count;    // 引用计数
} SharedString;

SharedString* create_shared_string(const char *str);
void add_reference(SharedString *shared);
void release_reference(SharedString *shared);
void modify_shared_string(SharedString **shared, const char *new_value);
void print_shared_string(SharedString *shared);