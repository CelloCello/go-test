rm ./test.conf
if [ ! -e ./test.conf ] ; then 
    cat > ./test.conf << EOF
db-driver: postgres
db-url: postgresql:///cms?host=/var/run/postgresql
logger: zap-prod
logurl: 'tcp://127.0.0.1:5170'
fs-path: /acus/runtime/
fs-default-acf-path: /acus/patterns/
web-bind: ':443'
web-fe: latest
EOF
fi