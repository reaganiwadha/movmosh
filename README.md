# movmosh
Mosh MOV files.

## Usage
```bash
movmosh -input IMG_0246.MOV -output output.mov --rate 0.5 --startp 20 --endp 96 --chunk 222 --mode copyswap
```

## Options
- `--input`: Input file.
- `--output`: Output file.
- `--rate`: How much the moshing occurs, can accept from 0-100.
- `--startp`: Start percentage of the file range to mosh.
- `--endp`: End percentage of the file range to mosh.
- `--chunk`: The chunksize of data to be corrupted.
- `--mode`: The moshing mode, can be `swap` (default), `copyswap`, `blackout`, `purerandom`.

## License
MIT
