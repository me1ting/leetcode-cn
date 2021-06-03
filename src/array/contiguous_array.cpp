#include <vector>
#include <unordered_map>
using std::unordered_map;
using std::vector;

class Solution
{
public:
    // https://leetcode-cn.com/problems/contiguous-array/
    // 前缀和+哈希表
    // 将0的值当作-1，问题转换为求和为0的最大长度，等价于两个前缀和的差为0，即两个前缀和相等
    // 使用哈希表记录前缀和n首次出现的位置，在一次遍历中记录最长前缀和
    int findMaxLength(vector<int> &nums)
    {
        int n = nums.size();
        int maxLen = 0;
        int preSum = 0;
        unordered_map<int, int> mp;

        mp[0] = -1;
        for (auto i = 0; i < n; i++)
        {
            auto val = nums[i];
            if (val == 0)
            {
                val = -1;
            }
            preSum += val;
            if (mp.count(preSum) > 0)
            {
                auto len = i - mp[preSum];
                if (len > maxLen)
                {
                    maxLen = len;
                }
            }
            else
            {
                mp[preSum] = i;
            }
        }
    }
};