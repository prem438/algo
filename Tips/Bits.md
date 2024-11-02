# Bits Operations

## Set the lowest set bits
```
x & (x-1)
```

## Extracts the lower bit set
```
x & ~(x-1)
```

## create mask of n bits
````
(1<<n)-1
````