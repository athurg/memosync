# Memos Discovery

发现身边的 [Memos](https://usememos.com) ， 并订阅同步到自己的 *Memos* 实例。

## 实现原理

- 将感兴趣的 *Memos* 实例视作自己 *Memos* 实例的一个用户。
- 周期性的检测新增的 *Memo* 并同步为对应用户发送的 *Memo*。
- 通过自己的 *Memos* 实例的 *Explore* 查看。


## 使用方法

下载对应平台的二进制文件，带参数运行即可。主要参数:

- `-h 自己Memos的访问地址`
- `-k 自己Memos的管理员用户的OpenID`
- `-targets 逗号分割的感兴趣的多个Memos实例的访问地址`

可选参数

- `-i 10m`: 同步周期，缺省为10m，也就是十分钟同步一次。

## 安装方法（Systemd）

- 下载或编译二进制文件，并放到 `/usr/local/bin/discovery_memo` 目录下
- 下载 `conf/memo_discovery.conf` ，放到 `/etc/systemd/system` 或 `/usr/local/lib/systemd/system/` 目录下
- 执行 `sudo systemctl daemon-reload && sudo systemctl enable memos_discovery` 安装服务
- 执行 `sudo systemctl start memos_discovery` 启动服务

如需查看日志，可执行 `journalctl -f -u memos_discovery`

## 重要提醒

由于会定期性的访问感兴趣的 *Memos* 实例，因此使用前请先征求对方的同意。

尤其避免被对方误判为爬虫、恶意访问者。

## TODO

- 等 *Memos* API 支持指定某一条 *Memo* `ID`作为起止点后，优化拉取的逻辑。当前的实现是直接先拉指定数量条，然后按照发送时间过滤
- 尽量争取把该功能合并到 *Memos* 项目中。

## 问题反馈

请在 *Memos* 的 [Telegram用户群](https://t.me/+-_tNF1k70UU4ZTc9) 联系 @athurg 。
