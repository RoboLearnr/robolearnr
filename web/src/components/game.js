import React from 'react';
import Map from './map'
import MapService from './../services/mapService';

class Game extends React.Component {
    constructor(props) {
        super(props);

        const mapService = new MapService();
        mapService.fetch().then((apiResponse) => {
            const grid = apiResponse.grid.map((row) => {
                return row.map((cell) => {
                    switch (cell) {
                        case "e":
                            return Game.cell("empty");
                        case "w":
                            return Game.cell("wall");
                        case "c":
                            return Game.cell("car");
                        default:
                            return Game.cell("empty");
                    }
                });
            });

            this.setState({grid});
        });

        this.state = {
            grid: [],
        }
    }


    componentWillMount() {
        document.addEventListener("keydown", (event) => {
            const KEY_R = 82;
            const KEY_UP = 38;
            switch( event.keyCode ) {
                case KEY_R:
                    this.rotateCar();
                    break;
                case KEY_UP:
                    this.moveCarForward();
                    break;
                default:
                    //console.log(event.keyCode);
                    break;
            }
        });
    }

    rotateCar() {
        const grid = this.state.grid.map((row) => {
            return row.map((cell) => {
                if (cell.type === "car") {
                    cell.options.rotation = cell.options.rotation + 90;
                    if (cell.options.rotation  >= 360) cell.options.rotation  = 0;
                }

                return cell
            })
        });
        this.setState({grid});
        this.playSound("step");
    }

    moveCarForward() {
        const oldPosition = this.findCarPosition();
        const [x, y] = oldPosition;
        const newPosition = {
            "right": [x, y+1],
            "down": [x+1, y],
            "left": [x, y-1],
            "up": [x-1, y],
        }[this.findCarDirection()];

        const [newX, newY] = newPosition;

        if (!(newX in this.state.grid) || !(newY in this.state.grid[newX])) {
            console.log("out of bound");
            this.playSound("wall");
            return;
        }
        if (this.state.grid[newX][newY].type === "wall") {
            console.log("Dont break the wall");
            this.playSound("wall");
            return;
        }

        this.swapTiles(oldPosition, newPosition)
    }

    swapTiles(oldPosition, newPosition) {
        const [x, y] = oldPosition;
        const [newX, newY] = newPosition;

        let oldCell = null;
        const grid = this.state.grid.map((row, xIndex) => {
            return row.map((cell, yIndex) => {
                if (xIndex === newX && yIndex === newY) {
                    return Object.assign({}, this.state.grid[x][y]);
                }
                else if (xIndex === x && yIndex === y) {
                    oldCell = Object.assign({}, this.state.grid[newX][newY]);
                    return oldCell;
                }

                return cell;
            });
        });

        this.setState({grid});
        this.playSound("step");
    }

    playSound(name) {
        const audio = new Audio(`/sounds/${name}.mp3`);
        audio.play();
    }

    findCarDirection() {
        const [x, y] = this.findCarPosition();
        const rotation = this.state.grid[x][y].options.rotation;

        return {0: "right", 90: "down", 180: "left", 270: "up"}[rotation];
    }

    findCarPosition() {
        let xIndex = 0;
        return this.state.grid.reduce((acc, row) => {
            xIndex++;
            let yIndex = -1;
            return row.reduce((rAcc, cell) => {
                yIndex++
                if (cell.type === "car") {
                    return [xIndex, yIndex];
                }
                return rAcc;
            }, acc);
        });
    }

    static cell(type) {
        let cell = {"type": type, "options": {}};

        if (type === "car") {
            cell.options.rotation = 0;
        }

        return cell
    }

    render() {
        return (
	  <div className="game">
            <div className="game-board">
              <Map grid={this.state.grid}/>
            </div>
	  </div>
        );
    }
}

export default Game
