# eqcp
EQEmu Control Panel

```
curl --request POST "localhost:8081/v1/petition" -d '{"values":test4", "accountname":"test5"}}'

curl --request PUT "localhost:8081/v1/bug/1" -d '{"values":{"charactername":"test4", "hp":"12345"}}'

curl --request PATCH "localhost:8081/v1/bug/1" -d '{"key":"charactername","value":"test5"}'

curl --request GET "localhost:8081/v1/bug/search/test"

curl --request GET "localhost:8081/v1/bug/1"
```

loginserver endpoints: (typically port 6000)
```
GET /v1/servers/list
/v1/account/create
/v1/account/create/external
/v1/account/credentials/validate/local
/v1/account/credentials/update/local
/v1/account/credentials/update/external
/v1/account/credentials/validate/external
```

openssl genrsa -out eqcp.rsa 512
openssl rsa -in eqcp.rsa -pubout > eqcp.rsa.pub

```sql
SELECT * FROM spawnentry se 
INNER JOIN spawn2 s2 ON s2.`spawngroupid` = se.`spawngroupID` 
INNER JOIN npc_types n ON n.id = se.npcid
WHERE s2.enabled = 1 AND s2.zone = "qeynos"

SELECT * FROM items WHERE id IN 
(SELECT item_id FROM lootdrop_entries 
WHERE lootdrop_id IN 
(SELECT lootdrop_id FROM loottable_entries 
WHERE loottable_id IN 
(SELECT loottable_id FROM spawnentry se 
INNER JOIN spawn2 s2 ON s2.`spawngroupid` = se.`spawngroupID` 
INNER JOIN npc_types n ON n.id = se.npcid
WHERE s2.enabled = 1 
AND s2.zone = "qeynos" 
AND loottable_id > 0)
)
)
```