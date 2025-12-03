func merge(intervals [][]int) [][]int {
    if len(intervals) == 0 {
        return [][]int{}
    }

    // 1. 按区间起点从小到大排序
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] < intervals[j][1]
        }
        return intervals[i][0] < intervals[j][0]
    })

    // 2. 依次扫描并合并
    res := [][]int{intervals[0]} // 先把第一个区间放进去
    for i := 1; i < len(intervals); i++ {
        inter2 := res[len(res)-1] // 结果数组中最后一个区间

        // 如果当前区间与 inter2 有重叠：当前起点 <= inter2 的终点
        if intervals[i][0] <= inter2[1] {
            // 合并：更新 last 的右端点为两者较大值
            if intervals[i][1] > inter2[1] {
                inter2[1] = intervals[i][1]
                res[len(res)-1] = inter2
            }
        } else {
            // 没有重叠，直接放进结果数组
            res = append(res, intervals[i])
        }
    }

    return res
}
