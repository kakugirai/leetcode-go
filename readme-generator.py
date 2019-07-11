#!/usr/bin/env python
# -*- coding: utf-8 -*-

import os
import requests
import json

# [![CircleCI](https://circleci.com/gh/kakugirai/leetcode-go.svg?style=svg)](https://circleci.com/gh/kakugirai/leetcode-go)

head = """# leetcode-go

[![Go Report Card](https://goreportcard.com/badge/github.com/kakugirai/leetcode-go)](https://goreportcard.com/report/github.com/kakugirai/leetcode-go)
[![Build Status](https://travis-ci.org/kakugirai/leetcode-go.svg?branch=master)](https://travis-ci.org/kakugirai/leetcode-go)

My LeetCode solutions primarily written in Go. Don't expect too much on those questions solved in other random languages.

## Wanna start your LeetCode journey?

Feel free to copy the cookies from your browser, write them into a `cookies.json` file,

```
{
    "__cfduid": "",
    "_ga": "",
    "LEETCODE_SESSION": "",
    "_gid": "",
    "__stripe_mid": "",
    "c_a_u": ""
}
```

and copy the `readme-generator.py` to your directory. Here you ≡Go!

## My progress

"""

stats_head = """
| Solved | Easy | Medium | Hard |
| :----: | :--: | :----: | :--: |
"""

questions_head = """

## Solutions

| Question ID | Question | Difficulty | Solution |
| :---------: | -------- | :--------: | :------: |
"""

if __name__ == "__main__":
    with open('cookies.json') as cookies_file:
        cookies = json.load(cookies_file)

    r = requests.get(
        "https://leetcode.com/api/problems/algorithms/", cookies=cookies)
    parsed_json = json.loads(r.content)

    stats_table = "|{}|{}|{}|{}|\n".format(
        parsed_json["num_solved"], parsed_json["ac_easy"], parsed_json["ac_medium"], parsed_json["ac_hard"])
    questions_table = ""
    for question in parsed_json["stat_status_pairs"]:
        if question["status"] == "ac":
            level = ["Easy", "Medium", "Hard"]
            question_id = question["stat"]["question_id"]
            question__title = question["stat"]["question__title"]
            question_level = level[question["difficulty"]["level"]-1]
            question__title_slug = question["stat"]["question__title_slug"]
            leetcode_link = "https://leetcode.com/problems/"
            done = os.listdir(os.path.join('src', question__title_slug))
            answers_path_arr = []
            for answer in done:
                lang = answer.split('.')[-1]
                answer_path = ""
                if lang == 'go':
                    answer_path = "[Go](src/{}/solution.go)".format(
                        question__title_slug)
                if lang == 'java':
                    answer_path = "[Java](src/{}/solution.java)".format(
                        question__title_slug)
                if lang == 'py':
                    answer_path = "[Python](src/{}/solution.py)".format(
                        question__title_slug)
                if lang == 'cpp':
                    answer_path = "[C++](src/{}/solution.cpp)".format(
                        question__title_slug)
                answers_path_arr.append(answer_path)
            answers_path = ", ".join(answers_path_arr)

            line = "|{}|[{}]({}{})|{}|{}|\n".format(question_id, question__title,
                                                    leetcode_link, question__title_slug, question_level, answers_path)
            questions_table = line + questions_table

    with open('README.md', 'w') as writer:
        writer.write(head + stats_head + stats_table +
                     questions_head + questions_table)
