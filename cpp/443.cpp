class Solution {
public:
    int compress(vector<char>& chars) {
        int cur = 0, size = chars.size();
        for (int i = 0, j = 0; j < size; i = j) {
            while (j < size && chars[i] == chars[j])
                ++j;
            chars[cur++] = chars[i];
            if (j - i == 1)
                continue;
            string count = to_string(j-i);
            for (auto c : count)
                chars[cur++] = c;
        }
        return cur;
    }
};
