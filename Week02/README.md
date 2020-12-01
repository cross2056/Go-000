Q:

我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

A:

包住最底下来自DAO接入层的error, 可以增加调用者堆栈信息，底层的sql.ErrNoRows是个标准库errors.New出来的SentinelError，只有基本信息