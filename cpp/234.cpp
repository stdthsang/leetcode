class Solution {
public:
    bool isPalindrome(ListNode* head) {
        if (head == nullptr || head->next == nullptr)
            return true;
        vector<int> v;
        ListNode *cur = head;
        while (cur != nullptr) {
            v.push_back(cur->val);
            cur = cur->next;
        }
        for (int i = 0, j = v.size() - 1; i < j; ++i, --j)
            if (v[i] != v[j])
                return false;
        return true;
    }
};

