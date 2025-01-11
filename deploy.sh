# 1. 构建新镜像
docker build -t deepseek-http:latest .

# 2. 停止旧容器（如果存在）
docker stop deepseek-http || true
docker rm deepseek-http || true

# 3. 启动新容器
docker run -d \
  --name deepseek-http \
  -p 8089:8089 \
  --restart unless-stopped \
  deepseek-http:latest

# 4. 查看容器日志
# docker logs -f deepseek-http

# 5. 测试服务
# wget --no-verbose --tries=1 --spider http://localhost:8089/health || echo "deloy Failed!!!"