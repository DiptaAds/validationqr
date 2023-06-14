systemctl daemon-reload
systemctl stop dipta.service
/usr/local/go/bin/go build ./main.go
cp /home/pi/dipta/app/app.service /lib/systemd/system/dipta.service
systemctl daemon-reload
systemctl enable dipta.service
systemctl start dipta.service
systemctl status dipta.service