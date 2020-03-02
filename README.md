# eqcp
EQEmu Control Panel

```
curl --request POST "localhost:8081/v1/petition" -d '{"values":test4", "accountname":"test5"}}'

curl --request PUT "localhost:8081/v1/bug/1" -d '{"values":{"charactername":"test4", "hp":"12345"}}'

curl --request PATCH "localhost:8081/v1/bug/1" -d '{"key":"charactername","value":"test5"}'

curl --request GET "localhost:8081/v1/bug/search/test"

curl --request GET "localhost:8081/v1/bug/1"

curl localhost:8081/v1/item/search?values[classes]=8

curl --request POST "localhost:8081/v1/loginserver/login" -d '{"username":test", "password":"test"}'

curl "localhost:8081/v1/inventory/search?values[charid]=693603"
```


openssl genrsa -out eqcp.rsa 512
openssl rsa -in eqcp.rsa -pubout > eqcp.rsa.pub

```sql
SELECT * FROM spawnentry se 
INNER JOIN spawn2 s2 ON s2.`spawngroupid` = se.`spawngroupID` 
INNER JOIN npc_types n ON n.id = se.npcid
WHERE s2.enabled = 1 AND s2.zone = "qeynos" 
```

