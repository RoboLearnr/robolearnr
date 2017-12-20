import React from 'react';
import Square from './square'

class Map extends React.Component {
    render() {
        return this.props.grid.map((row, xIndex) => (
            <div key={xIndex} className="board-row">{
                row.map((cell, yIndex) => (
                    <Square
                        type={this.props.grid[xIndex][yIndex].type}
                        key={xIndex+","+yIndex}
                        rotation={this.props.grid[xIndex][yIndex].options.rotation}
                    />
                ))
            }</div>
        ))
    }
}

export default Map
