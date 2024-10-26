cd cmd/business

pkill business
echo "停止business服务"
nohup ./business &
echo "启动business服务"

cd ../logic
pkill logic
echo "停止logic服务"
nohup ./logic &
echo "启动logic服务"

cd ../connect
pkill connect
echo "停止connect服务"
sleep 2
nohup ./connect &
echo "启动connect服务"

cd ../file
pkill file
echo "停止file服务"
nohup ./file &
echo "启动file服务"

