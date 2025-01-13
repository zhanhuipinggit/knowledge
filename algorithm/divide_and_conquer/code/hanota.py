def hanota(n, source, target,auxiliary):
    """
        汉诺塔问题的递归实现
        :param n: 盘子的数量
        :param source: 源柱子名称
        :param target: 目标柱子名称
        :param auxiliary: 辅助柱子名称
        """

    if n == 1:
        print(f"将盘子 1 从 {source} 移动到 {target}")
    else:
        hanota(n-1,source, auxiliary, target)
        print(f"将盘子 {n} 从 {source} 移动到 {target}")
        hanota(n-1,auxiliary,target,source)

