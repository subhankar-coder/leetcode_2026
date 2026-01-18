#include <iostream>
#include <vector>

using namespace std;
class Solution {
public:
    int repeatedNTimes(vector<int>& nums) {
        unordered_map<int,int> mp;
        for(int num: nums){
        	if (mp.find(num) != mp.end()){
        		return num;
        	}else{
        		mp[num]++;
        	}
        }

        return -1;
    }
};

int main(int argc, char const *argv[])
{
	Solution sol;
	vector<int> nums = {1,2,3,3};
	cout<<sol.repeatedNTimes(nums)<<endl;
	return 0;
}
