class Solution {
public:
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
        return dummy->next;
    }
};
