import heapq


class MinStack:

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.stk = []
        self.h = []

    def push(self, x: int) -> None:
        self.stk.append(x)
        heapq.heappush(self.h, x)

    def pop(self) -> None:
        ele = self.stk[-1]
        self.stk.pop()
        self.h.remove(ele)

    def top(self) -> int:
        return self.stk[-1]

    def getMin(self) -> int:
        return heapq.nsmallest(1, self.h)


if __name__ == "__main__":
    obj = MinStack()
    obj.push(4)
    obj.push(3)
    obj.push(5)
    obj.push(6)
    obj.push(7)
    obj.pop()
    param_3 = obj.top()
    param_4 = obj.getMin()
# Your MinStack object will be instantiated and called as such:
