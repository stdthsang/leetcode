/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        if (head == nullptr)
            return head;
        ListNode *cur = head;
        for (int i = 1; i < n; ++i)
            cur = cur->next;
        ListNode *dummy = new ListNode(0);
        dummy->next = head;
        ListNode *pre = dummy;
        while (cur->next != nullptr) {
            pre = pre->next;
            cur = cur->next;
        }
        pre->next = pre->next->next;
        return dummy->next;
    }
};
