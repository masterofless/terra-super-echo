'use strict';

class LikeButton extends React.Component {
  constructor(props) {
    super(props);
    this.state = { liked: false };
    // This binding is necessary to make `this` work in the callback
    this.handleClick = this.handleClick.bind(this);
  }

  handleClick() {
      this.setState({ liked: true });
    /*
      fetch('http://jsonplaceholder.typicode.com/users')
        .then(res => res.json())
        .then((data) => {
          this.setState({ liked: false });
        })
        .catch(console.log)
        */
  }

  render() {
    return (
      <button onClick={this.handleClick}>
        {this.state.liked ? 'ON' : 'OFF'}
      </button>
    );
  }
}

ReactDOM.render(
  <LikeButton />,
  document.querySelector('#like_button_container')
);
