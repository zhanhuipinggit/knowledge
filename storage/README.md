# 介绍
### 存储知识（数据存储知识）

#### <span style="color: red;">为了表述更清晰，所有的代码或者文本都严格按照书本格式来编撰的，工作量巨大，若是有所收益麻烦点个start </span>


- **数据存储基础**：了解数据存储的基础知识，包括不同类型的存储介质（例如硬盘驱动器HDD、固态硬盘SSD、NVMe）、存储架构（例如直接附加存储DAS、网络附加存储NAS、存储区域网络SAN）和存储协议（例如SCSI、SATA、NVMe）。
- **数据库管理系统（DBMS）**：了解数据库的工作原理，包括关系型数据库（SQL）和非关系型数据库NoSQL，它们的模式以及如何使用SQL或其他查询语言与它们交互。
- **数据持久化**：了解数据如何保存和检索，包括文件系统、对象存储和数据序列化格式（例如JSON、XML、BSON）的使用。
- **数据冗余和复制**：熟悉确保数据持久性的技术，例如RAID配置、数据镜像和跨不同地理位置的复制。
- **数据备份和恢复**：了解备份策略，包括全备份、增量备份和差异备份，以及在数据丢失情况下的恢复流程。
- **存储优化**：掌握优化存储性能和容量的技术，包括数据去重、压缩和分层存储策略。

###  Table of Contents

[**storage**](https://github.com/zhanhuipinggit/knowledge/storage)
- [ceph](./ceph)
  - [**1. Ceph 简介与基本概念**](./ceph/1.ceph简介与基本概念)
    - **[1.1 什么是 Ceph](./ceph/1.ceph简介与基本概念/1.什么是ceph.md)**
    - **[1.2 Ceph 的应用场景](./ceph/1.ceph简介与基本概念/2.Ceph的应用场景.md)**
    - **[1.3 Ceph 的特点](./ceph/1.ceph简介与基本概念/3.Ceph的特点.md)**

  - [2. Ceph 架构概览](./ceph/2.Ceph架构概览)
    - **[2.1 Ceph 集群的基本组成](./ceph/2.Ceph架构概览/1.ceph集群的基本组成.md)**
    - **[2.2 数据存储结构](./ceph/2.Ceph架构概览/2.数据存储结构.md)**
    - **[2.3 数据一致性与复制](./ceph/2.Ceph架构概览/3.数据一致性与复制.md)**

  - [3. Ceph Monitor (MON)](./ceph/3.CephMonitor(MON))
    - **[3.1 MON 的角色和功能](./ceph/3.CephMonitor(MON)/1.MON的角色和功能.md)**
    - **[3.2 MON 的协议](./ceph/3.CephMonitor(MON)/2.MON的协议.md)**
    - **[3.3 Quorum 与故障恢复](./ceph/3.CephMonitor(MON)/3.Quorum与故障恢复.md)**

  - [4. Ceph OSD (Object Storage Daemon)](./ceph/4.CephOSD(ObjectStorageDaemon))
    - **[4.1 OSD 的工作原理](./ceph/4.CephOSD(ObjectStorageDaemon)/1.OSD的工作原理.md)**
    - **[4.2 数据的分布与副本](./ceph/4.CephOSD(ObjectStorageDaemon)/2.数据的分布与副本.md)**
    - **[4.3 数据复制与恢复](./ceph/4.CephOSD(ObjectStorageDaemon)/3.数据复制与恢复.md)**

  - [5. Ceph CRUSH 算法](./ceph/5.CephCRUSH算法)
    - **[5.1 CRUSH 算法概述](./ceph/5.CephCRUSH算法/1.CRUSH%20算法概述.md)**
    - **[5.2 数据分布与负载均衡](./ceph/5.CephCRUSH算法/2.数据分布与负载均衡.md)**
    - **[5.3 CRUSH 规则和调整](./ceph/5.CephCRUSH算法/3.CRUSH%20规则和调整.md)**

  - [6. Ceph MDS (Metadata Server)](./ceph/6.CephMDS(Metadata Server))
    - **[6.1 MDS 的作用与功能](./ceph/6.CephMDS(Metadata%20Server)/1.MDS的作用与功能.md)**
    - **[6.2 MDS 的工作流程](./ceph/6.CephMDS(Metadata%20Server)/2.MDS的工作流程.md)**
    - **[6.3 MDS 集群与负载均衡](./ceph/6.CephMDS(Metadata%20Server)/3.MDS集群与负载均衡.md)**

  - [7. Ceph 客户端与接口](./ceph/7.Ceph客户端与接口)
    - **[7.1 Ceph 客户端架构](./ceph/7.Ceph客户端与接口/1.Ceph%20客户端架构.md)**
    - **[7.2 对象存储（RADOS）](./ceph/7.Ceph客户端与接口/2.对象存储(RADOS).md)**
    - **[7.3 块存储（RBD）](./ceph/7.Ceph客户端与接口/3.块存储（RBD）.md)**
    - **[7.4 文件存储（CephFS](./ceph/7.Ceph客户端与接口/4.文件存储（CephFS）.md)**

  - [8. Ceph 的数据一致性与恢复](./ceph/8.Ceph的数据一致性与恢复)
    - **[8.1 CRUSH 与数据一致性](./ceph/8.Ceph的数据一致性与恢复/1.CRUSH%20与数据一致性.md)**
    - **[8.2 数据恢复与自愈能力](./ceph/8.Ceph的数据一致性与恢复/2.数据恢复与自愈能力.md)**
    - **[8.3 故障检测与修复机制](./ceph/8.Ceph的数据一致性与恢复/3.故障检测与修复机制.md)**

  - [9. Ceph 性能优化与调优](./ceph/9.Ceph性能优化与调优)
    - **[9.1 Ceph 的性能瓶颈](./ceph/9.Ceph性能优化与调优/1.Ceph%20的性能瓶颈.md)**
    - **[9.2 调优策略](./ceph/9.Ceph性能优化与调优/2.调优策略.md)**
    - **[9.3 监控与故障排除](./ceph/9.Ceph性能优化与调优/3.监控与故障排除.md)**

  - [10. Ceph 的集群部署与维护](./ceph/10.Ceph的集群部署与维护)
    - **[10.1 集群的部署与安装](./ceph/10.Ceph的集群部署与维护/1.集群的部署与安装.md)**
    - **[10.2 集群扩展与升级](./ceph/10.Ceph的集群部署与维护/2.集群扩展与升级.md)**
    - **[10.3 集群健康检查与问题解决](./ceph/10.Ceph的集群部署与维护/3.集群健康检查与问题解决.md)**

- [clickhouse](./clickhouse)
- [mysql](./mysql)
- [redis](./redis)
- [tidb](./tidb)
