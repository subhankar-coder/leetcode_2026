"""
    1 2 3
    _____
    |_|_|   1
    |_|_|   2
    |_|_|   3
"""



from typing import List


class Solution:
    def maximizeSquareHoleArea(self, n: int, m: int, hBars: List[int], vBars: List[int]) -> int:
        start_v,end_v,start_h,end_h = 1,m+2,1,n+2
        hBars.sort()
        vBars.sort()

        maxV,maxH,vlen,hlen,count = 2,2,len(vBars),len(hBars),1

        for i in range(1,vlen):
            if vBars[i]-vBars[i-1] ==1:
                count+=1
            else:
                maxV = max(maxV,count+1)
                count=1
        maxV = max(maxV,count+1)
        print(maxV)
        count=1
        for i in range(1,hlen):
            if hBars[i]-hBars[i-1] ==1:
                count+=1
            else:
                maxH = max(maxH,count+1)
                count = 1
        maxH = max(maxH,count+1)

        return min(maxH,maxV)**2





if __name__ == "__main__":
    s = Solution()
    n,m,hBars,vBars = 4,40,[5,3,2,4],[36,41,6,34,33]
    print(s.maximizeSquareHoleArea(n,m,hBars,vBars))
