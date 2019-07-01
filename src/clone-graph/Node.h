//
// Created by Girai Kaku on 2019-07-01.
//

#include <vector>

// Definition for a Node.
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
