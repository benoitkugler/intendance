../structgen/structgen --source=server/models/models.go \
    --mode=sql:server/models/scans.go \
    --mode=sql_test:server/models/scans_test.go \
    --mode=sql_gen:create.sql \
    --mode=rand:server/models/data_test.go

../structgen/structgen --source=server/controller/types.go --mode=ts:frontend/src/logic/types.ts 

goimports -w server/models/scans.go
goimports -w server/models/scans_test.go
goimports -w server/models/data_test.go