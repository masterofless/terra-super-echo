'use strict';

class LikeButton extends React.Component {
  constructor(props) {
    super(props);
    this.state = { utterance: '' };
    // This binding is necessary to make `this` work in the callback
    this.handleClick = this.handleClick.bind(this);
  }

  handleClick() {
    this.setState({});
    fetch('/beforfe/utterance')
      .then(res => res.json())
      .then((data) => {
        this.setState({ utterance: data.utterance });
        console.log(this.state.utterance);
      })
      .catch(console.log)
  }

  render() {
    return (
      <div>
        <button onClick={this.handleClick}>Terra Super Echo!</button>
        <p />
        <label htmlFor="utterance">Tell us your sayings!</label>
        <textarea id="story" name="utterance" value={this.state.utterance} readOnly />
      </div>
    );
  }
}

ReactDOM.render(
  <LikeButton />,
  document.querySelector('#like_button_container')
);
