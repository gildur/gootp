gootp
=====

Implementation of RFC-2289 One-Time Password 

Usage
-----

Download and use from command line:
```sh
go get github.com/gildur/gootp

otp-md4 seq seed
otp-md5 seq seed
otp-sha1 seq seed
```

Or use in your code: 
```go
import (
	"github.com/gildur/gootp"
)
...
answer := gootp.OtpMd4(seq, seed, passphrase)
answer := gootp.OtpMd5(seq, seed, passphrase)
answer := gootp.OtpSha1(seq, seed, passphrase)
```

License
-------

<a href="https://raw.githubusercontent.com/gildur/gootp/master/LICENSE">MIT</a>
