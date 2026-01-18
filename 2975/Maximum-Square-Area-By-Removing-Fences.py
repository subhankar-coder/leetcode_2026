from typing import List

MOD = 10**9+7
class Solution:
    def maximizeSquareArea(self, m: int, n: int, hFences: List[int], vFences: List[int]) -> int:
        hFences.insert(0,1)
        hFences.append(m)

        vFences.insert(0,1)
        vFences.append(n)
        
        mp = {}
        vlen,hlen,maxArea = len(vFences),len(hFences),-1
        for i in range(0,hlen-1):
            for j in range(i+1,hlen):
                base = abs(hFences[j]-hFences[i])
                mp[base]=1

        for i in range(0,vlen-1):
            for j in range(i+1,vlen):
                base = abs(vFences[j]-vFences[i])
                if base in mp:
                    maxArea = max(maxArea,base**2)
        return maxArea%MODı



if __name__ == "__main__":
    m,n,hFences,vFences = 6,7,[2],[4]
    s = Solution()
    print(s.maximizeSquareArea(m,n,hFences,vFences))
    