# `kathisto`

kathisto (greek for _render_) is a lightweight docker base image that performs full server-side rendering using Headless Chrome in order to serve pre-rendered pages of your Angular/React web-apps to search engine bots/crawlers.

This project was developed after I came to the realization that bots /crawlers were not indexing anything other than the head and un-rendered body of my Javascript based web-applications.

### Your Dockerfile
```sh
FROM sdwolfe32/kathisto:latest
ADD dist/ /dist/
```

### Deploy it!
```sh
docker build myapp
docker run --name myapp_container -p 80:80 myapp
```

### Example
If you need further clarification on how kathisto operates or perhaps a concrete example check out [the example dir](example/).

### Credit
 * [Headless Chrome](https://developers.google.com/web/updates/2017/04/headless-chrome) for Page rendering

The MIT License (MIT)
=====================

Copyright © 2018 Steven Wolfe

Permission is hereby granted, free of charge, to any person
obtaining a copy of this software and associated documentation
files (the “Software”), to deal in the Software without
restriction, including without limitation the rights to use,
copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the
Software is furnished to do so, subject to the following
conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.