import heapq


class Solution:
    def maxSlidingWindow(self, nums, k):
        max_heap = []
        max_elem = []
        for i in range(k):
            heapq.heappush(max_heap, (-nums[i], i))
        max_elem.append(-max_heap[0][0])
        for i in range(k, len(nums)):
            heapq.heappush(max_heap, (-nums[i], i))
            while max_heap[0][1] <= i - k:
                heapq.heappop(max_heap)
            max_elem.append(-max_heap[0][0])
        return max_elem