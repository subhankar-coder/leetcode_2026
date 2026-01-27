import heapq
from typing import List
class Solution:
    def minCost(self, n: int, edges: List[List[int]]) -> int:
        graph = [[] for _ in range(n)]
        for u,v,w in edges:
            graph[u].append((v,w))
            graph[v].append((u,2*w))
        visited = [float("inf")]*n
        heap = [(0,0)]
        visited[0] = 0
        while heap:
            u,d = heapq.heappop(heap)
            if d > visited[u]:
                continue
            for v,w in graph[u]:
                if visited[v] > visited[u] + w:
                    visited[v] = visited[u] + w
                    heapq.heappush(heap,(v,visited[v]))
        return visited[-1] if visited[-1]!=float("inf") else -1




if __name__ == "__main__":
    sol = Solution()
    edges,n = [[2,3,25],[2,1,18],[3,1,2]],4
    print(sol.minCost(n,edges))
        