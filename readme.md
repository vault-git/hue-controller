# Project
GO cli tool that enables you to control Phillips smart lights through the HUE Bridge.

On first use a config file needs to be created, this can be done using the "--register"
command and following the printed instructions.

## Example

```
Usage of /tmp/go-build2690300984/b001/exe/cmd:
  -br float
        Controls the brightness of the given light. [0 - 100] (default -1)
  -colorx float
        Controls the X Coordinate in the color diagram. [0.0 - 1.0] (default -1)
  -colory float
        Controls the Y Coordinate in the color diagram. [0.0 - 1.0] (default -1)
  -light string
        ID of the light to control.
  -list
        Lists all registered lights.
  -register
        Creates a config and registers a new HUE api key.
```

Output of --list:

```
    hue_color_2: Id: 32feab4c-db96-409c-bd28-aef734c5c315, State: Off, Brightness: 1.2, ColorCoords: 0.4317, 0.4026
    hue_color_1: Id: 97f2d988-f2dc-4150-9ba5-f788472c569b, State: Off, Brightness: 1.2, ColorCoords: 0.4317, 0.4026
    hue_color800_1: Id: 8535734d-7136-45e6-a551-70b5f2782fa8, State: On, Brightness: 100.0, ColorCoords: 0.4300, 0.4000
```

To turn on a light and set the color coordinates:

```
    ./hue-controller --light 8535734d-7136-45e6-a551-70b5f2782fa8 --br 100 --colorx 0.5 --colory 0.4
```
