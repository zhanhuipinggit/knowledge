# Ceph 
 **Ceph 是一个开源的分布式存储系统，设计用于提供高性能、高可扩展性和高可靠性的存储解决方案**

## [**1. Ceph 简介与基本概念**](./1.ceph简介与基本概念)
- **[1.1 什么是 Ceph](./1.ceph简介与基本概念/1.什么是ceph.md)**：定义 Ceph，介绍它作为一个分布式存储系统的基本目标和特点。
- **[1.2 Ceph 的应用场景](./1.ceph简介与基本概念/2.Ceph的应用场景.md)**：讲解 Ceph 在大规模数据存储、云计算和容器环境中的使用场景。
- **[1.3 Ceph 的特点](./1.ceph简介与基本概念/3.Ceph的特点.md)**：分布式、可扩展、高可用性、自愈能力、无单点故障。

## [2. Ceph 架构概览](./2.Ceph架构概览)
- **[2.1 Ceph 集群的基本组成](./2.Ceph架构概览/1.ceph集群的基本组成.md)**：介绍 Ceph 集群的核心组件，包括 **Monitor (MON)**、**OSD (Object Storage Daemon)**、**MDS (Metadata Server)** 和 **Client**。
- **[2.2 数据存储结构](./2.Ceph架构概览/2.数据存储结构.md)**：讲解 Ceph 的存储机制，如何使用 CRUSH 算法来进行数据分布。
- **[2.3 数据一致性与复制](./2.Ceph架构概览/3.数据一致性与复制.md)**：介绍 Ceph 如何保证数据的一致性、可靠性及如何进行副本管理。

## [3. Ceph Monitor (MON)](./3.CephMonitor(MON))
- **[3.1 MON 的角色和功能](./3.CephMonitor(MON)/1.MON的角色和功能.md)**：介绍 Ceph Monitor 的作用，如何帮助维护集群的状态和健康。
- **[3.2 MON 的协议](./3.CephMonitor(MON)/2.MON的协议.md)**：讲解 MON 如何进行集群协调，包括集群的元数据和状态管理。
- **[3.3 Quorum 与故障恢复](./3.CephMonitor(MON)/3.Quorum与故障恢复.md)**：阐述 MON 的选举机制及其如何确保集群的高可用性。

## [4. Ceph OSD (Object Storage Daemon)](./4.CephOSD(Object Storage Daemon))
- **[4.1 OSD 的工作原理](./4.CephOSD(Object%20Storage%20Daemon)/1.OSD的工作原理.md)**：介绍 OSD 组件如何处理数据存储和读取，如何与 MON 协作来管理数据的持久化。
- **[4.2 数据的分布与副本](./4.CephOSD(Object%20Storage%20Daemon)/2.数据的分布与副本.md)**：详细讲解数据的分布机制，如何利用 CRUSH 算法计算数据存储的位置，并通过副本来提高可靠性。
- **[4.3 数据复制与恢复](./4.CephOSD(Object%20Storage%20Daemon)/3.数据复制与恢复.md)**：深入剖析 OSD 的副本管理与数据恢复机制。

## [5. Ceph CRUSH 算法](./5.CephCRUSH算法)
- **[5.1 CRUSH 算法概述](./5.CephCRUSH算法/1.CRUSH%20算法概述.md)**：介绍 CRUSH 算法的基本概念和原理，如何根据集群的拓扑结构进行数据的分布。
- **[5.2 数据分布与负载均衡](./5.CephCRUSH算法/2.数据分布与负载均衡.md)**：深入讲解如何利用 CRUSH 算法将数据均衡地分配到集群中的 OSD 节点上。
- **[5.3 CRUSH 规则和调整](./5.CephCRUSH算法/3.CRUSH%20规则和调整.md)**：讲解如何通过调整 CRUSH 规则来优化集群性能和存储策略。

## [6. Ceph MDS (Metadata Server)](./6.CephMDS(Metadata Server))
- **[6.1 MDS 的作用与功能](./6.CephMDS(Metadata%20Server)/1.MDS的作用与功能.md)**：介绍 MDS 在 Ceph 文件系统中的作用，尤其是在 CephFS 中，它如何管理文件系统的元数据。
- **[6.2 MDS 的工作流程](./6.CephMDS(Metadata%20Server)/2.MDS的工作流程.md)**：讲解 MDS 如何与 OSD 配合，处理客户端的文件操作请求，以及如何进行元数据的缓存和同步。
- **[6.3 MDS 集群与负载均衡](./6.CephMDS(Metadata%20Server)/3.MDS集群与负载均衡.md)**：解释 MDS 如何通过多副本和负载均衡来保证高效的元数据管理。

## [7. Ceph 客户端与接口](./7.Ceph客户端与接口)
- **[7.1 Ceph 客户端架构](./7.Ceph客户端与接口/1.Ceph%20客户端架构.md)**：介绍 Ceph 客户端的不同接口，如 RADOS、RBD、CephFS 和 RGW。
- **[7.2 对象存储（RADOS）](./7.Ceph客户端与接口/2.对象存储(RADOS).md)**：讲解 Ceph 如何通过 RADOS（Reliable Autonomic Distributed Object Store）提供底层的对象存储服务。
- **[7.3 块存储（RBD）](./7.Ceph客户端与接口/3.块存储（RBD）.md)**：介绍 Ceph 的块存储服务，如何通过 RBD 提供虚拟机磁盘、块设备等存储。
- **[7.4 文件存储（CephFS](./7.Ceph客户端与接口/4.文件存储（CephFS）.md)**：讲解 CephFS 作为分布式文件系统的工作原理，如何支持 POSIX 标准。

## [8. Ceph 的数据一致性与恢复](./8.Ceph的数据一致性与恢复)
- **[8.1 CRUSH 与数据一致性](./8.Ceph的数据一致性与恢复/1.CRUSH%20与数据一致性.md)**：解释 Ceph 如何使用 CRUSH 算法以及 OSD 副本来保证数据一致性。
- **[8.2 数据恢复与自愈能力](./8.Ceph的数据一致性与恢复/2.数据恢复与自愈能力.md)**：深入分析 Ceph 如何在 OSD 故障后自动恢复数据，确保集群的高可用性。
- **[8.3 故障检测与修复机制](./8.Ceph的数据一致性与恢复/3.故障检测与修复机制.md)**：讲解 Ceph 如何检测和修复集群中的故障，保障数据的完整性和一致性。

## [9. Ceph 性能优化与调优](./9.Ceph性能优化与调优)
- **[9.1 Ceph 的性能瓶颈](./9.Ceph性能优化与调优/1.Ceph%20的性能瓶颈.md)**：分析 Ceph 在不同应用场景中的性能瓶颈，例如 I/O 性能、网络延迟、磁盘吞吐量等。
- **[9.2 调优策略](./9.Ceph性能优化与调优/2.调优策略.md)**：讲解如何调整 Ceph 集群的配置和硬件，以优化性能，例如调整 CRUSH 规则、优化 OSD、调整网络配置等。
- **[9.3 监控与故障排除](./9.Ceph性能优化与调优/3.监控与故障排除.md)**：介绍如何使用 Ceph 的监控工具（如 Ceph Dashboard、Ceph CLI、Ceph metrics）来监控集群状态、调优性能以及故障排除。

## [10. Ceph 的集群部署与维护](./10.Ceph的集群部署与维护)
- **[10.1 集群的部署与安装](./10.Ceph的集群部署与维护/1.集群的部署与安装.md)**：介绍 Ceph 集群的安装步骤，使用工具（如 ceph-deploy、cephadm）部署集群。
- **[10.2 集群扩展与升级](./10.Ceph的集群部署与维护/2.集群扩展与升级.md)**：讲解如何扩展 Ceph 集群，添加 OSD 或 MON 节点，如何在不中断服务的情况下进行升级和维护。
- **[10.3 集群健康检查与问题解决](./10.Ceph的集群部署与维护/3.集群健康检查与问题解决.md)**：阐述如何检测集群的健康状态、识别潜在问题，并使用 Ceph 工具进行排错。
