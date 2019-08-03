class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        vector<int> n1, n2;
        ListNode *p1 = l1, *p2 = l2;
        while (p1 != nullptr) {
            n1.push_back(p1->val);
            p1 = p1->next;
        }
        while (p2 != nullptr) {
            n2.push_back(p2->val);
            p2 = p2->next;
        }
        int s1 = n1.size() - 1, s2 = n2.size() - 1, carry = 0;
        ListNode *dummy = new ListNode(0);
        while (s1 >= 0 && s2 >= 0) {
            int temp = carry + n1[s1--] + n2[s2--];
            ListNode *l = new ListNode(temp%10);
            carry =  temp / 10;
            l->next = dummy->next;
            dummy->next = l;
        }
        while (s1 >= 0) {
            int temp = carry + n1[s1--];
            ListNode *l = new ListNode(temp%10);
            carry =  temp / 10;
            l->next = dummy->next;
            dummy->next = l;
        }
        while (s2 >= 0) {
            int temp = carry + n2[s2--];
            ListNode *l = new ListNode(temp%10);
            carry =  temp / 10;
            l->next = dummy->next;
            dummy->next = l;
        }
        if (carry) {
            ListNode *l = new ListNode(carry);
            l->next = dummy->next;
            dummy->next = l;
        }
        return dummy->next;
    }
};
