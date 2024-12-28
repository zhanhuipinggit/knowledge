#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// 定义内存池结构体
typedef struct MemoryPool {
    void *memory;         // 内存池起始地址
    size_t block_size;    // 单个内存块大小
    int block_count;      // 内存块数量
    void **free_blocks;   // 空闲块列表
    int free_count;       // 空闲块数量
} MemoryPool;

// 初始化内存池
MemoryPool* create_memory_pool(size_t block_size, int block_count) {
    MemoryPool *pool = (MemoryPool *)malloc(sizeof(MemoryPool));
    pool->block_size = block_size;
    pool->block_count = block_count;
    pool->free_count = block_count;

    // 分配内存池
    pool->memory = malloc(block_size * block_count);
    pool->free_blocks = (void **)malloc(block_count * sizeof(void *));

    // 初始化空闲块链表
    for (int i = 0; i < block_count; i++) {
        pool->free_blocks[i] = (char *)pool->memory + i * block_size;
    }

    return pool;
}

// 分配内存块
void* memory_pool_alloc(MemoryPool *pool) {
    if (pool->free_count == 0) {
        return NULL; // 无空闲块
    }

    void *block = pool->free_blocks[--pool->free_count];
    return block;
}

// 释放内存块
void memory_pool_free(MemoryPool *pool, void *block) {
    pool->free_blocks[pool->free_count++] = block;
}

// 销毁内存池
void destroy_memory_pool(MemoryPool *pool) {
    free(pool->memory);
    free(pool->free_blocks);
    free(pool);
}

// 示例测试代码
int main() {
    size_t block_size = 32; // 每块大小 32 字节
    int block_count = 10;  // 总共 10 块

    // 创建内存池
    MemoryPool *pool = create_memory_pool(block_size, block_count);

    // 分配三个内存块
    void *block1 = memory_pool_alloc(pool);
    void *block2 = memory_pool_alloc(pool);
    void *block3 = memory_pool_alloc(pool);

    printf("Allocated blocks: %p, %p, %p\n", block1, block2, block3);

    // 释放一个内存块
    memory_pool_free(pool, block2);
    printf("Block %p freed.\n", block2);

    // 再次分配一个内存块
    void *block4 = memory_pool_alloc(pool);
    printf("Reallocated block: %p\n", block4);

    // 销毁内存池
    destroy_memory_pool(pool);
    return 0;
}