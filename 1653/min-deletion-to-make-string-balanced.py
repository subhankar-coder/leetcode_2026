class Solution:
    def minimumDeletions(self, s: str) -> int:
        n = len(s)
        prefix,suffix,pre_sum,suf_sum = [0]*n,[0]*n,1 if s[0]=='b' else 0,1 if s[-1]=='a' else 0 
        for i,j in zip(range(1,n,1),range(n-2,-1,-1)):
            prefix[i],suffix[j] = pre_sum,suf_sum
            if s[i]=='b': pre_sum+=1
            if s[j]=='a': suf_sum+=1

        sm = float("inf")
        for i in range(n):
            sm = min(sm,prefix[i]+suffix[i])
        return sm



if __name__ == "__main__":
    s = Solution()
    print(s.minimumDeletions("aababbab"))