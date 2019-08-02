class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        int size = nums.size();
        vector<node> nodes;
        for (int i = 0; i < size; ++i) {
            node temp(i, nums[i]);
            nodes.push_back(temp);
        }
        sort(nodes.begin(), nodes.end(), [](const node &a, const node &b) { return a.val < b.val; });
        vector<int> ret;
        for (int i = 0, j = size - 1; i < j; ) {
            if (nodes[i].val + nodes[j].val == target) {
                int sub1 = nodes[i].index < nodes[j].index ? nodes[i].index : nodes[j].index;
                int sub2 = nodes[i].index < nodes[j].index ? nodes[j].index : nodes[i].index;
                ret.push_back(sub1);
                ret.push_back(sub2);
                return ret;
            }
            if (nodes[i].val + nodes[j].val < target)
                ++i;
            else --j;
        }
        return ret;
    }
    vector<int> twoSum_1(vector<int>& nums, int target) {
        int size = nums.size();
        unordered_map<int, int> m;
        vector<int> ret;
        for (int i = 0; i < size; ++i) {
            if (m.count(target-nums[i])) {
                ret.push_back(m[target-nums[i]]);
                ret.push_back(i);
                return ret;
            }
            m[nums[i]] = i;
        }
        return ret;
    }
private:
    struct node {
        int index;
        int val;
        node(int i, int v): index(i), val(v){}
    };
};
