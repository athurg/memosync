# Memos Discovery

通过 *Explore* 发现身边的 [Memos](https://usememos.com)

## 实现方案

- 将感兴趣的 *Memos* 实例作为自己 *Memos* 实例的一个用户。
- 周期性的检测新增的 *Memo* 并同步为给该用户
- 通过自己的 *Memos* 实例的 *Explore* 查看感兴趣的 *Memos*。

## 安装（Systemd）

- 下载或编译二进制文件`discovery_memo`，并放到 `/usr/local/bin/` 下
- 参考 [conf/memos_discovery.service](conf/memos_discovery.service) 创建自己的Systemd配置
- 将Systemd配置文件放在 `/etc/systemd/system` 或 `/usr/local/lib/systemd/system/` 下
- 执行 `sudo systemctl daemon-reload && sudo systemctl enable memos_discovery` 安装服务
- 执行 `sudo systemctl start memos_discovery` 启动服务

如需查看日志，可执行 `journalctl -f -u memos_discovery`

## 使用方法

- 首先在 *Memos* 中，以你要订阅的 *Memos* 的URL为用户名创建用户，比如: `https://usememos.com`
- 然后运行本服务即可

命令行参数说明：

- `-h http://my.usememos.com`: 必须，自己 *Memos* 实例的地址
- `-k xxx`: 必须，自己 *Memos* 实例的管理员OpenID
- `-i 30m`: 可选，检查的周期，默认为十分钟

## 重要提醒

由于会定期性的访问感兴趣的 *Memos* 实例，因此使用前请先征求对方的同意。

尤其避免被对方误判为爬虫、恶意访问者。

## TODO

- 等 *Memos* API 支持指定某一条 *Memo* `ID`作为起止点后，优化拉取的逻辑。当前的实现是直接先拉指定数量条，然后按照发送时间过滤
- 尽量争取把该功能合并到 *Memos* 项目中。

## 问题反馈

请在 *Memos* 的 [Telegram用户群](https://t.me/+-_tNF1k70UU4ZTc9) 里提出你的反馈。
