# paste


## storage entities

See `entities.proto` for schema.

## storage keys

### stored in database

* `user:<user-id>` - user entity serialized with protobuf
* `user:<user-id>:paste:<bookmarked-at>` - paste ID that user with given `user-id` bookmarked at `bookmarked-at` unix nanosecond time.

* `paste:<paste-id>` - paste entity serialized with protobuf
* `paste:<expire-at>:expire` - paste ID that should expire and given `expire-at` unix nanosecond time

* `oauth:<oauth-id>:user` - user ID for given oauth key


### stored in cache

* `session:<session-key>` - user ID that is using session with given key
