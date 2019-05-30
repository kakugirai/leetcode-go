#!/usr/bin/env python
# -*- coding: utf-8 -*-

import requests
import json

head = """# leetcode-go

[![Go Report](https://goreportcard.com/badge/github.com/kakugirai/leetcode-go)](https://goreportcard.com/report/github.com/kakugirai/leetcode-go)

My LeetCode solutions written in Go.
"""

stats_head = """
| Solved | Easy | Medium | Hard |
| :----: | :--: | :----: | :--: |
"""

questions_head = """
| Question ID | Question | Difficulty | Solution |
| :---------: | -------- | :--------: | :------: |
"""

foot = """### Questions solved in C++ or Java or Python (don't support Go)

| Question ID | Question | Difficulty | Solution |
| :---------: | -------- | :--------: | :------: |
|278|[First Bad Version](https://leetcode.com/problems/first-bad-version)|Easy|[Solution](src/first-bad-version)|
|374|[Guess Number Higher or Lower](https://leetcode.com/problems/guess-number-higher-or-lower)|Easy|[Solution](src/guess-number-higher-or-lower)|
"""

if __name__ == "__main__":
    with open('cookies.json') as cookies_file:
        cookies = json.load(cookies_file)

    r = requests.get(
        "https://leetcode.com/api/problems/algorithms/", cookies=cookies)
    parsed_json = json.loads(r.content)

    stats_table = "|{}|{}|{}|{}|\n".format(parsed_json["num_solved"], parsed_json["ac_easy"], parsed_json["ac_medium"], parsed_json["ac_hard"])
    questions_table = ""
    for question in parsed_json["stat_status_pairs"]:
        if question["status"] == "ac":
            level = ["Easy", "Medium", "Hard"]
            question_id = question["stat"]["question_id"]
            question__title = question["stat"]["question__title"]
            question_level = level[question["difficulty"]["level"]-1]
            question__title_slug = question["stat"]["question__title_slug"]
            leetcode_link = "https://leetcode.com/problems/"
            line = "|{}|[{}]({}{})|{}|[Solution](src/{})|\n".format(question_id, question__title,
                                                                    leetcode_link, question__title_slug, question_level, question__title_slug)
            questions_table = line + questions_table

    with open('README.md', 'w') as writer:
        writer.write(head + stats_head + stats_table + questions_head + questions_table + foot)
