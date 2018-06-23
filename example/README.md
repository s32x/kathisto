# kathisto-demo
Using kathisto is extremely simple. This example aims to demonstrate how kathisto works, and how you can use it to deploy pre-rendered web-apps easily.

## The problem
Below we have a basic `index.html` file that is the entrypoint for our application. It's fairly standard (in the case of Angular) and provides little context on the content to be indexed by search engines. All that exists is a `head` and a basic `body` containing one or many bundle files that are responsible for replacing the `<app-root>` content with actual website content.

Unfortunately, the below html is all that is generally seen by search-engine crawlers which obviously leaves a bot with very little information on how to index our pages. What we want a bot to see is a fully populated body that a search engine can index and rank based on our actual website content.
```
<!doctype html>
<html>

<head>
    <meta charset="utf-8">
    <title>kathisto Demo</title>
</head>

<body>
    <app-root></app-root>
    <script src="bundle.js"></script>
</body>

</html>
```
The javascript `bundle.js` - _Many bots/crawlers don't take this into consideration_
```
document.getElementsByTagName("app-root")[0].innerHTML = "This is your javascript rendered website content!.";
```
## The solution
Angular and React both offer different software/Node server-side rendering packages (ex Universal) that I'm sure work great. However because I was unable to get Universal working on my Angular4 applications, I decided to build a middleware in Go that rendered the content using PhantomJS, cache it on my server, and serve the cached pre-rendered pages from our Go/PhantomJS middleware.

What a typical bot/crawler will then see (specifically in the case of the above `/dist/` files) is the following...
```
<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8">
    <title>kathisto Demo</title>
</head>

<body>
    <app-root>This is your javascript rendered website content!.</app-root>
    <script src="bundle.js"></script>
</body>

</html>
```
This leaves the search-engine to make better decisions on where to index and rank your web pages and helps increase organic, search-engine traffic to your apps!
## How to use
The dockerfile is very simple...
```
FROM entrik/kathisto
ADD dist/ /dist/
```
As you can see all that needs to happen is you need to build a new image containing your `/dist/` folder.
### Build it
```
docker build --no-cache -t kathisto-demo .
```
### Run it
```
docker run --name=kathisto-demo -p 80:80 kathisto-demo
```
