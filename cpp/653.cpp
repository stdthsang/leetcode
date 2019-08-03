class Solution {
public:
    bool findTarget(TreeNode* root, int k) {
        stack<TreeNode*> stk;
        unordered_map<int, bool> m;
        TreeNode *cur = root;
        while (cur != nullptr || !stk.empty()) {
            while (cur != nullptr) {
                stk.push(cur);
                cur = cur->left;
            }
            if (!stk.empty()) {
                cur = stk.top();
                stk.pop();
                if (m.count(k-cur->val)) 
                    return true;
                m[cur->val] = true;
                cur = cur->right;
            }
        }
        return false;
    }
};
