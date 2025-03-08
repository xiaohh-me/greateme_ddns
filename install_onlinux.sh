#!/bin/bash

# 检查是否以 sudo 权限运行
if [ "$EUID" -ne 0 ]; then
  echo "请使用 sudo 权限运行此脚本。"
  exit 1
fi

# 获取脚本自身实际目录
SCRIPT_DIR=$(dirname "$(readlink -f "$0")")

# 安装至指定目录 /usr/local
cp -r $SCRIPT_DIR /usr/local

# 创建变量值并赋予执行权限
SERVICE_PATH="/usr/local/ddns"
SERVICE_FILE_EXE="greateme_ddns"
SERVICE_FILE_CONF="conf/config.ini"

chmod +x $SERVICE_PATH/$SERVICE_FILE_EXE

# 创建 systemd 服务单元文件
SERVICE_FILE="/etc/systemd/system/ddns_aliyun.service"
echo "[Unit]" > $SERVICE_FILE
echo "Description=Greateme DDNS Service" >> $SERVICE_FILE
echo "After=network.target" >> $SERVICE_FILE
echo "" >> $SERVICE_FILE
echo "[Service]" >> $SERVICE_FILE
echo "ExecStart=$SERVICE_PATH/$SERVICE_FILE_EXE $SERVICE_PATH/$SERVICE_FILE_CONF" >> $SERVICE_FILE
echo "Restart=on-failure" >> $SERVICE_FILE
echo "RestartSec=5" >> $SERVICE_FILE
echo "" >> $SERVICE_FILE
echo "[Install]" >> $SERVICE_FILE
echo "WantedBy=multi-user.target" >> $SERVICE_FILE

# 重新加载 systemd 配置
systemctl daemon-reload

# 启用并启动服务
systemctl enable ddns_aliyun.service
systemctl start ddns_aliyun.service

# 检查服务状态
systemctl status ddns_aliyun.service
