# RoboLearnr

RoboLearnr is simple to use tool to enable beginners to start programming.

Is is a all-in-one tool written in [Go](https://golang.org/) to render maps and allow some basic instruction through
an api. To simplify getting strarted some sdk's are provided in different languages.

## Installation

Download the [latest release](https://github.com/RoboLearnr/robolearnr/releases).

## Example

Start the server and open http://127.0.0.1:9000

```bash
./robolearnr[.exe] https://raw.githubusercontent.com/RoboLearnr/robolearnr/master/maps/robolearn.txt

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

