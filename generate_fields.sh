cd server/models &&
    ../../../scaneo/scaneo models.go &&
    cd ../..

../struct2ts/struct2ts --output=frontend/src/logic/types.ts --source=server/views/types.go

go run macros/enums.go >frontend/src/logic/enums.ts
