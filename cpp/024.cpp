class Solution {
public:
    ListNode* swapPairs(ListNode* head) {
        ListNode *dummy = new ListNode(0);
        ListNode *tail = dummy;
        while (head != nullptr && head->next != nullptr) {
            ListNode *temp = head->next->next;
            head->next->next = nullptr;

            tail->next = head->next;
            tail->next->next = head;
            head->next = nullptr;

            tail = tail->next->next;

            head = temp;
        }
        if (head != nullptr)
            tail->next = head;
        return dummy->next;
    }
};
