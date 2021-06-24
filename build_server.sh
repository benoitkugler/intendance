git pull &&
echo "Compiling..." &&
go build server/main.go &&
echo "./main généré. Running dry..." && 
./main -dry && 
echo "Done."