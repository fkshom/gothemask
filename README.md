# gothemask

Gothemask mask the vendor support file, for example `request support information (Juniper)`, `show tech (Cisco)`.

## Prepare

```
git clone https://github.com/fkshom/gothemask
cd gothemask
go build
```

## Usage
```
./themask --type TYPE SRCFILE [DSTFILE]

TYPE:
  - juniper  : mask request support information
  - ipaddress: mask only ipaddresses
```
