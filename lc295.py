import heapq


class MedianFinder:

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.min_heap = []
        self.max_heap = []

    def addNum(self, num: int) -> None:
        heapq.heappush(self.max_heap, -num)
        elem = -heapq.heappop(self.max_heap)
        heapq.heappush(self.min_heap, elem)
        n = len(self.min_heap)
        m = len(self.max_heap)
        if m < n:
            heapq.heappush(self.max_heap, -heapq.heappop(self.min_heap))

    def findMedian(self) -> float:
        n = len(self.min_heap)
        m = len(self.max_heap)
        if m == n:
            elem1 = -heapq.nsmallest(1, self.max_heap)[0]
            elem2 = heapq.nsmallest(1, self.min_heap)[0]
            return (elem1 + elem2) * 0.5
        else:
            return -heapq.nsmallest(1, self.max_heap)[0]


if __name__ == "__main__":
    obj = MedianFinder()
    obj.addNum(-1)
    print(obj.findMedian())
    obj.addNum(-2)
    print(obj.findMedian())
    obj.addNum(-3)
    print(obj.findMedian())
    obj.addNum(-4)
    print(obj.findMedian())
    obj.addNum(-5)
    print(obj.findMedian())
    obj.addNum(1)
    print(obj.findMedian())
