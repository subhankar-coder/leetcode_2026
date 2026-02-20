import java.util.*;



class Solution{

  public static TreeNode balanceBST(TreeNode root) {

    ArrayList<Integer> list = new ArrayList<>();
    inOrderTravershal(root, list);
    
    return buildBalancedTree(list,0,list.size()-1);
  }

  private static TreeNode buildBalancedTree(List<Integer> val,int left, int right){
    if (left > right){
      return null;
    }
    if(left == right){
      return new TreeNode(val.get(left));
    }

    int mid = left+(right-left)/2;
    TreeNode leftNode = buildBalancedTree(val,left,mid-1);
    TreeNode rightNode = buildBalancedTree(val,mid+1,right);

    TreeNode current=new TreeNode(val.get(mid));
    current.left = leftNode;
    current.right = rightNode;

    return current;
  }
  public static void inOrderTravershal(TreeNode root,ArrayList<Integer> list){
      if(root == null){
      return ;
    }
    inOrderTravershal(root.left, list);
    list.add(root.val);
    inOrderTravershal(root.right, list);
  }
  public static void main(String[] args) {
      // Solution s = new Solution();
      TreeNode root = new TreeNode(1,null,new TreeNode(2,null,new TreeNode(3,null,new TreeNode(4,null,null))));
      System.out.println(balanceBST(root).val);
  }
  static class TreeNode{
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode(){}
    TreeNode(int val){
      this.val = val;
    }
    TreeNode(int val,TreeNode left,TreeNode right){
      this.val = val;
      this.left = left;
      this.right = right;
    }
  }
}
