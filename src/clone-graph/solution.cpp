//
// Created by Girai Kaku on 2019-07-01.
//

#include <unordered_map>

/*
// Definition for a Node.
#include <vector>

class Node
{
public:
    int val;
    std::vector<Node *> neighbors;

    Node() {}

    Node(int _val, std::vector<Node *> _neighbors)
    {
        val = _val;
        neighbors = _neighbors;
    }
};
*/

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
        std::vector<Node *> vec = {};
        Node *root = new Node(node->val, vec);
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