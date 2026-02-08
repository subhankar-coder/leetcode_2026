from typing import Optional

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

# class Solution:
#     def isBalanced(self, root: Optional[TreeNode]) -> bool:
#         if root == None:
#             return True

#         leftHeight,rightHeight = self.getHeight(root.left),self.getHeight(root.right)
#         return False if abs(rightHeight-leftHeight) > 1 else self.isBalanced(root.left) and self.isBalanced(root.right)

#     def getHeight(self,root:Optional[TreeNode]) -> int:
#         if root == None:
#             return -1
#         return max(self.getHeight(root.left),self.getHeight(root.right))+1


class Solution:
    def isBalanced(self, root: Optional[TreeNode]) -> bool:
        isPossible = True
        def isBalancedHelper(root:Optional[TreeNode]) -> int:
            nonlocal isPossible
            if root == None:
                return 0

            l,r = isBalancedHelper(root.left),isBalancedHelper(root.right)
            if abs(r-l) > 1:
                isPossible = False
            return max(l,r)+1
        isBalancedHelper(root)
        return isPossible



if __name__ == "__main__":
    root = TreeNode(3, TreeNode(9),TreeNode(20,TreeNode(15),TreeNode(7)))
    s = Solution()
    print(s.isBalanced(root))