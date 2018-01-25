# RoboLearnr

## Installation

Download the [latest release](https://github.com/NoUseFreak/robolearnr/releases).

## Example

Download a sample map.

```bash
wget https://raw.githubusercontent.com/NoUseFreak/robolearnr/master/maps/maze_simple.txt
```

Start the server and open http://127.0.0.1:9000

```bash
./robolearnr[.exe] https://raw.githubusercontent.com/NoUseFreak/robolearnr/master/maps/maze_simple.txt

```

Write your program.

```python
import robolearnr
import time

rl = robolearnr.Robolearn()
rl.reset()

while not rl.on_goal():
    while not rl.before_obstacle():
        rl.forward()
        time.sleep(0.05)
    rl.rotate()
```

Run your program.

```
pip install robolearnr-python
python program.py
```

## Credits

 - Sound effects: [freeSFX](http://www.freesfx.co.uk)
                  
