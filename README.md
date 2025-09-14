# GOBOL-LSP

Language Server for [GOBOL](https://github.com/jtomaspm/GOBOL)

## Install 

### Linux
```bash
curl -fsS https://raw.githubusercontent.com/jtomaspm/GOBOL-LSP/refs/heads/main/scripts/install.sh | sh
```

## Run

```bash
GOBOL-LSP
```

### Run flags

#### Logs

Logs are disabled by default. To enable specify a path using:
```bash
GOBOL-LSP -log_path <path>
```

#### Interface

Default interface is stdio, to change use:
```bash
GOBOL-LSP -interface <interface>
```

Interfaces:
 - stdio ✅
 - tcp   ❌ (not supported currently)