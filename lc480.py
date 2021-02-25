import heapq
from collections import Counter


class Solution:
    def medianSlidingWindow(self, nums, k):
        self.min_heap = []
        self.max_heap = []
        self.min_map = Counter()
        self.max_map = Counter()
        self.min_heap_cnt = 0
        self.max_heap_cnt = 0
        medians = []
        for i in range(0, k):
            self._add(nums[i])
        medians.append(self.get_median())

        for i in range(k, len(nums)):
            self._remove(nums[i - k])
            self._add(nums[i])
            medians.append(self.get_median())
        return medians

    def get_median(self):
        if self.min_heap_cnt == self.max_heap_cnt:
            return (self.__get_min_from_min_heap() + self.__get_max_from_max_heap()) * 0.5
        elif self.min_heap_cnt > self.max_heap_cnt:
            return self.__get_min_from_min_heap()
        else:
            return self.__get_max_from_max_heap()

    def __get_min_from_min_heap(self):
        elem = self.min_heap[0]
        while self.min_map[elem] == 0:
            heapq.heappop(self.min_heap)
            elem = self.min_heap[0]
        return elem

    def __get_max_from_max_heap(self):
        elem = -self.max_heap[0]
        while self.max_map[elem] == 0:
            heapq.heappop(self.max_heap)
            elem = -self.max_heap[0]
        return elem

    def _add(self, elem):
        if self.min_heap_cnt > 0 and elem < self.min_heap[0]:
            heapq.heappush(self.max_heap, -elem)
            self.max_map[elem] += 1
            self.max_heap_cnt += 1
        else:
            heapq.heappush(self.min_heap, elem)
            self.min_map[elem] += 1
            self.min_heap_cnt += 1
        self._balance_size()

    def _balance_size(self):
        while abs(self.max_heap_cnt - self.min_heap_cnt) > 1:
            if self.max_heap_cnt > self.min_heap_cnt:
                self.__get_max_from_max_heap()
                elem = -heapq.heappop(self.max_heap)
                heapq.heappush(self.min_heap, elem)
                self.min_map[elem] += 1
                self.max_map[elem] -= 1
                self.max_heap_cnt -= 1
                self.min_heap_cnt += 1
            elif self.min_heap_cnt > self.max_heap_cnt:
                self.__get_min_from_min_heap()
                elem = heapq.heappop(self.min_heap)
                heapq.heappush(self.max_heap, -elem)
                self.min_map[elem] -= 1
                self.max_map[elem] += 1
                self.min_heap_cnt -= 1
                self.max_heap_cnt += 1

    def _remove(self, elem):
        if self.max_map[elem] > 0:
            self.max_map[elem] -= 1
            self.max_heap_cnt -= 1
            if self.max_map[elem] == 0:
                del self.max_map[elem]
        elif self.min_map[elem] > 0:
            self.min_map[elem] -= 1
            self.min_heap_cnt -= 1
            if self.min_map[elem] == 0:
                del self.min_map[elem]
        self._balance_size()


if __name__ == "__main__":
    # [1, 2, 3], 4, 2, 3, 1, 4, 2 = 2
    # 1, [2, 3, 4,] 2, 3, 1, 4, 2 = 3
    # 1, 2, [3, 4, 2], 3, 1, 4, 2 = 3
    # 1, 2, 3, [4, 2, 3], 1, 4, 2 = 3
    # 1, 2, 3, 4, [2, 3, 1], 4, 2 = 1
    # 1, 2, 3, 4, 2, [3, 1, 4], 2 = 3
    # 1, 2, 3, 4, 2, 3, [1, 4, 2] = 2

    nums = [1, 2, 3, 4, 2, 3, 1, 4, 2]
    k = 3
    nums = [1,3,-1,-3,5,3,6,7]
    k = 3
    obj = Solution()
    print(obj.medianSlidingWindow(nums, k))
