//
// Created by Girai Kaku on 2019-07-01.
//

#include "Node.h"
#include <unordered_map>

class Solution
{
public:
    Node *cloneGraph(Node *node)
    {
        if (node == NULL)
            return NULL;
        if (m.find(node) != m.end())
        {
            return m[node];
        }
        std::vector<Node *> nbs = {};
        Node *root = new Node(node->val, nbs);
        m[node] = root;
        for (const auto &n : node->neighbors)
        {
            root->neighbors.push_back(cloneGraph(n));
        }
        return root;
    }

private:
    std::unordered_map<Node *, Node *> m;
};