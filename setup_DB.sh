dropdb --if-exists intendance && 
createdb intendance && 
psql intendance < create.sql && 
psql intendance < init_dev.sql
