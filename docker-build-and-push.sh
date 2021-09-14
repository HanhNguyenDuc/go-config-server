echo "building smartivr-config-server:$1"
sudo docker login -u smarthub -p smarthub@123! dockerhub.smartcall.ai
sudo docker build -t smartivr-config-server:$1 .
sudo docker tag smartivr-config-server:$1 dockerhub.smartcall.ai/smartivr-config-server:$1
sudo docker push dockerhub.smartcall.ai/smartivr-config-server:$1

