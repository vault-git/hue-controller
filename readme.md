# Project
GO cli tool that enables you to control Phillips smart lights through the HUE Bridge.

On first use a config file needs to be created, this can be done using the "--register"
command and following the printed instructions.

## Example

Output of --list:

```
hue_color_2:
        Id: 32feab4c-db96-409c-bd28-aef734c5c315,
        State: Off,
        Brightness: 1.2,
        Color: #90C1FF
hue_color_1:
        Id: 97f2d988-f2dc-4150-9ba5-f788472c569b,
        State: Off,
        Brightness: 1.2,
        Color: #90C1FF
hue_color800_1:
        Id: 8535734d-7136-45e6-a551-70b5f2782fa8,
        State: On,
        Brightness: 71.9,
        Color: #FF6B10
```

To turn on a light and set the color to #00FF99:

```
    ./hue-controller --light 8535734d-7136-45e6-a551-70b5f2782fa8 --br 100 --rgb '#00ff99'
```
