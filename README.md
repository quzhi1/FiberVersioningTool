# FiberVersioningTool
I am writing a automatic versioning library.

With this library, an API server can support many backward incompatible versions without changing its core logic.

The sample API server support `POST localhost:8080/`

## API version 1.0 schema
### Request body
```json
{
    "name": "Zhi Qu"
}
```

## Request header
- Nylas-Api-Version: 1.0
- Metadata: <key>:<value>

## Query parameters
- language-code=<lang>-<region>

### Response body
```json
{
    "id": "af9c3b85-86c5-4771-b07f-22440de6d582",
    "name": "Zhi Qu",
    "created_time": 1657150876
}
```

### Reponse header
- Metadata: <key>:<value>

## API version 1.1 schema

### Request body
```json
{
    "first_name": "Zhi",
    "last_name": "Qu"
}
```

## Request header
- Nylas-Api-Version: 1.0
- Client-Metadata: hello:world

## Query parameters
- lang:<lang>
- region:<region>

### Response body
```json
{
    "id": "c497668b-89c7-48cd-b659-867f3287b023",
    "name": "Zhi Qu"
}
```

### Reponse header
- Client-Metadata: hello:world

## How to play with it
Version 1.1 is not backward compatible with version 1.0. Specifically:
1. Request body: `name` field splitted into two fields `first_name` and `last_name`.
2. Response body: `created_time` field got deleted.
3. Request header: `Metadata: <key:value>` becomes `Client-Metadata:<key:value>`.
4. Response header: Same change as request header.
5. Query parameter: `language-code=<lang>-<region>` becomes `lang:<lang>` and `region:<region>`

You can use `Nylas-Api-Version` header to control your versioning.

If you want to use 1.0, then
```bash
curl --location --request POST 'localhost:8080/?language-code=en-US' \
--header 'Nylas-API-Version: 1.0' \
--header 'Metadata: hello:world' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Zhi Qu"
}'
```

If you want to use 1.1, then
```bash
curl --location --request POST 'localhost:8080?region=US&lang=en' \
--header 'Nylas-API-Version: 1.1' \
--header 'Client-Metadata: hello:world' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name": "Zhi",
    "last_name": "Qu"
}'
```

You can see both version give you the result. The middleware automatically handles versioning.
