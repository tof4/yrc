# Youkai Relay Chat

Even more experimental than Lanchat!

## Planned features

* Channels and private messages
* Chat history
* SSH authorization

## Not planned features

* File transfer

## API

API is designed to be both human and machine readable. It"s similar to SQL.

* Single query consists of arguments separated by spaces.
* First argument determines how rest of query will be treated.
* First argument have to be in lower case.
* Argument containing a space have to be surrounded by `" "` marks.
* If you want to send `"` in argument write it as `\"`.
* Query have to be ended with `\n`.
* Query have to be encoded in UTF-8.
* Query from client to server are commands.
* Query from server to client are events.

## Server

* Connection to YRC server is made via SSH.

## Definitions

* `id` - User"s nickname. Have to be unique.
* `timestamp` - Date and time in `YYYY-MM-DD|HH:mm:ss`

## Commands

| Command | Schema             | Description       | Example              |
| ------- | ------------------ | ----------------- | -------------------- |
| send    | `send "[content]"` | Send message.     | `sent "hello world"` |
| nick    | `nick "[id]"`      | Set nickname.     | `nick "test"`        |
| exit    | `exit`             | Close connection. | `exit`               |

## Events

| Event   | Schema                                           | Description            | Example                                                     |
| ------- | ------------------------------------------------ | ---------------------- | ----------------------------------------------------------- |
| message | `message from "[id]" at [timestamp] "[content]"` | New message received.  | `message from "test" at 1970-01-01\|00:00:00 "hello world"` |
| renamed | `renamed "[id]" to "[new id]"`                   | User changed nickname. | `renamed "default" to "new name"`                           |

## Database

Default server implemetation uses plain file system as database.

* usr
  * name
    * passwordhash
    * publickey
* chl
  * name
    * messages
    * subscribers