'use strict';

const puppeteer = require('puppeteer');
const fs = require('fs');

const questions = fs.readFileSync("questions.txt", {
  encoding: "utf8"
}).split('\n');

(async() => {
  for (let i = 0; i < questions.length; i++) {
    const browser = await puppeteer.launch();
    const page = await browser.newPage();
    const link = "https://leetcode.com/problems/" + questions[i]
    await page.goto(link, {
      waitUntil: 'networkidle0'
    });
    const element = await page.$("#question-title");
    const title = await page.evaluate(element => element.textContent, element);
    console.log("|[%s](%s)|[Solution](src/%s)|", title, link, questions[i]);
    await browser.close();
  }
})();