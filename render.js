var system = require('system');
var args = system.args;
var page = require('webpage').create();
page.settings.userAgent = args[2];

// Return an error if there's javascript errors on render
page.onError = function(msg, trace) {
    console.log(msg);
    phantom.exit(1);
};

// Dump the page to the console
page.open(args[1], function() {
    console.log(page.content);
    phantom.exit(0);
});