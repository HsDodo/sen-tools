# debug 远程调试
dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient attach <pid>
# 查日志
less -f /var/log/syslog
# 查进程
ps aux | grep <process_name>
# 查端口
netstat -anp | grep <port>
# 查端口占用进程
lsof -i:<port>
# 列出所有进程
ps -ef | grep <process_name>
# 杀死进程
kill -9 <pid>
# 杀死进程
killall -9 <process_name>
# 查看服务状态
systemctl status <service_name>
# 启动服务
systemctl start <service_name>
# 停止服务
systemctl stop <service_name>
# 重启服务
systemctl restart <service_name>       