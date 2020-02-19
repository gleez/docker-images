#!/usr/bin/env node
var prerender = require("prerender");

const options = {
  chromeLocation: "/usr/bin/chromium-browser",
  pageDoneCheckInterval: process.env.PAGE_DONE_CHECK_INTERVAL || 500,
  pageLoadTimeout: process.env.PAGE_LOAD_TIMEOUT || 20000,
  waitAfterLastRequest: process.env.WAIT_AFTER_LAST_REQUEST || 500,
  chromeFlags: [
    "--disable-software-rasterizer",
    "--disable-dev-shm-usage",
    "--no-sandbox",
    "--headless",
    "--disable-gpu",
    "--remote-debugging-port=9222",
    "--hide-scrollbars"
  ]
};
console.log("Starting with options:", options);

// borrow from https://github.com/prerender/prerender/blob/master/server.js
var server = prerender(options);

server.use(prerender.sendPrerenderHeader());
//server.use(prerender.blockResources());
server.use(prerender.removeScriptTags());
server.use(prerender.httpHeaders());

server.start();
