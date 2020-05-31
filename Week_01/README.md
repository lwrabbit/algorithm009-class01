# 学习笔记
## 思维导图
### 目的
给自己的零碎算法知识来一次runtime.GC()，根据学习阶段循序渐进的来完善自己的算法思维导图
### 状态
```
数据结构与算法
-> 存储结构
   -> 线性存储结构
      -> 基础线性存储结构           
         -> 数组
            -> insert/delete:O(n)
            -> lookup:O(1)仅指下标随机访问
         -> 链表
            -> insert/delete:O(1)
            -> lookup:O(n)
            -> 跳表(基于有序链表，目的是利用二分查找思想达到O(logn)的lookup)
               -> insert/delete/lookup:O(logn)
      -> 操作受限线性存储结构
         -> 栈stack
            -> FILO
            -> insert/delete:O(1)
            -> lookup:O(n) 栈内元素无序
         -> 队列queue
            -> FIFO
            -> insert/delete:O(1)
            -> lookup:O(n) 栈内元素无序
         -> 双端队列Deque
            -> 两端都可以出队和入队
         -> 优先级队列
            -> 基于堆排序优先出队           
```
## 技能训练
### 数组
#### 暴力枚举N个元素组合
```
for i := 0; i < len(array)-N; i++{
    for j := i+1; j < len(array)-N+1; j++{
        ...
        // 假设n是最内一层循环，上一层loop变量是m
        for n := m+1; n < len(array); n++{
            //[i,j,...,n]组合枚举值计算
        }
    } 
}
```
#### 左右指针向中间收敛
```
//一般先进行排序
sort.Ints(array)
for i, j := 0, len(array)-1; i < j; {
    if a[i] < a[j] {
        //todo something
        i++
    } else {
        //todo something
        j--
    }
    //todo something
}
```
### 链表
#### 指针遍历
可以想象成数组下标的i,k暴露枚举
```
//反转链表 cur->当前node；pre->当前node.prev
cur.next,pre,cur = pre,cur,cur.next
```
#### 快慢指针
主要解决算法场景
- 链表是否有环
- 寻找链表中点
- 寻找链表倒数第k个元素

