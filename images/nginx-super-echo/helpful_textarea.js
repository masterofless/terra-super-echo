'use strict';

class UsefulTextarea extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return <textarea/>;
  }
}

const textContainer = document.querySelector('#helpful_text_container');
const el = React.createElement(UsefulTextarea)
ReactDOM.render(el, textContainer);
