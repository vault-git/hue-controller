# Project
GO cli tool that enables you to control Phillips smart lights through the HUE Bridge.

On first use a config file needs to be created, this can be done using the "--register"
command and following the printed instructions.

## Example

```
Usage of /tmp/go-build3443625824/b001/exe/cmd:
  -br float
        Controls the brightness of the given light. [0 - 100] (default -1)
  -colorx float
        Controls the X Coordinate in the color diagram. [0.0 - 1.0] (default -1)
  -colory float
        Controls the Y Coordinate in the color diagram. [0.0 - 1.0] (default -1)
  -light string
        Name of the light to control.
  -list
        Lists all registered lights.
  -register
        Creates a config and registers a new HUE api key.
```
