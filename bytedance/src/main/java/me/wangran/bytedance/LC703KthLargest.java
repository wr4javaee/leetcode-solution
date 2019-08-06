package me.wangran.bytedance;

import java.util.PriorityQueue;

/**
 * @author Wang Ran
 * @date 2019-08-02
 * @desc
 */
class LC703KthLargest {

    private PriorityQueue<Integer> minHeap;
    private int k;


    public LC703KthLargest(int k, int[] nums) {
        minHeap =  new PriorityQueue<Integer>(k);
        this.k = k;
        for (int i = 0; i < nums.length; i++) {
            add(nums[i]);
        }
    }

    public int add(int val) {
        int heapSize = minHeap.size();
        // 若堆还有空间, 则加入元素
        if (heapSize < k) {
            minHeap.add(val);
        } else {
            // 获得堆顶
            int heapTop = minHeap.peek();
            if (heapTop < val) {
                // 来了一个更大的数, 把堆顶删除
                minHeap.poll();
                minHeap.add(val);
            }
        }

        return minHeap.peek();
    }

}


