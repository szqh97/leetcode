#include <iostream>
#include <cstdio>
using namespace std;
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};
class Solution {
    public:
        ListNode *detectCycle(ListNode *head) {
            auto p1 = head;
            auto p2 = head;
            auto p3 = head;

            if (head == NULL) {
                return NULL;
            }
            while (p1 != NULL and p2 != NULL) {
                p1 = p1->next;
                if (p1 == NULL) {
                    return NULL;
                }

                for (auto i = 0; i < 2; ++i) {
                    p2 = p2->next;
                    if (p2 == NULL) return NULL;
                }

                if (p1 == p2 ) {
                    while (p1 != p3) {
                        p1 = p1->next;
                        if (p1 == p2) {
                            p3 = p3->next;
                        }
                    }
                    return p1;
                }
            }
            return NULL;
        }
};

int main()
{
    auto l1 = new ListNode(1);
    auto l2 = new ListNode(2);
    Solution s;
    auto r = s.detectCycle(l1);
    printf("%x %x\n", l1, r);

    return 0;
}
