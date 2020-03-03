# eqcp
EQEmu Control Panel


## Client
```
nvm use lts/erbium
make client starts a dev environment
make client-ssr makes a SSR-production copy
make client-build generates a static version of client

```
curl --request POST "localhost:8081/v1/petition" -d '{"values":test4", "accountname":"test5"}}'

curl --request PUT "localhost:8081/v1/bug/1" -d '{"values":{"charactername":"test4", "hp":"12345"}}'

curl --request PATCH "localhost:8081/v1/bug/1" -d '{"key":"charactername","value":"test5"}'

curl --request GET "localhost:8081/v1/bug/search/test"

curl --request GET "localhost:8081/v1/bug/1"

curl localhost:8081/v1/item/search?values[classes]=8?limit=2

curl --request POST "localhost:8081/v1/loginserver/login" -d '{"username":test", "password":"test"}'

curl "localhost:8081/v1/inventory/search?values[charid]=693603"
```

consider auto generating rsa if not provided? with https://github.com/tmc/genkey ?

```bash
openssl genrsa -out eqcp.rsa 512
openssl rsa -in eqcp.rsa -pubout > eqcp.rsa.pub
```

e.g. 
genkey -bits 4096 > /dev/null  2.44s user 0.02s system 99% cpu 2.458 total
genkey -bits 4096 > /dev/null  6.81s user 0.03s system 99% cpu 6.839 total
genkey -bits 4096 > /dev/null  2.17s user 0.01s system 99% cpu 2.174 total
vs
openssl genrsa 4096  0.76s user 0.00s system 99% cpu 0.771 total
openssl genrsa 4096  0.47s user 0.00s system 99% cpu 0.475 total
openssl genrsa 4096  0.80s user 0.00s system 99% cpu 0.807 total
openssl genrsa 4096  1.22s user 0.00s system 99% cpu 1.231 total


```sql
SELECT * FROM spawnentry se 
INNER JOIN spawn2 s2 ON s2.`spawngroupid` = se.`spawngroupID` 
INNER JOIN npc_types n ON n.id = se.npcid
WHERE s2.enabled = 1 AND s2.zone = "qeynos" 


SELECT item_id FROM lootdrop_entries WHERE lootdrop_id IN
(SELECT lootdrop_id FROM loottable_entries WHERE loottable_id IN
(SELECT n.loottable_id FROM spawnentry se 
INNER JOIN spawn2 s2 ON s2.`spawngroupid` = se.`spawngroupID` 
INNER JOIN npc_types n ON n.id = se.npcid
WHERE s2.enabled = 1 AND s2.zone = "qeynos"
)
)
```

