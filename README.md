# Giyn's LeetCode travel

## 二分查找

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

## 双指针

快指针慢指针或者左指针右指针。

## 二叉树

### 前序遍历

迭代：

```go
func preorderTraversal(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return
}
```

递归：

```go
func preorderTraversal(root *TreeNode) (ans []int) {
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}
```

### 中序遍历

迭代：

```go
func inorderTraversal(root *TreeNode) (ans []int) {
	var stack []*TreeNode
	cur := root
	for cur != nil || len(stack) != 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, cur.Val)
			cur = cur.Right
		}
	}
	return
}
```

递归：

```go
func inorderTraversal(root *TreeNode) (ans []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		ans = append(ans, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}
```

### 后序遍历

迭代：

```go
func postorderTraversal(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	Reverse(ans, 0, len(ans)-1)
	return
}
```

递归:

```go
func postorderTraversal(root *TreeNode) (ans []int) {
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		ans = append(ans, node.Val)
	}
	postorder(root)
	return
}
```

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

## 排序算法

### 冒泡排序

思路：两两交换，大的放在后面，第一次排序后最大值已在数组末尾。因为两两交换，需要 `n-1` 趟排序（比如 10 个数，需要 9 趟排序）。

代码实现要点：**两个 for 循环，外层循环控制排序的趟数，内层循环控制比较的次数**。**每趟过后，比较的次数都应该要减 1**。

![849589-20171015223238449-2146169197](https://gitee.com/giyn/image-storage/raw/master/images/849589-20171015223238449-2146169197.gif)

```go
func BubbleSort(nums []int) {
	// 排序的趟数
	for i := 0; i < len(nums)-1; i++ {
		var isChange bool
		// 当前需要排序的趟数
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isChange = true
			}
		}
		if !isChange {
			break
		}
	}
}
```

### 选择排序

思路：找到最小（大）元素，存放到排序序列的起始位置，然后，再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。以此类推，直到所有元素均排序完毕，因此需要 `n-1` 趟排序。

代码实现要点：**两个 for 循环，外层循环控制排序的趟数，内层循环找到当前趟数的最大值，随后与当前趟数组最后的一位元素交换**。

![849589-20171015224719590-1433219824](https://gitee.com/giyn/image-storage/raw/master/images/849589-20171015224719590-1433219824.gif)

```go
func SelectionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		var minIndex = i
		// 扫描到底
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}
```

### 插入排序

思路：将一个元素插入到已有序的数组中，在初始时未知是否存在有序的数据，因此将元素**第一个元素看成是有序的**。与有序的数组进行比较，**比它大则直接放入，比它小则移动数组元素的位置，找到个合适的位置插入**。当只有一个数时，则不需要插入了，因此需要 `n-1` 趟排序。

代码实现：一个 for 循环内嵌一个 while 循环实现，外层 for 循环控制需要排序的趟数，while 循环找到合适的插入位置（并且插入的位置不能小于 0）。

![849589-20171015225645277-1151100000](https://gitee.com/giyn/image-storage/raw/master/images/849589-20171015225645277-1151100000.gif)

```go
func InsertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		prev := i - 1
		for prev >= 0 && nums[prev] > cur {
			nums[prev+1] = nums[prev] // 前移
			prev--
		}
		nums[prev+1] = cur // 插入
	}
}
```

### 快速排序

思路：在数组中找一个节点，比它小的放在节点左边，比它大的放在节点右边。一趟下来，比节点小的在左边，比节点大的在右边，不断执行这个操作。

步骤：

快速排序使用分治法来把一个串（list）分为两个子串（sub-lists）。

- 从数组中挑出一个元素，称为 “基准”（pivot）；
- 重新排序数组，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数组的中间位置。这个称为分区（partition）操作；
- 递归地（recursive）把小于基准值元素的子数组和大于基准值元素的子数组排序。

![849589-20171015230936371-1413523412](https://gitee.com/giyn/image-storage/raw/master/images/849589-20171015230936371-1413523412.gif)

```go
func QuickSort(nums []int, left, right int) {
	i, j := left, right
	pivot := nums[(left+right)/2] // 基准
	for i <= j {
		// 找到比基准大的数
		for nums[i] < pivot {
			i++
		}
		// 找到比基准小的数
		for nums[j] > pivot {
			j--
		}
		// 此时已经找到了比基准小(右)和大(左)的数,交换它们
		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	if left < j {
		QuickSort(nums, left, j)
	}
	if right > i {
		QuickSort(nums, i, right)
	}
}
```

### 归并排序

思路：将两个已排好序的数组合并成一个有序的数组。**将元素分隔开来，看成是有序的数组，进行比较合并**。不断拆分和合并，直到只有一个元素。

步骤：

1. 把长度为 n 的输入序列分成两个长度为 n/2 的子序列；
2. 对这两个子序列分别采用归并排序；
3. 将两个排序好的子序列合并成一个最终的排序序列。

代码实现：在第一趟排序时实质是两个元素（看成是两个已有序的数组）来进行合并，不断执行这样的操作，最终数组有序，拆分左边，右边，合并…

归并排序是建立在归并操作上的一种有效的排序算法。该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。将已有序的子序列合并，得到完全有序的序列；即先使每个子序列有序，再使子序列段间有序。若将两个有序表合并成一个有序表，称为二路归并。

<img src="https://gitee.com/giyn/image-storage/raw/master/images/image-20220321210353687.png" alt="image-20220321210353687" style="zoom:80%;" />

![849589-20171015230557043-37375010](https://gitee.com/giyn/image-storage/raw/master/images/849589-20171015230557043-37375010.gif)

```go
func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	left, right := nums[:mid], nums[mid:]
	return Merge(MergeSort(left), MergeSort(right))
}

func Merge(left, right []int) (ans []int) {
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			ans = append(ans, left[0])
			left = left[1:]
		} else {
			ans = append(ans, right[0])
			right = right[1:]
		}
	}
	for len(left) != 0 {
		ans = append(ans, left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		ans = append(ans, right[0])
		right = right[1:]
	}
	return
}
```

### 堆排序

思路：堆排序使用到了完全二叉树的一个特性，根节点比左孩子和右孩子都要大（小），完成**一次建堆**的操作实质上是比较根节点和左孩子、右孩子的大小，大（小）的交换到根节点上，**直至最大（小）的节点在树顶**。随后与数组最后一位元素进行交换。

代码实现：只要左子树或右子树大于当前根节点，则替换。替换后会导致下面的子树发生了变化，因此同样需要进行比较，直至各个节点实现父>子这么一个条件。

1. 将初始待排序关键字序列 (R1,R2….Rn) 构建成大（小）顶堆，此堆为初始的无序区；
2. 将堆顶元素 R[1] 与最后一个元素 R[n] 交换，此时得到新的无序区 (R1,R2,……Rn-1) 和新的有序区 (Rn)，且满足 R[1,2…n-1]<=R[n]；
3. 由于交换后新的堆顶 R[1] 可能违反堆的性质，因此需要对当前无序区 (R1,R2,……Rn-1) 调整为新堆，然后再次将 R[1] 与无序区最后一个元素交换，得到新的无序区 (R1,R2….Rn-2) 和新的有序区 (Rn-1,Rn)。不断重复此过程直到有序区的元素个数为 n-1，则整个排序过程完成。

![img](https://images2017.cnblogs.com/blog/849589/201710/849589-20171015231308699-356134237.gif)

```go
func HeapSort(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		lastLen := length - i
		Heapify(nums, lastLen)
		if i < length {
			nums[0], nums[lastLen-1] = nums[lastLen-1], nums[0]
		}
	}
}

func Heapify(nums []int, length int) {
	if length <= 1 {
		return
	}
	depth := length/2 - 1 // 二叉树深度
	for i := depth; i >= 0; i-- {
		top := i
		left := 2*i + 1
		right := 2*i + 2
		if left <= length-1 && nums[left] > nums[top] {
			top = left
		}
		if right <= length-1 && nums[right] > nums[top] {
			top = right
		}
		if top != i {
			nums[i], nums[top] = nums[top], nums[i]
		}
	}
}
```

### 希尔排序

思路：希尔排序实质上就是插入排序的增强版，希尔排序将数组分隔成 n 组来进行插入排序，**直至该数组宏观上有序，**最后再进行插入排序时就不用移动那么多次位置了。

先将整个待排序的记录序列分割成为若干子序列分别进行直接插入排序。

步骤：

1. 选择一个增量序列 t1，t2，…，tk，其中 ti>tj，tk=1；
2. 按增量序列个数 k，对序列进行 k 趟排序；
3. 每趟排序，根据对应的增量 ti，将待排序列分割成若干长度为 m 的子序列，分别对各子表进行直接插入排序。仅增量因子为 1 时，整个序列作为一个表来处理，表长度即为整个序列的长度。

代码思路：希尔增量一般是 `gap = gap / 2`，只是比普通版插入排序多了一个 for 循环而已。

![img](https://images2018.cnblogs.com/blog/849589/201803/849589-20180331170017421-364506073.gif)

```go
func ShellSort(nums []int) {
	for gap := len(nums) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(nums); i++ {
			j := i
			cur := nums[i]
			for j-gap >= 0 && nums[j-gap] > cur {
				nums[j] = nums[j-gap]
				j -= gap
			}
			nums[j] = cur
		}
	}
}
```

### 计数排序

计数排序不是基于比较的排序算法，其核心在于将输入的数据值转化为键存储在额外开辟的数组空间中。作为一种线性时间复杂度的排序，计数排序**要求输入的数据必须是有确定范围的整数**。

步骤：

1. 找出待排序的数组中最大和最小的元素；
2. 统计数组中每个值为 i 的元素出现的次数，存入数组 C 的第 i 项；
3. 对所有的计数累加（从 C 中的第一个元素开始，每一项和前一项相加）；
4. 反向填充目标数组：将每个元素i放在新数组的第 C(i) 项，每放一个元素就将 C(i) 减去 1。

![849589-20171015231740840-6968181](https://gitee.com/giyn/image-storage/raw/master/images/849589-20171015231740840-6968181.gif)

```go
func CountingSort(nums []int) {
	maxVal := Max(nums)
	minVal := Min(nums)
	bucket := make([]int, maxVal-minVal+1)
	for _, num := range nums {
		bucket[num-minVal]++
	}
	sortedIndex := 0
	for i := 0; i < maxVal-minVal+1; i++ {
		for bucket[i] > 0 {
			nums[sortedIndex] = i + minVal
			sortedIndex++
			bucket[i]--
		}
	}
}
```

### 基数排序

基数排序是按照低位先排序，然后收集；再按照高位排序，然后再收集；依次类推，直到最高位。有时候有些属性是有优先级顺序的，先按低优先级排序，再按高优先级排序。最后的次序就是高优先级高的在前，高优先级相同的低优先级高的在前。

步骤：

1. 取得数组中的最大数，并取得位数；
2. nums 为原始数组，从最低位开始取每个位组成 radix 数组；
3. 对 radix 进行计数排序（利用计数排序适用于小范围数的特点）；

![849589-20171015232453668-1397662527](https://gitee.com/giyn/image-storage/raw/master/images/849589-20171015232453668-1397662527.gif)

```go
func RadixSort(nums []int) {
	maxVal := Max(nums)
	maxBit := 0 // 最大值的位数
	for maxVal > 0 {
		maxBit++
		maxVal /= 10
	}
	tmp := make([]int, len(nums))
	count := new([10]int)
	for i, radix := 0, 1; i < maxBit; i, radix = i+1, radix*10 {
		for j := 0; j < 10; j++ {
			count[j] = 0
		}
		for j := 0; j < len(nums); j++ {
			count[(nums[j]/radix)%10]++
		}
		for j := 1; j < 10; j++ {
			count[j] += count[j-1]
		}
		for j := len(nums) - 1; j >= 0; j-- {
			tmp[count[(nums[j]/radix)%10]-1] = nums[j]
			count[(nums[j]/radix)%10]--
		}
		for j := 0; j < len(nums); j++ {
			nums[j] = tmp[j]
		}
	}
}
```

### 桶排序

桶排序是计数排序的升级版。它利用了函数的映射关系，高效与否的关键就在于这个映射函数的确定。桶排序（Bucket sort）的工作的原理：假设输入数据服从均匀分布，将数据分到有限数量的桶里，每个桶再分别排序（有可能再使用别的排序算法或是以递归方式继续使用桶排序进行排）。

步骤：

1. 设置一个定量的数组当作空桶；
2. 遍历输入数据，并且把数据一个一个放到对应的桶里去；
3. 对每个不是空的桶进行排序；
4. 从不是空的桶里把排好序的数据拼接起来。 

![image-20220322121117832](https://gitee.com/giyn/image-storage/raw/master/images/image-20220322121117832.png)

```go
func BucketSort(nums []int) {
	bucketNum := len(nums) / 2 // 桶数
	maxVal := Max(nums)
	buckets := make([][]int, bucketNum)
	// 分配入桶
	index := 0
	for i := 0; i < bucketNum; i++ {
		// 分配桶 index = value * (n-1) /k
		index = nums[i] * (bucketNum - 1) / maxVal
		buckets[index] = append(buckets[index], nums[i])
	}
	// 桶内排序
	tmpPos := 0
	for i := 0; i < bucketNum; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			sort.Ints(buckets[i])
			copy(nums[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}
}
```

- 稳定：如果 a 原本在 b 前面，而 a=b，排序之后 a 仍然在 b 的前面。
- 不稳定：如果 a 原本在 b 的前面，而 a=b，排序之后 a 可能会出现在 b 的后面。

| 排序算法 | 平均时间复杂度 | 最好     | 最坏      | 空间复杂度 | 稳定性 |
| :------- | :------------- | :------- | :-------- | :--------- | :----- |
| 冒泡排序 | O(n^2)         | O(n)     | O(n^2)    | O(1)       | 稳定   |
| 快速排序 | O(nlogn)       | O(nlogn) | O(n^2)    | O(logn)    | 不稳定 |
| 插入排序 | O(n^2)         | O(n)     | O(n^2)    | O(1)       | 稳定   |
| 希尔排序 | O(n^1.3)       | O(n)     | O(nlog2n) | O(1)       | 不稳定 |
| 选择排序 | O(n^2)         | O(n^2)   | O(n^2)    | O(1)       | 不稳定 |
| 堆排序   | O(nlogn)       | O(nlogn) | O(nlogn)  | O(1)       | 不稳定 |
| 归并排序 | O(nlogn)       | O(nlogn) | O(nlogn)  | O(n)       | 稳定   |
| 桶排序   | O(n+k)         | O(n+k)   | O(n+k)    | O(n+k)     | 稳定   |
| 计数排序 | O(n+k)         | O(n+k)   | O(n+k)    | O(k)       | 稳定   |
| 基数排序 | O(n*k)         | O(n*k)   | O(n*k)    | O(n+k)     | 稳定   |
