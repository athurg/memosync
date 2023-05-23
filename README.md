# Memosync

订阅感兴趣的 [Memos](https://usememos.com) ，并在 **Explore** 中浏览。

## 安装（Systemd）

- 下载或编译二进制文件`memosync`，并放到 `/usr/local/bin/` 下
- 在 */etc/systemd/system/* 中，参考 [memosync.service](memosync.service) 创建自己的Systemd配置文件
- 执行 `sudo systemctl daemon-reload` 加载服务配置
- 执行 `sudo systemctl enable memos_discovery` 设置自动启动
- 执行 `sudo systemctl start memos_discovery` 启动服务

如需查看日志，可执行 `journalctl -f -u memosync`

## 使用方法

- 在自己的 *Memos* 中，创建一个新用户，用户名为要订阅的 *Memos* 的URL
- 打开自己 *Memos* 的 **Expolre** 频道
- Enjoy

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
