# Giyn's LeetCode travel

## 数组

### 二分查找

```go
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
```

### 双指针

快指针慢指针或者左指针右指针。

## 字符串



## 动态规划

![](https://img-blog.csdnimg.cn/20210117171307407.png)

### 0-1 背包问题

二维 DP 数组：

```c++
// weight数组的大小 就是物品个数
for(int i = 1; i < weight.size(); i++) { // 遍历物品
    for(int j = 0; j <= bagweight; j++) { // 遍历背包容量
        if (j < weight[i]) dp[i][j] = dp[i - 1][j]; 
        else dp[i][j] = max(dp[i - 1][j], dp[i - 1][j - weight[i]] + value[i]);
    }
}
```

一维 DP 数组：

```c++
for(int i = 0; i < weight.size(); i++) { // 遍历物品
    for(int j = bagWeight; j >= weight[i]; j--) { // 遍历背包容量
        dp[j] = max(dp[j], dp[j - weight[i]] + value[i]);
    }
}
```

### 组合类问题公式

```c++
dp[j] += dp[j - nums[i]]
```

### 完全背包问题

```c++
for(int i = 0; i < weight.size(); i++) { // 遍历物品
    for(int j = weight[i]; j < bagWeight; j++) { // 遍历背包容量
        dp[j] = max(dp[j], dp[j - weight[i]] + value[i]);
    }
}
```

### 排列组合问题

如果求组合数就是外层for循环遍历**物品**，内层for遍历**背包**。

如果求排列数就是外层for遍历**背包**，内层for循环遍历**物品**。

### 多重背包问题

转换为0-1背包问题

```c++
void test_multi_pack() {
    vector<int> weight = {1, 3, 4};
    vector<int> value = {15, 20, 30};
    vector<int> nums = {2, 3, 2};
    int bagWeight = 10;
    for (int i = 0; i < nums.size(); i++) {
        while (nums[i] > 1) { // nums[i]保留到1，把其他物品都展开
            weight.push_back(weight[i]);
            value.push_back(value[i]);
            nums[i]--;
        }
    }

    vector<int> dp(bagWeight + 1, 0);
    for(int i = 0; i < weight.size(); i++) { // 遍历物品
        for(int j = bagWeight; j >= weight[i]; j--) { // 遍历背包容量
            dp[j] = max(dp[j], dp[j - weight[i]] + value[i]);
        }
        for (int j = 0; j <= bagWeight; j++) {
            cout << dp[j] << " ";
        }
        cout << endl;
    }
    cout << dp[bagWeight] << endl;

}
```

改变遍历个数

```c++
void test_multi_pack() {
    vector<int> weight = {1, 3, 4};
    vector<int> value = {15, 20, 30};
    vector<int> nums = {2, 3, 2};
    int bagWeight = 10;
    vector<int> dp(bagWeight + 1, 0);


    for(int i = 0; i < weight.size(); i++) { // 遍历物品
        for(int j = bagWeight; j >= weight[i]; j--) { // 遍历背包容量
            // 以上为01背包，然后加一个遍历个数
            for (int k = 1; k <= nums[i] && (j - k * weight[i]) >= 0; k++) { // 遍历个数
                dp[j] = max(dp[j], dp[j - k * weight[i]] + k * value[i]);
            }
        }
        // 打印一下dp数组
        for (int j = 0; j <= bagWeight; j++) {
            cout << dp[j] << " ";
        }
        cout << endl;
    }
    cout << dp[bagWeight] << endl;
}
```

### 背包问题总结

动态规划五部法：

1. 确定dp数组（dp table）以及下标的含义 
2. 确定递推公式 
3. dp数组如何初始化 
4. 确定遍历顺序 
5. 举例推导dp数组

背包递推公式

问能否能装满背包（或者最多装多少）：dp[j] = max(dp[j], dp[j - nums[i]] + nums[i])，对应题目如下：

- 动态规划：416.分割等和子集 
- 动态规划：1049.最后一块石头的重量 II

问装满背包有几种方法：dp[j] += dp[j - nums[i]]，对应题目如下：

- 动态规划：494.目标和 
- 动态规划：518.零钱兑换 II 
- 动态规划：377.组合总和Ⅳ 
- 动态规划：70.爬楼梯

问背包装满最大价值：dp[j] = max(dp[j], dp[j - weight[i]] + value[i])，对应题目如下：

- 动态规划：474.一和零

问装满背包所有物品的最小个数：dp[j] = min(dp[j - coins[i]] + 1, dp[j])，对应题目如下：

- 动态规划：322.零钱兑换 
- 动态规划：279.完全平方数

#### 遍历顺序

#### 0-1背包问题

- 二维dp数组0-1背包问题先遍历物品还是先遍历背包都是可以的，且第二层for循环是**从小到大遍历**。 
- 一维dp数组0-1背包问题只能先遍历物品再遍历背包容量，且第二层for循环是从大到小遍历。

#### 完全背包问题

纯完全背包的一维dp数组实现，先遍历物品还是先遍历背包都可以，且第二层for循环是从小到大遍历。

但是仅仅是纯完全背包的遍历顺序是这样的，题目稍有变化，两个for循环的先后顺序就不一样了。

- 如果求组合数就是外层for循环遍历物品，内层for遍历背包。 
- 如果求排列数就是外层for遍历背包，内层for循环遍历物品。
- 如果求最小数，那么两层for循环的先后顺序就无所谓了

求组合数：动态规划：518.零钱兑换II

求排列数：动态规划：377.组合总和 Ⅳ、动态规划：70.爬楼梯进阶版（完全背包）

求最小数：动态规划：322.零钱兑换、动态规划：279.完全平方数

Reference：https://www.programmercarl.com/
