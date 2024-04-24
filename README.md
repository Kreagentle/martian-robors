## Run instructions
Requires go 1.22.

```
go run main.go
```

## Run tests

```
go test ./...
```

## Description 

Develop a program to track the movement of robots on a rectangular grid, representing the surface of Mars. The program must interpret instructions sent from Earth and report the final position of each robot.

#### Robot Behavior:

Robots are defined by their grid coordinates (x, y) and orientation (N, S, E, W).
Instructions consist of strings containing "L", "R", and "F" for left, right, and forward movements.
Movement instructions cause robots to turn or move one grid point in their current orientation.

#### Grid Boundaries and Lost Robots:

The grid is rectangular and bounded.
If a robot moves off the grid, it is lost forever, leaving a "scent" at its last position.
Future robots encountering a scent at a grid point will ignore instructions to move off the grid from that point.

## Input
The input starts with the upper-right coordinates of the rectangular world, with the lower-left assumed to be (0, 0). Following this, there is a sequence of robot positions and instructions. Each robot's position includes initial coordinates and orientation (N, S, E, W), with instructions provided separately. Robots are processed one by one, executing all instructions before the next robot begins.

The maximum coordinate value is 50.
Instruction strings are less than 100 characters long.

## Output
For each robot position and instruction, the output displays the final grid position and orientation. If a robot falls off the grid, "LOST" is printed alongside the position and orientation.

## Sample

After viewing the resulting coordinates, the program prompts the user with the question: "When do you want to stop the program?" Users can input "no" to indicate their desire to halt the program at that point.

```
Size of Mars: 5 3
Robot initial coordinates and direction #1: 1 1 E
Robot route №1: RFRFRFRF
The result coordinates are: 1 1 E
Do you want to continue adding robots? (yes/no): yes
Robot initial coordinates and direction #2: 3 2 N
Robot route №2: FRRFLLFFRRFLL
The result coordinates are: 3 3 N LOST
Do you want to continue adding robots? (yes/no): yes
Robot initial coordinates and direction #3: 0 3 W
Robot route №3: LLFFFLFLFL
The result coordinates are: 2 3 S
Do you want to continue adding robots? (yes/no): no
```

## What to do next
There are some areas on improvement and further development of the program. First, it could read and write files smoothly, making data handling simpler. Next, adding a mode to show how the program works step by step would help users understand it better. Putting the program in Docker would speed up setup. Lastly, improving the interface would make it easier for anyone to use. These changes would make the program more useful, understandable, and user-friendly, meeting a variety of needs and being easier to work with.