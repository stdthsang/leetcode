class Solution {
public:
    int minSwapsCouples(vector<int>& row) {
        int size = row.size(), ret = 0;
        for (int i = 0; i < size; i += 2) {
            if (row[i+1] == (row[i] ^ 1))
                continue;
            ++ret;
            for (int j = i + 2; j < size; ++j) {
                if (row[j] == (row[i] ^ 1)) {
                    swap(row[j], row[i+1]);
                    break;
                }
            }
        }
        return ret;
    }
};
