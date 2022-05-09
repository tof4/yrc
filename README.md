# Youkai Relay Chat

Even more experimental than Lanchat!

## API

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

* `id` - User's username. Have to be unique.
* `timestamp` - UTC Unix time.

## Commands

| Command | Schema                       | Description                     | Example                   |
| ------- | ---------------------------- | ------------------------------- | ------------------------- |
| send    | `send [channel] "[content]"` | Send message.                   | `send main "hello world"` |
| read    | `read [channel] [count]`     | Get last messages from channel. | `read main 10`            |
| exit    | `exit`                       | Close connection.               | `exit`                    |

## Events

| Event   | Schema                                           | Description           | Example                                      |
| ------- | ------------------------------------------------ | --------------------- | -------------------------------------------- |
| message | `message [channel] [id] [timestamp] "[content]"` | New message received. | `message main test 1650036004 "hello world"` |

## Database

Default server implemetation uses plain file system as database.

* usr
  * name
    * passwordhash
    * publickey
* chl
  * name
    * chat
    * members

## Tools
 * serv - Host server
 * channeladd - Create channel
 * channeldel - Delete channel
 * useradd - Create user
 * userdel - Delete user
 * usermod - Add or remove user from channel