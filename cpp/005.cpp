class Solution {
public:
    string longestPalindrome(string s) {
        int size = s.size();
        string ret;
        for (int i = 0; i < size; ++i) {
            string temp = palindrome(s, i, i , size);
            if (temp.size() > ret.size())
                ret = temp;
            temp = palindrome(s, i, i+1, size);
            if (temp.size() > ret.size())
                ret = temp;
        }
        return ret;
    }

    string palindrome(string s, int i, int j, int size) {
        while (i >= 0 && j < size && s[i] == s[j]) {
            --i;
            ++j;
        }
        return s.substr(i+1,j-i-1);
    }

};
