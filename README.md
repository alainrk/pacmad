# Pacmad
[![YourActionName Actions Status](https://github.com/alainrk/pacmad/workflows/Go/badge.svg)](https://github.com/alainrk/pacmad/actions)


https://user-images.githubusercontent.com/1529268/147325699-6fbac3bc-441d-4eca-ab79-5551e63419a8.mov


## Usage
```
$ ./pacmad
# Follow interactive instructions
```

## Dev

### Test everything recursively
```
chmod +x test.sh
./test.sh
```

### Build everything recursively
```
go build ./...
```

### Cross compilation for multiple arch
```
chmod +x crosscompile.sh
./crosscompile.sh
```

### Development
Branch *main* is protected and requires PR.
