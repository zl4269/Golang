//C++实现素数筛
#include<iostream>
#include<vector>
using namespace std;

//素数筛实现:从2开始，每次除以前面的元素，如果能整除，则说明不是素数，就置为-1
vector<int> Filter(vector<int> nums) {
    int size = nums.size();
    for (int i = 0; i < size; i++) {
        if (nums[i] != -1) {
            for (int j = i + 1; j < size; j++) {
                if (nums[j] % nums[i] == 0) {
                    nums[j] = -1;
                }
            }
        }
    }
    //将所有素数保留
    vector<int> result;
    for (int i = 0; i < size; i++) {
        if (nums[i] != -1) {
            result.push_back(nums[i]);
        }
    }
    return result;
}

int main() {
    int n;//从2开始的n个元素
    cin >> n;
    vector<int> nums;
    for (int i = 0; i < n; i++) {
        nums.push_back(i+2); //注意：一定要从2开始
    }
    vector<int> ans = Filter(nums);
    for (auto e : ans) {
        cout << e << " ";
    }
    return 0;
}