class Solution {
public:
    string countAndSay(int n) {
        if (n < 1)
            return "";
        map<int, string> m = {{1,"1"}, {2,"2"}, {3,"3"}, {4,"4"}, {5,"5"}, {6,"6"}, {7,"7"}, {8,"8"}, {9,"9"}};
        string ret = "1";
        while (n > 1) {
            --n;
            string cur;
            int size = ret.size();
            for (int i = 0, j = 0; j < size; ) {
                while (j < size && ret[j] == ret[i])
                    ++j;
                cur += m[j-i] + ret[i];
                i = j;
                if (j == size)
                    break;
            }
            ret = cur;
        }
        return ret;
    }
};
