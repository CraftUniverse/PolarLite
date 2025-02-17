# PolarLite v1.0

The polarlite format is a lightweight format to store minecraft worlds, it's based of the [Polar format](https://github.com/hollow-cube/polar/blob/main/FORMAT.md), but with a few changes.

polarlite is always compressed with zstd. And is serialized with [MessagePack](https://msgpack.org/).

### Header

| Name         | Type  | Notes     |
| ------------ | ----- | --------- |
| Magic Number | int   | `PolrLte` |
| Version      | short |           |
| World        | world |           |

### World

| Name        | Type         | Notes                              |
| ----------- | ------------ | ---------------------------------- |
| Min Section | byte         | For example, -4 in a vanilla world |
| Max Section | byte         | For example, 19 in a vanilla world |
| Chunks      | array[chunk] | Chunk data                         |

### Chunk

Entities or some other extra data field needs to be added to chunks in the future.

| Name           | Type                | Notes                                                                                |
| -------------- | ------------------- | ------------------------------------------------------------------------------------ |
| Chunk X        | varint              |                                                                                      |
| Chunk Z        | varint              |                                                                                      |
| Sections       | array[section]      | `maxSection-minSection+1` entries                                                    |
| Block Entities | array[block entity] |                                                                                      |
| Heightmap Mask | int                 | A mask indicating which heightmaps are present. See `AnvilChunk` for flag constants. |
| Heightmaps     | array[bytes]        | One heightmap for each bit present in Heightmap Mask                                 |

### Sections

| Name               | Type        | Notes                                                     |
| ------------------ | ----------- | --------------------------------------------------------- |
| Is Empty           | bool        | If set, nothing follows                                   |
| Block Palette Data | array[long] | See the anvil format for more information about this type |
| Biome Palette Data | array[long] | See the anvil format for more information about this type |

### Block Entity

| Name            | Type   | Notes                         |
| --------------- | ------ | ----------------------------- |
| Chunk Pos       | int    |                               |
| Block Entity ID | string |                               |
| Has NBT Data    | bool   | If unset, NBT Data is omitted |
| NBT Data        | nbt    |                               |
