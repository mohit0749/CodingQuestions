# your code goes hereclass Node:
class Node:
    def __init__(self, value, next_node):
        self.value = value
        self.next = next_node

    def __hash__(self) -> int:
        return hash(self.value)


class LRU:
    def __init__(self, size):
        self.size = size
        self.current_node_cnt = 1
        self.__addr_hash = dict()
        self.head_ptr = None
        self.tail_ptr = None

    def get(self, node):
        return self.__addr_hash[hash(node)]

    def add(self, value):
        node = Node(value, None)
        if self.head_ptr is None:
            self.head_ptr = node
            self.__update_hash(node, node)
        if self.tail_ptr is None:
            self.tail_ptr = node
            self.__update_hash(node, node)
        else:
            if self.current_node_cnt < self.size:
                self.tail_ptr.next = node
                self.tail_ptr = self.tail_ptr.next
                self.current_node_cnt += 1
                self.__update_hash(node, node)
            else:
                prev_node = self.__check_if_node_exists(node)
                if prev_node is not None:
                    if self.tail_ptr == prev_node:
                        return
                    self.tail_ptr.next = node
                    self.tail_ptr = self.tail_ptr.next
                    self.__delete_lr_node(prev_node)
                    self.__update_hash(node, node)
                else:

                    temp_node = self.head_ptr
                    self.head_ptr = self.head_ptr.next
                    self.tail_ptr.next = node
                    self.tail_ptr = self.tail_ptr.next
                    self.__update_hash(node, node)
                    del self.__addr_hash[hash(temp_node)]
                    del (temp_node)

    def __delete_lr_node(self, node):
        if node.next is None:
            node = None
        else:
            self.__update_hash(node.next, node)
            temp_node = Node(node.value, node.next)
            node.value = node.next.value
            node.next = node.next.next
            del (temp_node)

    def __update_hash(self, node, new_addr):
        self.__addr_hash[hash(node.value)] = new_addr

    def __check_if_node_exists(self, node):
        if hash(node) in self.__addr_hash:
            return self.__addr_hash[hash(node)]
        return None

    def print_lru(self):
        print(self.__addr_hash)
        temp_head = self.head_ptr
        while temp_head is not None:
            print(temp_head.value)
            temp_head = temp_head.next


if __name__ == "__main__":
    lru = LRU(size=5)
    lru.add(1)
    lru.add(2)
    lru.add(3)
    lru.add(4)
    lru.add(5)
    lru.add(2)
    lru.add(2)
    lru.add(2)
    lru.add(5)
    lru.add(4)
    lru.add(3)
    lru.add(2)
    lru.add(1)
    lru.print_lru()