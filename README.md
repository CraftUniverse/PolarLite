# PolarLite

The [PolarLite format](./FORMAT.md) is a lightweight format to store minecraft worlds, it's based of the [Polar format](https://github.com/hollow-cube/polar/blob/main/FORMAT.md), but with a few changes.

### Table of Contents

- [Format Changes](#format-changes)
- [CLI](#cli)
- [License](#license)
- [Contributers](#contributers)

## Format Changes

A list of changes between [Polar](https://github.com/hollow-cube/polar/blob/main/FORMAT.md) and [PolarLite](./FORMAT.md)

- Made `zstd` compression a must
- Changed from [`Network Buffer`](https://javadoc.minestom.net/net/minestom/server/network/NetworkBuffer.html) to [`MessagePack`](https://msgpack.org/)
- Changed file extension from `.polar` to `.polarlite`
- Changed `Magic Number` from `Polr` to `PolrLte`
- Removed Array length fields
- Removed `User Data`-fields
- Removed `Block Palette`'s
  - `Block Palette Data` using `Blockstate IDs`
  - `Blockstate IDs` are providied by the Registry
- Removed `Biome Palette`'s
  - `Biome Palette Data` using `Biome IDs`
  - `Biome IDs` are providied by the Registry
- Removed `Block Light` and `Sky Light`
- Removed `Has ID` from `Block Entity`
  - `Block Entity ID` will be `null` instead

## CLI

```
polarlite
    --help, -h
        Show this message and exists

    --version, -v
        Prints version and exists

    --world, -w <string>
        Path to the world folder

    --output, -o <string>
        Path to the output file
```

## License

[PolarLite](https://github.com/CraftUniverse/PolarLite) © 2025 by [CraftUniverse](https://github.com/CraftUniverse) is licensed under [Creative Commons Attribution-ShareAlike 4.0 International](https://creativecommons.org/licenses/by-sa/4.0/?ref=chooser-v1)

## Contributers

- [@Turboman3000](https://github.com/Turboman3000)
