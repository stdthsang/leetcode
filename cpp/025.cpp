class Solution {
public:
    ListNode* reverseKGroup(ListNode* head, int k) {
        ListNode *dummy = new ListNode(0);
        ListNode *tail = dummy;
        while (head != nullptr) {
            ListNode *kth = findKth(head, k);
            if (kth == nullptr) {
                tail->next = head;
                return dummy->next;
            }
            ListNode *temp = head;
            ListNode *next = kth->next;
            kth->next = nullptr;
            for (int i = 0; i < k; ++i) {
                ListNode *tmp = head->next;
                head->next = tail->next;
                tail->next = head;
                head = tmp;
            }
            tail = temp;
            head = next;
        }
        return dummy->next;
    }
    
    ListNode* findKth(ListNode* head, int k) {
        for (int i = 1; i < k && head != nullptr; ++i)
            head = head->next;
        return head;
    }
};
