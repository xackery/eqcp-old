# eqcp
EQEmu Control Panel


curl --request POST "localhost:8081/v1/petition" -d '{"values":test4", "accountname":"test5"}}'
curl --request PUT "localhost:8081/v1/bug/" -d '{"values":{"charactername":"test4", "hp":"12345"}}'
curl --request PATCH "localhost:8081/v1/bug/1" -d '{"key":"charactername","value":"test5"}'
curl --request PUT "localhost:8081/v1/bug/test" -d '{"key":"charactername","value":"test4"}'
curl --request GET "localhost:8081/v1/bug/search/test"
curl --request GET "localhost:8081/v1/bug/1"