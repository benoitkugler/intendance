cd frontend ;
npm run build ;
echo "Copying ..." ;
rm -rf ../server/static/app/*
cp -r dist/* ../server/static/app ;
echo "Done"