class Solution {
public:
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        if (lists.empty())
            return nullptr;
        while (lists.size() > 1) {
            int size = lists.size();
            vector<ListNode*> temp;
            for (int i = 0; i < size - 1; i += 2) {
                temp.push_back(mergeTwoLists(lists[i], lists[i+1]));
            }
            if (size % 2 != 0)
                temp.push_back(lists[size-1]);
            lists = temp;
        }
        return lists[0]; 
    }
    ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
        if (l1 == nullptr)
            return l2;
        if (l2 == nullptr)
            return l1;
        ListNode *dummy = new ListNode(0);
        ListNode *tail = dummy;
        while (l1 != nullptr && l2 != nullptr) {
            if (l1->val <= l2->val) {
                ListNode *temp = l1->next;
                l1->next = nullptr;
                tail->next = l1;
                l1 = temp;
            } else {
                ListNode *temp = l2->next;
                l2->next = nullptr;
                tail->next = l2;
                l2 = temp;
            }
            tail = tail->next;
        }
        if (l1 != nullptr)
            tail->next = l1;
        if (l2 != nullptr)
            tail->next = l2;
        tail = dummy->next;
        delete dummy;
        return tail;
    }
};
