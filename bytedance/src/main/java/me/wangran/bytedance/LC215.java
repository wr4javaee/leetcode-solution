package me.wangran.bytedance;

import java.util.PriorityQueue;

/**
 * @author Wang Ran
 * @date 2019-08-02
 * @desc
 */
public class LC215 {
    public int findKthLargest(int[] nums, int k) {
        PriorityQueue<Integer> minHeap = new PriorityQueue<Integer>(k);
        for (int i = 0; i < nums.length; i++) {

            if (minHeap.size() < k) {
                minHeap.add(nums[i]);
                continue;
            }

            if (minHeap.peek() < nums[i]) {
                minHeap.poll();
                minHeap.add(nums[i]);
                continue;
            }

            continue;
        }

        return minHeap.peek();

    }
}
