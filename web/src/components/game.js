import React from 'react';
import Map from './map'
import MapService from './../services/mapService';
import fetch from 'node-fetch'
import ReconnectingWebSocket from 'reconnectingwebsocket'

class Game extends React.Component {
    constructor(props) {
        super(props);

        const mapService = new MapService();
        mapService.fetch()
            .then((apiResponse) => {
                this._setGrid(apiResponse)
            })
            .catch((err) => {
                console.error("Could not fetch map.")
            });

        this.state = {
            grid: [],
        }
    }

    _setGrid(map) {
        const grid = map.grid.map((row) => {
            return row.map((cell) => {
                switch (cell) {
                    case "e":
                        return Game.cell("empty");
                    case "w":
                        return Game.cell("wall");
                    case "g":
                        return Game.cell("goal");
                    case "c":
                        return Game.cell("car", {"rotation": map.car.rotation});
                    default:
                        return Game.cell("empty");
                }
            });
        });

        this.setState({grid});
        // this.playSound("step");
    }

    componentWillMount() {
        const ws = new ReconnectingWebSocket("ws://127.0.0.1:9000/ws");

        ws.onmessage = (evt) => {
            this._handleRpc(JSON.parse(evt.data));
        };

        window.onbeforeunload = function(event) {
            ws.close();
        };

        document.addEventListener("keydown", (event) => {
            const KEY_R = 82;
            const KEY_UP = 38;
            const KEY_F = 70;
            switch( event.keyCode ) {
                case KEY_R:
                    fetch("http://127.0.0.1:9000/api/rotate")
                        .catch((err) => {console.error('Backend down')});
                    break;
                case KEY_F:
                case KEY_UP:
                    fetch("http://127.0.0.1:9000/api/forward")
                        .catch((err) => {console.error('Backend down')});
                    break;
                default:
                    break;
            }
        });
    }

    _handleRpc(event) {
        switch (event.action) {
            case "rotate":
                this.rotateCar();
                break;
            case "forward":
                this.moveCarForward();
                break;
            case "map":
                this._setGrid(event.map);
                break;
            default:
                break;
        }
    }

    playSound(name) {
        const audio = new Audio(`/sounds/${name}.mp3`);
        audio.play();
    }

    static cell(type, options) {
        options = options || {};

        return {"type": type, "options": options};
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
