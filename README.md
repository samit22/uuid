# UUID

> Command line tool (CLI) to generate UUID or list of UUIDs.

> Required Go version 1.17+

### Installation
Run the following command
```bash
 go install github.com/samit22/uuid@latest

 uuid
 ```
If you get error something like
```
# golang.org/x/sys/unix
../../../.gvm/pkgsets/go1.19/global/pkg/mod/golang.org/x/sys@v0.0.0-20190624142023-c5567b49c5d0/unix/zsyscall_darwin_amd64.go:28:3: //go:linkname must refer to declared function or variable
```
Run: `go get -u golang.org/x/sys`
This is due to the older version of sys
### Available Commands:
 - uuid
    - Default command
 - uuid v4
    - uuid v4
        - Generates single uuid.
    -  uuid v4 <i>10</i>
        - Generates list of 10 UUIDs. 10 here is an argument, can be between 1-100.
 - uuid check
    - uuid <b> uuid_1 </b>
        - Checks if the uuid is a valid one.
    - uuid <b> uuid_1 </b> &nbsp; <b>uuid_2 </b>
        - Validates each uuid format.
        - You can send as many argument as you want separated by a space.
 - uuid empty
    - Generates an empty uuid.
    - Can be used as uuid nil.
