#include<vector>
#include<iostream>
#include <algorithm>
#include <iterator>

using namespace std;


class Solution {
public:
    vector<int> plusOne(vector<int>& digits) {
    	int buffer =1;
    	vector<int> result;
        for(auto it = digits.rbegin();it!=digits.rend();++it){
        	int current_digit = *it+buffer;
        	result.emplace(result.begin(),current_digit%10);
        	buffer = current_digit/10;
        }
        if(buffer!=0){
        	result.emplace(result.begin(),buffer);
        }
        return result;
    }
};

int main(){
	vector<int> v = {1,2,3};
	Solution sol;
	vector<int> res = sol.plusOne(v);
	copy(res.begin(),res.end(),ostream_iterator<int>(cout,","));
	return 0;
}