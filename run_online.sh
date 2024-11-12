stop(){
	pkill business
	echo "停止business服务"

	pkill logic
	echo "停止logic服务"

	pkill connect
	echo "停止connect服务"

	cd ../file
	pkill file


}
start(){
	pkill business
	echo "停止business服务"
	nohup ./business > log_business.log 2>&1 &
	echo "启动business服务"



	pkill logic
	echo "停止logic服务"
	nohup ./logic > log_logic.log 2>&1 &
	echo "启动logic服务"


	pkill connect
	echo "停止connect服务"
	sleep 2
	nohup ./connect > log_connect.log 2>&1 &
	echo "启动connect服务"

	cd ../file
	pkill file
	echo "停止file服务"
	nohup ./file &
	echo "启动file服务"
	
}


 
case "$1" in  
start)  
start  
;;  
stop)  
stop  
;;    
restart)  
stop  
start  
;;  
*)  
printf 'Usage: %s {start|stop|restart}\n' "$prog"  
exit 1  
;;  
esac