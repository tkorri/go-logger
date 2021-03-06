# go-logger

Go library for message logging

Reads sql files from a directory and imports them to the database. Files that
have been already imported are skipped.

This has currently been tested only with PostgreSQL but should work with other
databases.

## Import

    import "github.com/tkorri/go-logger"

## Usage

```go
import (
    "github.com/tkorri/go-logger"
)

logger.Init()
logger.SetLogLevel("INFO")

logger.V("This message is not shown")
logger.I("This message is shown")
```

You can also configure logging output with SetLogFileLocation and SetTimeFormat

```go

logger.SetLogFileLocation("/home/tkorri/logs/golang.log")
logger.SetTimeFormat("2006-01-02 15:04:05")

```

## Documentation

Documentation is available at
[godoc.org](http://godoc.org/github.com/tkorri/go-logger).


## License

Copyright (c) 2015 Taneli Korri

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
