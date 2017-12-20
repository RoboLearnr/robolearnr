import React from 'react';

class Square extends React.Component {
  render() {
    let classNames = [
      "square",
      this.props.type,
      "rotation-"+(this.props.rotation || 0)
    ];

    return (
      <div className={classNames.join(" ")}>
      </div>
    );
  }
}

export default Square
