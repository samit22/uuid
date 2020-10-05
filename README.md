# UUID

> Command line tool (CLI) to generate UUID or list of UUIDs.

> Required Go version 1.13+

### Installation
Run the following command
```bash
 go get github.com/samit22/uuid

 go install github.com/samit22/uuid

 uuid
 ```

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
