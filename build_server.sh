git pull &&
echo "Compiling..." &&
go build server/main.go &&
echo "./main created. Running dry..." && 
./main -dry && 
echo "Done."