echo off
set IP_SERVER="192.168.18.115"
set GOOS=linux
set GOARCH=amd64
@REM go mod tidy
@REM go build main.go

@REM scp ./app.service pi@192.168.18.115:/home/pi/dipta/dipta.service
@REM scp ./main pi@192.168.18.115:/home/pi/dipta
scp -r ./* pi@192.168.18.115:/home/pi/dipta/app
ssh -t pi@192.168.18.115 "cd /home/pi/dipta/app && chmod +X install.sh && ./install.sh"
