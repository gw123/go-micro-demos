# Qrcode Service

This is the Qrcode service

Generated with

```
micro new qrcode
```

## Usage

Generate the proto code

```
make proto
```

Run the service

```
micro run .
```

call getQrcode
```
micro call qrcode Qrcode.GetQrcode '{"content": "91828939832", "size":12 ,"format": 1}'

{
        "msg": "content: 91828939832, size: 12, format: 2"
}
```