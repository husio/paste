# paste


## storage entities

See `entities.proto` for schema.

## storage keys

* `user:<user-id>:id` - user entity serialized with protobuf
* `user:<oauth-id>:oauth` - user ID for given oauth key
* `user:<user-id>:paste:<bookmarked-at>` - paste ID that user with given `user-id` bookmarked at `bookmarked-at` unix nanosecond time.

* `paste:<paste-id>:id` - paste entity serialized with protobuf
* `paste:<expire-at>:expire` - paste ID that should expire and given `expire-at` unix nanosecond time


