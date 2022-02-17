# `kathisto`

kathisto (greek for _render_) is a lightweight docker base image that performs full server-side rendering using Headless Chrome in order to serve pre-rendered pages of your Angular/React web-apps to search engine bots/crawlers.

This project was developed after I came to the realization that bots /crawlers were not indexing anything other than the head and un-rendered body of my Javascript based web-applications.

### Your Dockerfile
```sh
FROM s32x/kathisto:latest
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

The BSD 3-clause License
========================

Copyright (c) 2022, s32x. All rights reserved.

Redistribution and use in source and binary forms, with or without modification,
are permitted provided that the following conditions are met:

 - Redistributions of source code must retain the above copyright notice,
   this list of conditions and the following disclaimer.

 - Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

 - Neither the name of kathisto nor the names of its contributors may
   be used to endorse or promote products derived from this software without
   specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
