 [English](./README.md) | 中文
# 介绍
### 存储知识（数据存储知识）

- **数据存储基础**：了解数据存储的基础知识，包括不同类型的存储介质（例如硬盘驱动器HDD、固态硬盘SSD、NVMe）、存储架构（例如直接附加存储DAS、网络附加存储NAS、存储区域网络SAN）和存储协议（例如SCSI、SATA、NVMe）。
- **数据库管理系统（DBMS）**：了解数据库的工作原理，包括关系型数据库（SQL）和非关系型数据库NoSQL，它们的模式以及如何使用SQL或其他查询语言与它们交互。
- **数据持久化**：了解数据如何保存和检索，包括文件系统、对象存储和数据序列化格式（例如JSON、XML、BSON）的使用。
- **数据冗余和复制**：熟悉确保数据持久性的技术，例如RAID配置、数据镜像和跨不同地理位置的复制。
- **数据备份和恢复**：了解备份策略，包括全备份、增量备份和差异备份，以及在数据丢失情况下的恢复流程。
- **存储优化**：掌握优化存储性能和容量的技术，包括数据去重、压缩和分层存储策略。

### 算法知识（算法知识）

- **算法复杂度**：理解时间和空间复杂度、大O表示法以及如何分析算法的效率。
- **数据结构**：了解各种数据结构，如数组、链表、栈、队列、树（例如二叉搜索树、AVL树）、图和哈希表。
- **排序和搜索算法**：熟悉不同的排序算法（例如快速排序、归并排序、堆排序）和搜索算法（例如二分搜索、深度优先搜索DFS、广度优先搜索BFS）。
- **图算法**：了解图遍历技术（例如迪杰斯特拉算法、A*算法、贝尔曼-福特算法）和特定图问题的算法（例如最短路径、最小生成树）。
- **动态规划和贪心算法**：了解解决问题的范式，涉及将复杂问题分解为更简单的子问题。
- **机器学习算法**：了解各种机器学习技术，包括监督学习和非监督学习，以及它们的应用。
- **算法设计**：能够设计和实现解决特定问题的算法，考虑效率、可扩展性和可维护性等因素。

### 网络编程知识（网络编程知识）

- **网络协议**：了解OSI模型和TCP/IP模型，以及每一层的协议（例如HTTP、FTP、TCP、UDP、ICMP）。
- **套接字编程**：了解如何创建和管理套接字进行网络通信，包括客户端-服务器架构和使用Berkeley套接字等API。
- **网络安全**：熟悉安全协议（例如SSL/TLS、VPN）和保护网络通信免受威胁的做法。
- **多线程和并发**：能够使用线程或异步编程同时处理多个网络连接。
- **分布式系统**：了解分布式计算、消息传递的概念，以及如何设计能够跨多台机器扩展的系统。
- **RESTful API设计**：了解设计和使用RESTful API的知识，包括HTTP方法、状态码和API版本控制。
- **网络工具和实用程序**：熟练使用Wireshark、curl和netcat等工具进行网络调试和分析。
- **负载均衡和代理服务器**：了解如何高效分配网络流量以及代理服务器在网络架构中的作用。


[**storage**](https://github.com/zhanhuipinggit/knowledge/storage)
- [ceph](./storage/ceph)
    - [**1. Ceph 简介与基本概念**](./storage/ceph/1.ceph简介与基本概念)
        - **[1.1 什么是 Ceph](./storage/ceph/1.ceph简介与基本概念/1.什么是ceph.md)**
        - **[1.2 Ceph 的应用场景](./storage/ceph/1.ceph简介与基本概念/2.Ceph的应用场景.md)**
        - **[1.3 Ceph 的特点](./storage/ceph/1.ceph简介与基本概念/3.Ceph的特点.md)**

    - [2. Ceph 架构概览](./storage/ceph/2.Ceph架构概览)
        - **[2.1 Ceph 集群的基本组成](./storage/ceph/2.Ceph架构概览/1.ceph集群的基本组成.md)**
        - **[2.2 数据存储结构](./storage/ceph/2.Ceph架构概览/2.数据存储结构.md)**
        - **[2.3 数据一致性与复制](./storage/ceph/2.Ceph架构概览/3.数据一致性与复制.md)**

    - [3. Ceph Monitor (MON)](./storage/ceph/3.CephMonitor(MON))
        - **[3.1 MON 的角色和功能](./storage/ceph/3.CephMonitor(MON)/1.MON的角色和功能.md)**
        - **[3.2 MON 的协议](./storage/ceph/3.CephMonitor(MON)/2.MON的协议.md)**
        - **[3.3 Quorum 与故障恢复](./storage/ceph/3.CephMonitor(MON)/3.Quorum与故障恢复.md)**

    - [4. Ceph OSD (Object Storage Daemon)](./storage/ceph/4.CephOSD(ObjectStorageDaemon))
        - **[4.1 OSD 的工作原理](./storage/ceph/4.CephOSD(ObjectStorageDaemon)/1.OSD的工作原理.md)**
        - **[4.2 数据的分布与副本](./storage/ceph/4.CephOSD(ObjectStorageDaemon)/2.数据的分布与副本.md)**
        - **[4.3 数据复制与恢复](./storage/ceph/4.CephOSD(ObjectStorageDaemon)/3.数据复制与恢复.md)**

    - [5. Ceph CRUSH 算法](./storage/ceph/5.CephCRUSH算法)
        - **[5.1 CRUSH 算法概述](./storage/ceph/5.CephCRUSH算法/1.CRUSH%20算法概述.md)**
        - **[5.2 数据分布与负载均衡](./storage/ceph/5.CephCRUSH算法/2.数据分布与负载均衡.md)**
        - **[5.3 CRUSH 规则和调整](./storage/ceph/5.CephCRUSH算法/3.CRUSH%20规则和调整.md)**

    - [6. Ceph MDS (Metadata Server)](./storage/ceph/6.CephMDS(MetadataServer))
        - **[6.1 MDS 的作用与功能](./storage/ceph/6.CephMDS(MetadataServer)/1.MDS的作用与功能.md)**
        - **[6.2 MDS 的工作流程](./storage/ceph/6.CephMDS(MetadataServer)/2.MDS的工作流程.md)**
        - **[6.3 MDS 集群与负载均衡](./storage/ceph/6.CephMDS(MetadataServer)/3.MDS集群与负载均衡.md)**

    - [7. Ceph 客户端与接口](./storage/ceph/7.Ceph客户端与接口)
        - **[7.1 Ceph 客户端架构](./storage/ceph/7.Ceph客户端与接口/1.Ceph%20客户端架构.md)**
        - **[7.2 对象存储（RADOS）](./storage/ceph/7.Ceph客户端与接口/2.对象存储(RADOS).md)**
        - **[7.3 块存储（RBD）](./storage/ceph/7.Ceph客户端与接口/3.块存储（RBD）.md)**
        - **[7.4 文件存储（CephFS](./storage/ceph/7.Ceph客户端与接口/4.文件存储（CephFS）.md)**

    - [8. Ceph 的数据一致性与恢复](./storage/ceph/8.Ceph的数据一致性与恢复)
        - **[8.1 CRUSH 与数据一致性](./storage/ceph/8.Ceph的数据一致性与恢复/1.CRUSH%20与数据一致性.md)**
        - **[8.2 数据恢复与自愈能力](./storage/ceph/8.Ceph的数据一致性与恢复/2.数据恢复与自愈能力.md)**
        - **[8.3 故障检测与修复机制](./storage/ceph/8.Ceph的数据一致性与恢复/3.故障检测与修复机制.md)**

    - [9. Ceph 性能优化与调优](./storage/ceph/9.Ceph性能优化与调优)
        - **[9.1 Ceph 的性能瓶颈](./storage/ceph/9.Ceph性能优化与调优/1.Ceph%20的性能瓶颈.md)**
        - **[9.2 调优策略](./storage/ceph/9.Ceph性能优化与调优/2.调优策略.md)**
        - **[9.3 监控与故障排除](./storage/ceph/9.Ceph性能优化与调优/3.监控与故障排除.md)**

    - [10. Ceph 的集群部署与维护](./storage/ceph/10.Ceph的集群部署与维护)
        - **[10.1 集群的部署与安装](./storage/ceph/10.Ceph的集群部署与维护/1.集群的部署与安装.md)**
        - **[10.2 集群扩展与升级](./storage/ceph/10.Ceph的集群部署与维护/2.集群扩展与升级.md)**
        - **[10.3 集群健康检查与问题解决](./storage/ceph/10.Ceph的集群部署与维护/3.集群健康检查与问题解决.md)**
- [clickhouse](./storage/clickhouse)
- [mysql](./storage/mysql)
- [redis](./storage/redis)
- [tidb](./storage/tidb)