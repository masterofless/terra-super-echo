'use strict';

class LikeButton extends React.Component {
  constructor(props) {
    super(props);
    this.state = { liked: false };
  }

  render() {
    if (this.state.liked) {
      return <h3>You liked this.</h3>;
    }

    return <button onClick={this.enhance()}>Like Me</button>;
  }

  enhance() {
      fetch('http://jsonplaceholder.typicode.com/users')
        .then(res => res.json())
        .then((data) => {
          this.setState({ liked: true });
        })
        .catch(console.log)
  }
}

const domContainer = document.querySelector('#like_button_container');
const el = React.createElement(LikeButton)
ReactDOM.render(el, domContainer);
