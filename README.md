# nwCraft

nwCraft is a tool to help with crafting in mmo's. 

Example using data in sampleData.json
```
~\go\src\github.com\chris-carpenter\nwCraft [update-readme ≡]> nwCraft.exe --item "Linen Leggings" --pretty
{
    "name": "Linen Leggings",
    "ingredients": {
        "Fibers": 40,
        "Iron Ore": 8,
        "Rawhide": 24
    }
}
```

Build with
```
go install
```

Command Options
```
[-i|--item] is required
usage: nwCraft [-h|--help] -i|--item "<value>" [-f|--file "<value>"]
               [-d|--debug-level (INFO|DEBUG)] [-p|--pretty]

               Provides shopping list for desired craft

Arguments:

  -h  --help         Print help information
  -i  --item         Item to craft
  -f  --file         File of craftData to load. Default: sampleData.json
  -d  --debug-level  Logging debug level
  -p  --pretty       Pretty output

```

A recipe consists of an item name and a map, item name -> count, of ingredients. The tool will break down the 
ingredients into the lowest level provided for each given ingredient.

Looks at [sampleData.json](./sampleData.json) for data file structure.

