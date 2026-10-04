from collections import deque


class ResilientQueue:
    def __init__(self):
        self._items = deque()

    def put(self, item):
        self._items.append(item)

    def get(self):
        return self._items.popleft()

    def __len__(self):
        return len(self._items)
