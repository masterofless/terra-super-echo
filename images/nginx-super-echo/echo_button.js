'use strict';

class LikeButton extends React.Component {
  constructor(props) {
    super(props);
    this.state = { utterances: '' };
    // This binding is necessary to make `this` work in the callback
    this.handleClick = this.handleClick.bind(this);
  }

  handleClick() {
      this.setState({ });
      fetch('http://jsonplaceholder.typicode.com/users')
        .then(res => res.json())
        .then((data) => {
          this.setState({ utterances: data });
        })
        .catch(console.log)
  }

  render() {
    return (
      <div>
        <button onClick={this.handleClick}>Terra Super Echo!</button>
        <p/>
        <label htmlFor="utterances">Tell us your sayings!</label>
        <textarea id="story" name="utterances" value={this.state.utterances} readOnly />
      </div>
    );
  }
}

ReactDOM.render(
  <LikeButton />,
  document.querySelector('#like_button_container')
);
