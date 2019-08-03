/**
 *  * Definition for singly-linked list.
 *   * struct ListNode {
 *    *     int val;
 *     *     ListNode *next;
 *      *     ListNode(int x) : val(x), next(NULL) {}
 *       * };
 *        */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode *dummy = new ListNode(0), *tail = dummy;
        int carry = 0;
        while (l1 && l2) {
            int temp = carry + l1->val + l2->val;
            tail->next = new ListNode(temp%10);
            carry = temp / 10;
            l1 = l1->next;
            l2 = l2->next;
            tail = tail->next;
        }
        while (l1) {
            int temp = carry + l1->val;
            tail->next = new ListNode(temp%10);
            carry = temp / 10;
            l1 = l1->next;
            tail = tail->next;
        }
        while (l2) {
            int temp = carry + l2->val;
            tail->next = new ListNode(temp%10);
            carry = temp / 10;
            l2 = l2->next;
            tail = tail->next;
        }
        if (carry)
            tail->next = new ListNode(carry);
        return dummy->next;
    }
};
