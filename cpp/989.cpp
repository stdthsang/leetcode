class Solution {
public:
    vector<int> addToArrayForm(vector<int>& A, int K) {
        vector<int> B;
        for (; K != 0; K /= 10) {
            B.push_back(K%10);
        }
        for (int i = 0, size = B.size(); i < size / 2; i++) 
            swap(B[i], B[size-i-1]);
        int i = A.size() - 1, j = B.size() - 1, carry = 0;
        vector<int> ret;
        while (i >= 0 && j >= 0) {
            int temp = carry + A[i--] + B[j--];
            ret.push_back(temp%10);
            carry = temp / 10;
        }
        while (i >= 0) {
            int temp = carry + A[i--];
            ret.push_back(temp%10);
            carry = temp / 10;
        }
        while (j >= 0) {
            int temp = carry + B[j--];
            ret.push_back(temp%10);
            carry = temp / 10;
        }
        if (carry)
            ret.push_back(carry);
        for (int i = 0, size = ret.size(); i < size / 2; i++) 
            swap(ret[i], ret[size-i-1]);
        return ret;
    }
};
