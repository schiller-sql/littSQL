# regenerate models from middleware.sql, requires manual work
# the gen tool from https://github.com/smallnest/gen#Generated-Samples has to be installed

rm -rf model || echo "could not delete the model directory" && exit 1

gen --connstr="host=localhost port=5432 dbname=postgres sslmode=disable" \
    --sqltype=postgres \
    --database="postgres" \
    --gorm --json
